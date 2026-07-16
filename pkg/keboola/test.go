package keboola

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"testing"
	"time"

	"github.com/keboola/go-utils/pkg/testproject"
	"github.com/stretchr/testify/require"

	"github.com/keboola/keboola-sdk-go/v2/pkg/client"
)

// GenerateRSAPublicKeyPEM generates a new RSA key pair and returns the public key
// as a PEM-encoded PKIX string, for Snowflake workspace keypair login in tests.
func GenerateRSAPublicKeyPEM(t *testing.T) string {
	t.Helper()

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err, "Failed to generate RSA private key")

	publicKeyPKIX, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	require.NoError(t, err, "Failed to marshal public key to PKIX")

	return string(pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyPKIX,
	}))
}

func APIClientForRandomProject(t *testing.T, ctx context.Context, opts ...testproject.Option) (*testproject.Project, *AuthorizedAPI) {
	t.Helper()

	project, err := testproject.GetTestProjectForTest(t, opts...)
	if err != nil {
		t.Fatal(err)
	}

	c := client.NewTestClient()

	publicAPI, err := NewPublicAPI(ctx, project.StorageAPIHost(), WithClient(&c))
	if err != nil {
		t.Fatal(err)
	}

	api := publicAPI.NewAuthorizedAPI(project.StorageAPIToken(), 1*time.Minute)

	return project, api
}

func APIClientForAnEmptyProject(t *testing.T, ctx context.Context, opts ...testproject.Option) (*testproject.Project, *AuthorizedAPI) {
	t.Helper()

	project, api := APIClientForRandomProject(t, ctx, opts...)
	_, err := api.CleanProjectRequest().Send(ctx)
	if err != nil {
		t.Fatalf(`cannot clean project "%d": %s`, project.ID(), err)
	}
	return project, api
}
