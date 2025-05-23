package keboola

import (
	"context"
	jsonLib "encoding/json"
	"fmt"
	"time"

	"github.com/relvacode/iso8601"

	"github.com/keboola/keboola-sdk-go/v2/pkg/request"
)

const (
	BucketPermissionRead  BucketPermission = "read"
	BucketPermissionWrite BucketPermission = "write"
)

// Token https://keboola.docs.apiary.io/#reference/tokens-and-permissions/token-verification/token-verification
type Token struct {
	Token                 string            `json:"token"` // set manually from request
	ID                    string            `json:"id"`
	Description           string            `json:"description"`
	IsMaster              bool              `json:"isMasterToken"`
	CanManageBuckets      bool              `json:"canManageBuckets"`
	CanManageTokens       bool              `json:"canManageTokens"`
	CanReadAllFileUploads bool              `json:"canReadAllFileUploads"`
	CanPurgeTrash         bool              `json:"canPurgeTrash"`
	Created               iso8601.Time      `json:"created"`
	Refreshed             iso8601.Time      `json:"refreshed"`
	Expires               *iso8601.Time     `json:"expires"`
	IsExpired             bool              `json:"isExpired"`
	IsDisabled            bool              `json:"isDisabled"`
	Owner                 TokenOwner        `json:"owner"`
	Admin                 *TokenAdmin       `json:"admin,omitempty"`
	Creator               *CreatorToken     `json:"creatorToken,omitempty"`
	BucketPermissions     BucketPermissions `json:"bucketPermissions,omitempty"`
	ComponentAccess       []string          `json:"componentAccess,omitempty"`
}

type BucketPermissions map[BucketID]BucketPermission

type BucketPermission string

// TokenAdmin - admin part of the token that should exist if the token is a master token.
type TokenAdmin struct {
	Name                 string   `json:"name"`
	ID                   int      `json:"id"`
	IsOrganizationMember bool     `json:"isOrganizationMember"`
	Role                 string   `json:"role"`
	Features             Features `json:"features"`
}

// TokenOwner - owner of Token.
type TokenOwner struct {
	ID                  int      `json:"id"`
	Name                string   `json:"name"`
	Features            Features `json:"features"`
	HasMysql            bool     `json:"hasMysql"`
	HasSynapse          bool     `json:"hasSynapse"`
	HasRedshift         bool     `json:"hasRedshift"`
	HasSnowflake        bool     `json:"hasSnowflake"`
	HasExasol           bool     `json:"hasExasol"`
	HasTeradata         bool     `json:"hasTeradata"`
	HasBigquery         bool     `json:"hasBigquery"`
	DefaultBackend      string   `json:"defaultBackend"`
	FileStorageProvider string   `json:"fileStorageProvider"`
}

type CreatorToken struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

type createTokenOptions struct {
	Description           string            `writeas:"description"`
	BucketPermissions     map[string]string `writeas:"bucketPermissions" writeoptional:"true"`
	ComponentAccess       []string          `writeas:"componentAccess" writeoptional:"true"`
	CanManageBuckets      bool              `writeas:"canManageBuckets"`
	CanReadAllFileUploads bool              `writeas:"canReadAllFileUploads"`
	CanPurgeTrash         bool              `writeas:"canPurgeTrash"`
	ExpiresIn             int               `writeas:"expiresIn" writeoptional:"true"`
}

type CreateTokenOption func(*createTokenOptions)

// WithDescription sets the token's description.
func WithDescription(description string) CreateTokenOption {
	return func(o *createTokenOptions) { o.Description = description }
}

// WithBucketPermission adds `bucket` to the set of buckets this token may read or write to, depending on the permission specified (`perm`).
func WithBucketPermission(bucketID BucketID, perm BucketPermission) CreateTokenOption {
	return func(o *createTokenOptions) {
		if o.BucketPermissions == nil {
			o.BucketPermissions = make(map[string]string)
		}
		o.BucketPermissions[bucketID.String()] = string(perm)
	}
}

// WithBucketPermissions set token buckets permissions.
func WithBucketPermissions(v BucketPermissions) CreateTokenOption {
	return func(o *createTokenOptions) {
		o.BucketPermissions = make(map[string]string)
		for k, v := range v {
			o.BucketPermissions[k.String()] = string(v)
		}
	}
}

// WithComponentAccess adds `component` to the list of components this token may access.
func WithComponentAccess(component string) CreateTokenOption {
	return func(o *createTokenOptions) { o.ComponentAccess = append(o.ComponentAccess, component) }
}

// WithCanManageBuckets gives the newly created token the ability to manage buckets.
func WithCanManageBuckets(canManageBuckets bool) CreateTokenOption {
	return func(o *createTokenOptions) { o.CanManageBuckets = canManageBuckets }
}

