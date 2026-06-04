package keboola

import (
	"context"
	"fmt"
	"net/http"

	"github.com/keboola/keboola-sdk-go/v2/pkg/keboola/management"
)

// StorageTokenExchanger exchanges Connection programmatic bearer tokens
// (kbc_at_* / kbc_pat_*) for legacy Storage tokens via Connection's internal
// auth-bridge resolver endpoint. Services use it to accept programmatic
// tokens on endpoints that authenticate with legacy Storage tokens: exchange
// first, then run the usual VerifyTokenRequest / NewAuthorizedAPI flow with
// the resolved token.
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
}

// StorageTokenExchangeError describes a failed token exchange.
// The message never contains token material.
type StorageTokenExchangeError struct {
	// StatusCode is the HTTP status returned by the resolver, 0 for
	// network/transport errors.
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
// anything else (5xx, timeout, network) -> 502.
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

	cfg := management.NewConfiguration()
	cfg.Servers = management.ServerConfigurations{{URL: connectionURL}}
	for _, opt := range opts {
		opt(cfg)
	}

	return &StorageTokenExchanger{api: management.NewAPIClientWithAuth(cfg, auth)}, nil
}

// Exchange resolves a programmatic Connection bearer token (kbc_at_* /
// kbc_pat_*) to the legacy Storage token of the given project. An optional
// "Bearer " scheme prefix on the subject token is accepted.
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

	return &StorageTokenExchangeResult{
		StorageToken: response.StorageToken,
		ProjectID:    response.ProjectID,
		TokenID:      response.TokenID,
		UserID:       response.UserID,
		ExpiresAt:    response.ExpiresAt,
	}, nil
}
