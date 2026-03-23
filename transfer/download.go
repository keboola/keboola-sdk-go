package transfer

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"gocloud.dev/blob"

	"github.com/keboola/keboola-sdk-go/v2/pkg/keboola"
	coreabs "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/storage_file_upload/abs"
	coregcs "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/storage_file_upload/gcs"
	cores3 "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/storage_file_upload/s3"
	"github.com/keboola/keboola-sdk-go/v2/transfer/abs"
	"github.com/keboola/keboola-sdk-go/v2/transfer/gcs"
	"github.com/keboola/keboola-sdk-go/v2/transfer/s3"
)

type downloadConfig struct {
	transport http.RoundTripper
}

// DownloadOption configures download behaviour.
type DownloadOption func(c *downloadConfig)

// WithDownloadTransport overrides the HTTP transport used for download requests.
func WithDownloadTransport(transport http.RoundTripper) DownloadOption {
	return func(c *downloadConfig) {
		c.transport = transport
	}
}

// Download downloads the entire file content.
func Download(ctx context.Context, file *keboola.FileDownloadCredentials) ([]byte, error) {
	if file.IsSliced {
		return nil, fmt.Errorf("cannot download a sliced file as a whole file")
	}
	return DownloadSlice(ctx, file, "")
}

// DownloadManifest downloads and parses the manifest for a sliced file, returning the list of slice names.
func DownloadManifest(ctx context.Context, file *keboola.FileDownloadCredentials) (keboola.SlicesList, error) {
	rawManifest, err := DownloadSlice(ctx, file, keboola.ManifestFileName)
	if err != nil {
		return nil, fmt.Errorf("cannot download manifest: %w", err)
	}

	manifest := &keboola.SlicedFileManifest{}
	err = json.Unmarshal(rawManifest, manifest)
	if err != nil {
		return nil, fmt.Errorf("cannot map file contents to manifest: %w", err)
	}

	dstURL, err := file.DestinationURL()
	if err != nil {
		return nil, err
	}
	res := keboola.SlicesList{}
	for _, slice := range manifest.Entries {
		if !strings.HasPrefix(slice.URL, dstURL) {
			return nil, fmt.Errorf(`slice URL "%s" seems wrong: %w`, slice.URL, err)
		}
		res = append(res, strings.TrimPrefix(slice.URL, dstURL))
	}
	return res, nil
}

// DownloadSlice downloads a single slice (or the whole file if slice is "").
func DownloadSlice(ctx context.Context, file *keboola.FileDownloadCredentials, slice string) (out []byte, err error) {
	reader, err := DownloadSliceReader(ctx, file, slice)
	if err != nil {
		return nil, err
	}
	out, err = io.ReadAll(reader)
	if closeErr := reader.Close(); err == nil && closeErr != nil {
		err = closeErr
	}
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DownloadReader returns an io.ReadCloser for the whole file (non-sliced).
func DownloadReader(ctx context.Context, file *keboola.FileDownloadCredentials) (io.ReadCloser, error) {
	return DownloadSliceReader(ctx, file, "")
}

// DownloadManifestReader returns an io.ReadCloser for the manifest file.
func DownloadManifestReader(ctx context.Context, file *keboola.FileDownloadCredentials) (io.ReadCloser, error) {
	return DownloadSliceReader(ctx, file, keboola.ManifestFileName)
}

// DownloadSliceReader returns an io.ReadCloser for the given slice.
func DownloadSliceReader(ctx context.Context, file *keboola.FileDownloadCredentials, slice string, opts ...DownloadOption) (io.ReadCloser, error) {
	c := downloadConfig{}
	for _, opt := range opts {
		opt(&c)
	}
	switch file.Provider {
	case coreabs.Provider:
		return abs.NewDownloadReader(ctx, file.ABSDownloadParams, slice, c.transport)
	case coregcs.Provider:
		return gcs.NewDownloadReader(ctx, file.GCSDownloadParams, slice, c.transport)
	case cores3.Provider:
		return s3.NewDownloadReader(ctx, file.S3DownloadParams, file.Region, slice, c.transport)
	default:
		return nil, fmt.Errorf(`unsupported provider "%s"`, file.Provider)
	}
}

// GetFileAttributes returns metadata (size, content-type, mod time) for the given slice.
func GetFileAttributes(ctx context.Context, file *keboola.FileDownloadCredentials, slice string, opts ...DownloadOption) (*keboola.FileAttributes, error) {
	c := downloadConfig{}
	for _, opt := range opts {
		opt(&c)
	}
	var attrs *blob.Attributes
	var err error
	switch file.Provider {
	case coreabs.Provider:
		attrs, err = abs.GetFileAttributes(ctx, file.ABSDownloadParams, slice, c.transport)
	case coregcs.Provider:
		attrs, err = gcs.GetFileAttributes(ctx, file.GCSDownloadParams, slice, c.transport)
	case cores3.Provider:
		attrs, err = s3.GetFileAttributes(ctx, file.S3DownloadParams, file.Region, slice, c.transport)
	default:
		return nil, fmt.Errorf(`unsupported provider "%s"`, file.Provider)
	}
	if err != nil {
		return nil, err
	}
	return &keboola.FileAttributes{
		ContentType: attrs.ContentType,
		ModTime:     attrs.ModTime,
		Size:        attrs.Size,
	}, nil
}
