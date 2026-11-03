// Hand-written companion to the generated client — not produced by
// openapi-generator, keep when regenerating.
package management

import "github.com/keboola/keboola-sdk-go/v2/pkg/request"

// ClientOption customizes the APIClient built by NewAPIClientWithAuth.
type ClientOption func(*clientOptions)

type clientOptions struct {
	retry *request.RetryConfig
}

// WithRetry makes the client retry a request according to the retry config, the same way the other API clients of
// the SDK do (see client.NewRetryRoundTripper). The request body is sent again on every attempt and the Auth headers
// are resolved again, so a rotated credential is used by the next attempt.
//
// It is opt-in because the Manage API also has operations that change data, and a retry of such an operation is up to
// the caller. Use it for a client that calls only read-only operations, or tune the Condition of the config.
func WithRetry(retry request.RetryConfig) ClientOption {
	return func(o *clientOptions) {
		o.retry = &retry
	}
}
