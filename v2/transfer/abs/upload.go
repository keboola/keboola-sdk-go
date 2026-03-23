package abs

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
	"gocloud.dev/blob"
	"gocloud.dev/blob/azureblob"

	coreabs "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/storage_file_upload/abs"
)

// ContainerURL builds the full URL for the given container using the connection string.
func ContainerURL(cs *coreabs.ConnectionString, containerName string) string {
	return runtime.JoinPaths(cs.BlobEndpoint, containerName) + "?" + cs.SharedAccessSignature
}

func NewUploadWriter(ctx context.Context, params *coreabs.UploadParams, slice string, transport http.RoundTripper) (*blob.Writer, error) {
	cs, err := coreabs.ParseConnectionString(params.Credentials.SASConnectionString)
	if err != nil {
		return nil, err
	}

	clientOptions := &container.ClientOptions{}
	if transport != nil {
		clientOptions.Transport = &http.Client{Transport: transport}
	}
	client, err := container.NewClientWithNoCredential(ContainerURL(cs, params.Container), clientOptions)
	if err != nil {
		return nil, err
	}
	b, err := azureblob.OpenBucket(ctx, client, nil)
	if err != nil {
		return nil, err
	}

	// Smaller buffer size for better progress reporting
	opts := &blob.WriterOptions{BufferSize: 512000}
	bw, err := b.NewWriter(ctx, coreabs.SliceKey(params.BlobName, slice), opts)
	if err != nil {
		return nil, fmt.Errorf(`opening blob "%s" failed: %w`, params.BlobName, err)
	}

	return bw, nil
}

func NewDownloadReader(ctx context.Context, params *coreabs.DownloadParams, slice string, transport http.RoundTripper) (*blob.Reader, error) {
	b, err := openBucketForDownload(ctx, params, transport)
	if err != nil {
		return nil, err
	}

	br, err := b.NewReader(ctx, coreabs.SliceKey(params.Path.BlobName, slice), nil)
	if err != nil {
		return nil, fmt.Errorf(`reader: opening blob "%s" failed: %w`, params.Path.BlobName, err)
	}

	return br, nil
}

func GetFileAttributes(ctx context.Context, params *coreabs.DownloadParams, slice string, transport http.RoundTripper) (*blob.Attributes, error) {
	b, err := openBucketForDownload(ctx, params, transport)
	if err != nil {
		return nil, err
	}

	return b.Attributes(ctx, coreabs.SliceKey(params.Path.BlobName, slice))
}

func openBucketForDownload(ctx context.Context, params *coreabs.DownloadParams, transport http.RoundTripper) (*blob.Bucket, error) {
	cs, err := coreabs.ParseConnectionString(params.Credentials.SASConnectionString)
	if err != nil {
		return nil, err
	}

	clientOptions := &container.ClientOptions{}
	if transport != nil {
		clientOptions.Transport = &http.Client{Transport: transport}
	}
	client, err := container.NewClientWithNoCredential(ContainerURL(cs, params.Path.Container), clientOptions)
	if err != nil {
		return nil, err
	}

	return azureblob.OpenBucket(ctx, client, nil)
}
