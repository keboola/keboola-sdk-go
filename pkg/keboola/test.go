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

// RSAKeyPair holds a generated RSA key pair in both PEM-encoded and raw form.
type RSAKeyPair struct {
	PrivateKeyPEM []byte
	PublicKeyPEM  []byte
	PrivateKey    *rsa.PrivateKey
	PublicKey     *rsa.PublicKey
}

// GenerateRSAKeyPairPKCS1 creates a new RSA key pair and returns the private key in PKCS1 format.
// This is useful for compatibility with systems that expect PKCS1 format (like some Snowflake configurations).
// Returns both PEM-encoded keys and the raw key objects.
func GenerateRSAKeyPairPKCS1(t *testing.T) *RSAKeyPair {
	t.Helper()

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err, "Failed to generate RSA private key")

	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	publicKeyPKIX, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	require.NoError(t, err, "Failed to marshal public key to PKIX")

	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyPKIX,
	})

	return &RSAKeyPair{
		PrivateKeyPEM: privateKeyPEM,
		PublicKeyPEM:  publicKeyPEM,
		PrivateKey:    privateKey,
		PublicKey:     &privateKey.PublicKey,
	}
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
