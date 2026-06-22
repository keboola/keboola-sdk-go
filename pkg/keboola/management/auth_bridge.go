// Hand-written companion to the generated client — not produced by
// openapi-generator, keep when regenerating.
package management

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// HeaderSubjectToken carries the programmatic Connection bearer token
// (kbc_at_* / kbc_pat_*) to the internal auth-bridge resolver endpoint.
// The value uses the "Bearer <token>" scheme.
const HeaderSubjectToken = "X-Subject-Token" //nolint:gosec // header name, not a credential

// AuthBridgeAPIService handles Connection's internal auth-bridge endpoints.
//
// The endpoints require internal Kubernetes ServiceAccount authentication
// (X-Kubernetes-Authorization) with the internal:auth-bridge:resolve-storage-token
// scope — build the client with NewAPIClientWithAuth and a KeboolaServiceAccountAuth.
type AuthBridgeAPIService service

// AuthBridgeAPI returns the hand-written auth-bridge service for this client.
// It is a method instead of a generated APIClient field so it survives
// client regeneration.
func (c *APIClient) AuthBridgeAPI() *AuthBridgeAPIService {
	return &AuthBridgeAPIService{client: c}
}

// AuthBridgeStorageTokenResolveRequest is the JSON body of the resolver endpoint.
type AuthBridgeStorageTokenResolveRequest struct {
	// ProjectID is the project whose Storage token should be resolved.
	ProjectID int `json:"projectId"`
}

// AuthBridgeStorageTokenResolveResponse is the resolver endpoint response.
type AuthBridgeStorageTokenResolveResponse struct {
	// StorageToken is the resolved legacy Storage token. Never log it.
	StorageToken string `json:"storageToken"`
	ProjectID    int    `json:"projectId"`
	TokenID      string `json:"tokenId"`
	UserID       string `json:"userId"`
	// ExpiresAt is the token expiration timestamp, nil for non-expiring tokens.
	ExpiresAt *string `json:"expiresAt"`
	// TokenDetail is the full token detail the resolver returns alongside the
	// legacy token — the same payload as the Storage API tokens/verify response —
	// so callers do not need a follow-up verify call (keboola/connection#7604).
	// Kept as raw JSON here because this generated-companion package must not
	// import pkg/keboola; the higher-level keboola.StorageTokenExchanger decodes
	// it into a keboola.Token. Never log it — it may carry token material.
	TokenDetail json.RawMessage `json:"tokenDetail"`
}

type ApiResolveStorageTokenRequest struct {
	ctx          context.Context
	ApiService   *AuthBridgeAPIService
	subjectToken string
	body         *AuthBridgeStorageTokenResolveRequest
}

// SubjectToken sets the programmatic Connection bearer token (kbc_at_* / kbc_pat_*)
// to resolve. An optional "Bearer " scheme prefix is accepted and normalized.
func (r ApiResolveStorageTokenRequest) SubjectToken(subjectToken string) ApiResolveStorageTokenRequest {
	r.subjectToken = subjectToken
	return r
}

// AuthBridgeStorageTokenResolveRequest sets the JSON body with the target project ID.
func (r ApiResolveStorageTokenRequest) AuthBridgeStorageTokenResolveRequest(
	body AuthBridgeStorageTokenResolveRequest,
) ApiResolveStorageTokenRequest {
	r.body = &body
	return r
}

func (r ApiResolveStorageTokenRequest) Execute() (*AuthBridgeStorageTokenResolveResponse, *http.Response, error) {
	return r.ApiService.ResolveStorageTokenExecute(r)
}

/*
ResolveStorageToken resolves a legacy Storage token for a programmatic Connection token.

Validates a programmatic Connection bearer token (kbc_at_* / kbc_pat_*) against the
given project and returns its decrypted legacy Storage token. The caller must be
Kubernetes-authenticated with the internal:auth-bridge:resolve-storage-token scope.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
	@return ApiResolveStorageTokenRequest
*/
func (a *AuthBridgeAPIService) ResolveStorageToken(ctx context.Context) ApiResolveStorageTokenRequest {
	return ApiResolveStorageTokenRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// ResolveStorageTokenExecute executes the request.
//
//	@return AuthBridgeStorageTokenResolveResponse
func (a *AuthBridgeAPIService) ResolveStorageTokenExecute(
	r ApiResolveStorageTokenRequest,
) (*AuthBridgeStorageTokenResolveResponse, *http.Response, error) {
	var returnValue *AuthBridgeStorageTokenResolveResponse

	if strings.TrimSpace(r.subjectToken) == "" {
		return returnValue, nil, &GenericOpenAPIError{error: "subjectToken is required"}
	}
	if r.body == nil {
		return returnValue, nil, &GenericOpenAPIError{error: "authBridgeStorageTokenResolveRequest is required"}
	}

	basePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthBridgeAPIService.ResolveStorageToken")
	if err != nil {
		return returnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	path := basePath + "/manage/internal/auth-bridge/resolve-storage-token"

	headerParams := map[string]string{
		"Content-Type":     "application/json",
		"Accept":           "application/json",
		HeaderSubjectToken: ensureBearerScheme(r.subjectToken),
	}

	req, err := a.client.prepareRequest(
		r.ctx, path, http.MethodPost, r.body, headerParams, url.Values{}, url.Values{}, []formFile{},
	)
	if err != nil {
		return returnValue, nil, err
	}

	httpResponse, err := a.client.callAPI(req)
	if err != nil || httpResponse == nil {
		return returnValue, httpResponse, err
	}

	body, err := io.ReadAll(httpResponse.Body)
	httpResponse.Body.Close()
	httpResponse.Body = io.NopCloser(bytes.NewBuffer(body))
	if err != nil {
		return returnValue, httpResponse, err
	}

	if httpResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  body,
			error: httpResponse.Status,
		}
		return returnValue, httpResponse, newErr
	}

	err = a.client.decode(&returnValue, body, httpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  body,
			error: err.Error(),
		}
		return returnValue, httpResponse, newErr
	}

	return returnValue, httpResponse, nil
}

// ensureBearerScheme normalizes a token to the canonical "Bearer <token>" form
// expected by the resolver's X-Subject-Token header. An existing scheme prefix
// is matched case-insensitively (HTTP auth schemes are case-insensitive, RFC
// 7235) so a lowercase "bearer …" value is not double-prefixed, and surrounding
// whitespace is trimmed.
func ensureBearerScheme(token string) string {
	token = strings.TrimSpace(token)

	const scheme = "bearer "
	if len(token) >= len(scheme) && strings.EqualFold(token[:len(scheme)], scheme) {
		return "Bearer " + strings.TrimSpace(token[len(scheme):])
	}

	return "Bearer " + token
}
