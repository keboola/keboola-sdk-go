package gcs

import (
	"fmt"
)

const Provider = "gcp"

type Path struct {
	Key    string `json:"key"`
	Bucket string `json:"bucket"`
}

//nolint:tagliatelle
type Credentials struct {
	ProjectID   string `json:"projectId"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type UploadParams struct {
	Path
	Credentials
}

type DownloadParams struct {
	Credentials Credentials `json:"gcsCredentials"`
	Path        Path        `json:"gcsPath"`
}

func (p *DownloadParams) DestinationURL() (string, error) {
	return fmt.Sprintf("gs://%s/%s", p.Path.Bucket, p.Path.Key), nil
}

func NewSliceURL(params *UploadParams, slice string) string {
	return fmt.Sprintf(
		"gs://%s/%s",
		params.Bucket,
		SliceKey(params.Key, slice),
	)
}

// SliceKey constructs the blob key for a given slice name.
func SliceKey(key, slice string) string {
	return key + slice
}
