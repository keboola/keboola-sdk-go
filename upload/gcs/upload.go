package gcs

import (
	"context"
	"fmt"
	"net/http"

	googlestorage "cloud.google.com/go/storage"
	"github.com/googleapis/gax-go/v2"
	"gocloud.dev/blob"
	"gocloud.dev/blob/gcsblob"
	"gocloud.dev/gcp"
	"golang.org/x/oauth2"

	coregcs "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/storage_file_upload/gcs"
)

func NewUploadWriter(ctx context.Context, params *coregcs.UploadParams, slice string, transport http.RoundTripper) (*blob.Writer, error) {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: params.AccessToken,
		TokenType:   params.TokenType,
	})

	if transport == nil {
		transport = gcp.DefaultTransport()
	}
	client, err := gcp.NewHTTPClient(transport, tokenSource)
	if err != nil {
		return nil, err
	}
	b, err := gcsblob.OpenBucket(ctx, client, params.Bucket, nil)
	if err != nil {
		return nil, err
	}

	var gcsClient *googlestorage.Client
	if b.As(&gcsClient) {
		gcsClient.SetRetry(
			googlestorage.WithBackoff(gax.Backoff{}),
			googlestorage.WithPolicy(googlestorage.RetryIdempotent),
		)
	} else {
		panic("Unable to access storage.Client through Bucket.As")
	}

	// Smaller buffer size for better progress reporting
	opts := &blob.WriterOptions{BufferSize: 512000}
	bw, err := b.NewWriter(ctx, coregcs.SliceKey(params.Key, slice), opts)
	if err != nil {
		return nil, fmt.Errorf(`opening blob "%s" failed: %w`, params.Key, err)
	}

	return bw, nil
}

func NewDownloadReader(ctx context.Context, params *coregcs.DownloadParams, slice string, transport http.RoundTripper) (*blob.Reader, error) {
	b, err := openBucketForDownload(ctx, params, transport)
	if err != nil {
		return nil, err
	}

	br, err := b.NewReader(ctx, coregcs.SliceKey(params.Path.Key, slice), nil)
	if err != nil {
		return nil, fmt.Errorf(`opening blob "%s" failed: %w`, params.Path.Key, err)
	}

	return br, nil
}

func GetFileAttributes(ctx context.Context, params *coregcs.DownloadParams, slice string, transport http.RoundTripper) (*blob.Attributes, error) {
	b, err := openBucketForDownload(ctx, params, transport)
	if err != nil {
		return nil, err
	}

	return b.Attributes(ctx, coregcs.SliceKey(params.Path.Key, slice))
}

func openBucketForDownload(ctx context.Context, params *coregcs.DownloadParams, transport http.RoundTripper) (*blob.Bucket, error) {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: params.Credentials.AccessToken,
		TokenType:   params.Credentials.TokenType,
	})

	if transport == nil {
		transport = gcp.DefaultTransport()
	}
	client, err := gcp.NewHTTPClient(transport, tokenSource)
	if err != nil {
		return nil, err
	}
	b, err := gcsblob.OpenBucket(ctx, client, params.Path.Bucket, nil)
	if err != nil {
		return nil, err
	}

	var gcsClient *googlestorage.Client
	if b.As(&gcsClient) {
		gcsClient.SetRetry(
			googlestorage.WithBackoff(gax.Backoff{}),
			googlestorage.WithPolicy(googlestorage.RetryIdempotent),
		)
	} else {
		panic("Unable to access storage.Client through Bucket.As")
	}

	return b, nil
}
