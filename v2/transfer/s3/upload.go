package s3

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
	s3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"gocloud.dev/blob"
	"gocloud.dev/blob/s3blob"

	cores3 "github.com/keboola/keboola-sdk-go/v2/pkg/keboola/storage_file_upload/s3"
)

func NewUploadWriter(ctx context.Context, params *cores3.UploadParams, region string, slice string, transport http.RoundTripper) (*blob.Writer, error) {
	// Create static configuration, don't load AWS ENVs
	var cfg aws.Config
	cfg.Region = region
	cfg.Credentials = credentials.NewStaticCredentialsProvider(
		params.Credentials.AccessKeyID,
		params.Credentials.SecretAccessKey,
		params.Credentials.SessionToken,
	)
	if transport != nil {
		cfg.HTTPClient = &http.Client{Transport: transport}
	}

	client := awss3.NewFromConfig(cfg)
	b, err := s3blob.OpenBucketV2(ctx, client, params.Bucket, nil)
	if err != nil {
		return nil, err
	}

	opts := &blob.WriterOptions{
		BeforeWrite: func(as func(any) bool) error {
			var req *awss3.PutObjectInput
			if as(&req) {
				req.ACL = s3types.ObjectCannedACL(params.ACL)
				req.ServerSideEncryption = s3types.ServerSideEncryption(params.Encryption)
			}
			return nil
		},
		// Smaller buffer size for better progress reporting
		// 5MB is AWS's minimum part size for multipart uploads
		BufferSize: cores3.MinUploadPartSize,
	}

	bw, err := b.NewWriter(ctx, cores3.SliceKey(params.Key, slice), opts)
	if err != nil {
		return nil, fmt.Errorf(`opening blob "%s" failed: %w`, params.Key, err)
	}

	return bw, nil
}

func NewDownloadReader(ctx context.Context, params *cores3.DownloadParams, region string, slice string, transport http.RoundTripper) (*blob.Reader, error) {
	b, err := openBucketForDownload(ctx, params, region, transport)
	if err != nil {
		return nil, err
	}

	opts := &blob.ReaderOptions{}
	br, err := b.NewReader(ctx, cores3.SliceKey(params.Path.Key, slice), opts)
	if err != nil {
		return nil, fmt.Errorf(`reader: opening blob "%s" failed: %w`, params.Path.Key, err)
	}
	return br, nil
}

func GetFileAttributes(ctx context.Context, params *cores3.DownloadParams, region string, slice string, transport http.RoundTripper) (*blob.Attributes, error) {
	b, err := openBucketForDownload(ctx, params, region, transport)
	if err != nil {
		return nil, err
	}

	return b.Attributes(ctx, cores3.SliceKey(params.Path.Key, slice))
}

func openBucketForDownload(ctx context.Context, params *cores3.DownloadParams, region string, transport http.RoundTripper) (*blob.Bucket, error) {
	cred := config.WithCredentialsProvider(
		credentials.NewStaticCredentialsProvider(
			params.Credentials.AccessKeyID,
			params.Credentials.SecretAccessKey,
			params.Credentials.SessionToken,
		),
	)

	var cfg aws.Config
	var err error
	if transport != nil {
		cfg, err = config.LoadDefaultConfig(ctx, cred, config.WithRegion(region), config.WithHTTPClient(&http.Client{Transport: transport}))
	} else {
		cfg, err = config.LoadDefaultConfig(ctx, cred, config.WithRegion(region))
	}
	if err != nil {
		return nil, err
	}

	client := awss3.NewFromConfig(cfg)
	return s3blob.OpenBucketV2(ctx, client, params.Path.Bucket, nil)
}
