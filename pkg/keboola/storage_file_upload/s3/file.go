package s3

import (
	"fmt"

	"github.com/relvacode/iso8601"
)

const (
	Provider = "aws"
	// MinUploadPartSize is AWS's minimum part size for multipart uploads (5MB).
	MinUploadPartSize = 5 * 1024 * 1024
)

type Path struct {
	Key    string `json:"key"`
	Bucket string `json:"bucket"`
}

//nolint:tagliatelle
type Credentials struct {
	AccessKeyID     string       `json:"AccessKeyId"`
	SecretAccessKey string       `json:"SecretAccessKey"`
	SessionToken    string       `json:"SessionToken"`
	Expiration      iso8601.Time `json:"Expiration"`
}

//nolint:tagliatelle
type UploadParams struct {
	Path
	Credentials Credentials `json:"credentials"`
	ACL         string      `json:"acl"`
	Encryption  string      `json:"x-amz-server-side-encryption"`
}

type DownloadParams struct {
	Credentials Credentials `json:"credentials"`
	Path        Path        `json:"s3Path"`
}

func (p *DownloadParams) DestinationURL() (string, error) {
	return fmt.Sprintf("s3://%s/%s", p.Path.Bucket, p.Path.Key), nil
}

func NewSliceURL(params *UploadParams, slice string) string {
	return fmt.Sprintf("s3://%s/%s", params.Bucket, SliceKey(params.Key, slice))
}

// SliceKey constructs the blob key for a given slice name.
func SliceKey(key, slice string) string {
	return key + slice
}
