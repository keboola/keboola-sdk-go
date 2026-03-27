package abs

import (
	"fmt"
	"strings"

	"github.com/relvacode/iso8601"
)

const Provider = "azure"

type ConnectionString struct {
	BlobEndpoint          string
	SharedAccessSignature string
}

type Credentials struct {
	SASConnectionString string       `json:"SASConnectionString"`
	Expiration          iso8601.Time `json:"expiration"`
}

type UploadParams struct {
	BlobName    string      `json:"blobName"`
	AccountName string      `json:"accountName"`
	Container   string      `json:"container"`
	Credentials Credentials `json:"absCredentials"`
}

type DownloadParams struct {
	Credentials Credentials `json:"absCredentials"`
	Path        struct {
		BlobName  string `json:"name"`
		Container string `json:"container"`
	} `json:"absPath"`
}

func (p *DownloadParams) DestinationURL() (string, error) {
	cs, err := ParseConnectionString(p.Credentials.SASConnectionString)
	if err != nil {
		return "", err
	}
	blobEndpoint := strings.ReplaceAll(cs.BlobEndpoint, "https://", "azure://")
	return fmt.Sprintf("%s/%s/%s", blobEndpoint, p.Path.Container, p.Path.BlobName), nil
}

func NewSliceURL(params *UploadParams, slice string) string {
	return fmt.Sprintf("azure://%s.blob.core.windows.net/%s/%s", params.AccountName, params.Container, SliceKey(params.BlobName, slice))
}

// SliceKey constructs the blob key for a given slice name.
func SliceKey(key, slice string) string {
	return key + slice
}

// ParseConnectionString parses an Azure Blob Storage SAS connection string.
func ParseConnectionString(str string) (*ConnectionString, error) {
	csMap := make(map[string]string)
	for item := range strings.SplitSeq(str, ";") {
		parts := strings.SplitN(item, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf(`connection string is malformed, it should contain key value pairs separated by semicolons`)
		}
		csMap[parts[0]] = parts[1]
	}
	cs := &ConnectionString{}

	if val, ok := csMap["BlobEndpoint"]; ok {
		cs.BlobEndpoint = val
	} else {
		return nil, fmt.Errorf(`connection string is missing "BlobEndpoint" part`)
	}

	if val, ok := csMap["SharedAccessSignature"]; ok {
		cs.SharedAccessSignature = val
	} else {
		return nil, fmt.Errorf(`connection string is missing "SharedAccessSignature" part`)
	}

	return cs, nil
}