// WithCanReadAllFileUploads allows access to all file uploads. Without this permission, only files uplaoded using the new token are accessible.
func WithCanReadAllFileUploads(canReadAllFileUploads bool) CreateTokenOption {
	return func(o *createTokenOptions) { o.CanReadAllFileUploads = canReadAllFileUploads }
}

// WithCanPurgeTrash allows this token to permanently delete configurations.
func WithCanPurgeTrash(canPurgeTrash bool) CreateTokenOption {
	return func(o *createTokenOptions) { o.CanPurgeTrash = canPurgeTrash }
}

// WithExpiresIn sets the time until the token expires.
func WithExpiresIn(expiresIn time.Duration) CreateTokenOption {
	return func(o *createTokenOptions) { o.ExpiresIn = int(expiresIn.Seconds()) }
}

// VerifyTokenRequest https://keboola.docs.apiary.io/#reference/tokens-and-permissions/token-verification/token-verification
func (a *PublicAPI) VerifyTokenRequest(token string) request.APIRequest[*Token] {
	result := &Token{}
	req := a.
		newRequest(StorageAPI).
		WithResult(result).
		WithGet("tokens/verify").
		AndHeader("X-StorageApi-Token", token).
		WithOnSuccess(func(_ context.Context, _ request.HTTPResponse) error {
			result.Token = token
			return nil
		})
	return request.NewAPIRequest(result, req)
}

// TokenDetailRequest https://keboola.docs.apiary.io/#reference/tokens-and-permissions/token/token-detail
func (a *AuthorizedAPI) TokenDetailRequest(tokenID string) request.APIRequest[*Token] {
	result := &Token{}
	req := a.
		newRequest(StorageAPI).
		WithResult(result).
		WithGet("tokens/{tokenId}").
		AndPathParam("tokenId", tokenID)
	return request.NewAPIRequest(result, req)
}	

// CreateTokenRequest https://keboola.docs.apiary.io/#reference/tokens-and-permissions/tokens-collection/create-token
func (a *AuthorizedAPI) CreateTokenRequest(opts ...CreateTokenOption) request.APIRequest[*Token] {
	options := &createTokenOptions{}
	for _, opt := range opts {
		opt(options)
	}

	result := &Token{}
	req := a.
		newRequest(StorageAPI).
		WithResult(result).
		WithPost("tokens").
		WithJSONBody(request.StructToMap(options, nil))
	return request.NewAPIRequest(result, req)
}

// ListTokensRequest https://keboola.docs.apiary.io/#reference/tokens-and-permissions/tokens-collection/list-all-tokens
func (a *AuthorizedAPI) ListTokensRequest() request.APIRequest[*[]*Token] {
	var result []*Token
	req := a.
		newRequest(StorageAPI).
		WithResult(&result).
		WithGet("tokens")
	return request.NewAPIRequest(&result, req)
}

// DeleteTokenRequest (no documentation).
func (a *AuthorizedAPI) DeleteTokenRequest(tokenID string) request.APIRequest[*Token] {
	result := &Token{}
	req := a.
		newRequest(StorageAPI).
		WithResult(result).
		WithDelete("tokens/{tokenId}").
		AndPathParam("tokenId", tokenID).
		WithOnError(ignoreResourceNotFoundError())
	return request.NewAPIRequest(result, req)
}

// RefreshTokenRequest https://keboola.docs.apiary.io/#reference/tokens-and-permissions/share-token/refresh-token
func (a *AuthorizedAPI) RefreshTokenRequest(tokenID string) request.APIRequest[*Token] {
	result := &Token{}
	req := a.
		newRequest(StorageAPI).
		WithResult(result).
		WithPost("tokens/{tokenId}/refresh").
		AndPathParam("tokenId", tokenID)
	return request.NewAPIRequest(result, req)
}

// ProjectID returns ID of project to which the token belongs.
func (t *Token) ProjectID() int {
	return t.Owner.ID
}

// ProjectName returns name of project to which the token belongs.
func (t *Token) ProjectName() string {
	return t.Owner.Name
}

func (r BucketPermissions) MarshalJSON() ([]byte, error) {
	raw := make(map[string]BucketPermission)
	for k, v := range r {
		raw[k.String()] = v
	}
	return jsonLib.Marshal(raw)
}

// UnmarshalJSON implements JSON decoding.
// The API returns empty array when the results field is empty.
func (r *BucketPermissions) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "[]" {
		*r = BucketPermissions{}
		return nil
	}
	// see https://stackoverflow.com/questions/43176625/call-json-unmarshal-inside-unmarshaljson-function-without-causing-stack-overflow
	raw := make(map[string]BucketPermission)
	if err := jsonLib.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("cannot decode bucket permissions: %w", err)
	}

	// convert key, string -> BucketID
	out := make(map[BucketID]BucketPermission)
	for k, v := range raw {
		bucketID, err := ParseBucketID(k)
		if err != nil {
			return fmt.Errorf("cannot decode bucket permissions: %w", err)
		}
		out[bucketID] = v
	}

	*r = out
	return nil
}
