package apptoken

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

// coordinateBytes is the size of a coordinate of a point of P-256.
const coordinateBytes = 32

// jwk is the part of a JSON Web Key that is used.
type jwk struct {
	Kty string `json:"kty"`
	Crv string `json:"crv"`
	Use string `json:"use"`
	Alg string `json:"alg"`
	Kid string `json:"kid"`
	X   string `json:"x"`
	Y   string `json:"y"`
}

// ParseJWKS returns the public keys of a JWKS by their key id. Only the keys of apps-proxy are accepted: EC keys on P-256, for
// signatures, with ES256. Anything else in the set is ignored, so that one unusable key (another type, curve or use, a key id that
// is not valid, a repeated key id, a point that is not on the curve) does not make the others unusable. The error is a document
// that is not a JWKS.
func ParseJWKS(body []byte) (map[string]*ecdsa.PublicKey, error) {
	var doc struct {
		Keys []jwk `json:"keys"`
	}
	if err := json.Unmarshal(body, &doc); err != nil {
		return nil, fmt.Errorf("app user token JWKS: decode: %w", err)
	}

	keys := make(map[string]*ecdsa.PublicKey, len(doc.Keys))
	for _, k := range doc.Keys {
		if k.Kty != KeyType || k.Crv != Curve || k.Use != KeyUse || k.Alg != Algorithm || !ValidKID(k.Kid) {
			continue
		}
		if _, repeated := keys[k.Kid]; repeated {
			continue
		}
		if key := k.publicKey(); key != nil {
			keys[k.Kid] = key
		}
	}

	return keys, nil
}

// publicKey returns the key, or nil when its coordinates are not a point of P-256.
func (k jwk) publicKey() *ecdsa.PublicKey {
	x, errX := base64.RawURLEncoding.DecodeString(k.X)
	y, errY := base64.RawURLEncoding.DecodeString(k.Y)
	if errX != nil || errY != nil || len(x) != coordinateBytes || len(y) != coordinateBytes {
		return nil
	}

	point := append(append([]byte{0x04}, x...), y...)
	key, err := ecdsa.ParseUncompressedPublicKey(elliptic.P256(), point)
	if err != nil {
		return nil
	}

	return key
}
