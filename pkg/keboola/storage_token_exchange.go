package keboola

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

// StorageTokenExchanger exchanges Connection programmatic bearer tokens
// (kbc_at_* / kbc_pat_*) for legacy Storage tokens via Connection's internal
// auth-bridge resolver endpoint. Services use it to accept programmatic
// tokens on endpoints that authenticate with legacy Storage tokens.
//
// The resolver returns the legacy Storage token together with its full token
// detail — the same payload as the Storage API tokens/verify response
// (keboola/connection#7604) — so a single exchange call yields everything the
// usual VerifyTokenRequest would: StorageTokenExchangeResult.Token is a
// ready-to-use *Token (with Token.Token set to the legacy token), and callers
// can build a NewAuthorizedAPI from it without a follow-up verify round-trip.
//
// The exchanger authenticates itself to Connection with a management.Auth,
// typically management.NewKeboolaServiceAccountAuth — the service's projected
// Kubernetes ServiceAccount JWT mapped to the
// internal:auth-bridge:resolve-storage-token scope.
type StorageTokenExchanger struct {
	api *management.APIClient
}

// StorageTokenExchangeResult is the outcome of a successful token exchange.
type StorageTokenExchangeResult struct {
	// StorageToken is the resolved legacy Storage token. Never log it.
	StorageToken string
	ProjectID    int
	// TokenID and UserID identify the resolved token for audit logging.
	TokenID string
	UserID  string
	// ExpiresAt is the raw expiration timestamp, nil for non-expiring tokens.
	ExpiresAt *string
	// Token is the full resolved token detail (the same payload the Storage API
	// tokens/verify returns), built directly from the resolver response with
	// Token.Token set to StorageToken. Because the resolver returns the detail
	// alongside the legacy token, callers can use it directly — e.g. to build a
	// NewAuthorizedAPI — without a follow-up VerifyTokenRequest.
	Token *Token
}

// StorageTokenExchangeError describes a failed token exchange.
// The message never contains token material.
type StorageTokenExchangeError struct {
	// StatusCode is the HTTP status returned by the resolver, 0 for
	// network/transport errors or a malformed resolver response (a 200 missing
	// storageToken/tokenDetail, or a token detail that fails to decode).
	StatusCode int
	err        error
}

func (e *StorageTokenExchangeError) Error() string {
	if e.StatusCode == 0 {
		return fmt.Sprintf("storage token exchange failed: %s", e.err)
	}

	return fmt.Sprintf("storage token exchange failed: resolver returned status %d", e.StatusCode)
}

func (e *StorageTokenExchangeError) Unwrap() error {
	return e.err
}

// ClientStatusCode maps the resolver error to the HTTP status a service
// should return to its caller: 400 invalid/missing project ID -> 400,
// 401 invalid subject token -> 401, 403 token cannot access project -> 403,
// anything else (5xx, timeout, network, malformed resolver response) -> 502.
func (e *StorageTokenExchangeError) ClientStatusCode() int {
	switch e.StatusCode {
	case http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden:
		return e.StatusCode
	default:
		return http.StatusBadGateway
	}
}

// StorageTokenExchangerOption customizes the StorageTokenExchanger.
type StorageTokenExchangerOption func(c *management.Configuration)

// WithExchangerHTTPClient sets the HTTP client used for resolver calls.
func WithExchangerHTTPClient(client *http.Client) StorageTokenExchangerOption {
	return func(c *management.Configuration) {
		c.HTTPClient = client
	}
}

// NewStorageTokenExchanger creates a StorageTokenExchanger calling the
// resolver on the given Connection host (e.g. "https://connection.keboola.com").
func NewStorageTokenExchanger(
	connectionURL string,
	auth management.Auth,
	opts ...StorageTokenExchangerOption,
) (*StorageTokenExchanger, error) {
	if connectionURL == "" {
		return nil, fmt.Errorf("connection URL must not be empty")
	}
	if auth == nil {
		return nil, fmt.Errorf("auth must not be nil")
	}

	// Trim a trailing slash so the resolver path (which already starts with "/")
	// is not appended as a double slash, which can trigger redirects/routing
	// issues depending on the deployment.
	connectionURL = strings.TrimRight(connectionURL, "/")

	cfg := management.NewConfiguration()
	cfg.Servers = management.ServerConfigurations{{URL: connectionURL}}
	for _, opt := range opts {
		opt(cfg)
	}

	api, err := management.NewAPIClientWithAuth(cfg, auth)
	if err != nil {
		return nil, err
	}

	return &StorageTokenExchanger{api: api}, nil
}

// Exchange resolves a programmatic Connection bearer token (kbc_at_* /
// kbc_pat_*) to the legacy Storage token of the given project. An optional
// "Bearer " scheme prefix on the subject token is accepted.
//
// On success the resolver returns the legacy token together with its full
// token detail, so the result already carries a ready-to-use *Token — no
// follow-up VerifyTokenRequest is needed.
//
// Failures are returned as *StorageTokenExchangeError — use its
// ClientStatusCode to map the failure to an HTTP response.
func (e *StorageTokenExchanger) Exchange(
	ctx context.Context,
	subjectToken string,
	projectID int,
) (*StorageTokenExchangeResult, error) {
	if !IsProgrammaticToken(subjectToken) {
		return nil, &StorageTokenExchangeError{
			StatusCode: http.StatusUnauthorized,
			err:        fmt.Errorf("subject token is not a programmatic Connection token"),
		}
	}

	response, httpResponse, err := e.api.AuthBridgeAPI().
		ResolveStorageToken(ctx).
		SubjectToken(subjectToken).
		AuthBridgeStorageTokenResolveRequest(management.AuthBridgeStorageTokenResolveRequest{ProjectID: projectID}).
		Execute()
	if httpResponse != nil && httpResponse.Body != nil {
		_ = httpResponse.Body.Close()
	}
	if err != nil {
		statusCode := 0
		if httpResponse != nil {
			statusCode = httpResponse.StatusCode
		}

		return nil, &StorageTokenExchangeError{StatusCode: statusCode, err: err}
	}

	// A 200 that lacks either the legacy token or a non-empty token detail (e.g.
	// a Connection deploy predating keboola/connection#7604) is an our-side
	// incident mapped to 502 — never echo the response, it may carry token
	// material. An empty object/array or null detail counts as missing.
	detailFields := map[string]json.RawMessage{}
	if response.StorageToken == "" ||
		json.Unmarshal(response.TokenDetail, &detailFields) != nil ||
		len(detailFields) == 0 {
		return nil, &StorageTokenExchangeError{
			err: fmt.Errorf("resolver response did not contain a storage token and its detail"),
		}
	}

	token := &Token{}
	if err := json.Unmarshal(response.TokenDetail, token); err != nil {
		return nil, &StorageTokenExchangeError{
			err: fmt.Errorf("resolver response token detail could not be decoded: %w", err),
		}
	}
	token.Token = response.StorageToken

	return &StorageTokenExchangeResult{
		StorageToken: response.StorageToken,
		ProjectID:    response.ProjectID,
		TokenID:      response.TokenID,
		UserID:       response.UserID,
		ExpiresAt:    response.ExpiresAt,
		Token:        token,
	}, nil
}
