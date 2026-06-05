// Hand-written companion to the generated client — not produced by
// openapi-generator, keep when regenerating.
package management

// Names of the security schemes defined in the Manage API OpenAPI spec
// (api/openapi.yaml). Use them as keys in the ContextAPIKeys map:
//
//	ctx = context.WithValue(ctx, management.ContextAPIKeys,
//	    map[string]management.APIKey{management.AuthSchemeAPIKey: {Key: token}})
const (
	// AuthSchemeAPIKey authenticates with a Manage API token
	// sent in the HeaderManageAPIToken header.
	AuthSchemeAPIKey = "ApiKeyAuth"
	// AuthSchemeK8s authenticates with a Kubernetes service-account token
	// sent in the HeaderKubernetesAuthorization header.
	AuthSchemeK8s = "K8sAuth"
)

// HTTP header names used by the security schemes above.
const (
	HeaderManageAPIToken          = "X-KBC-ManageApiToken" //nolint:gosec // header name, not a credential
	HeaderKubernetesAuthorization = "X-Kubernetes-Authorization"
)
