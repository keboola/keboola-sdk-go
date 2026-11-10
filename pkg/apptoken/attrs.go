package apptoken

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	// MaxAttrNames is the most attribute names a token carries.
	MaxAttrNames = 100
	// MaxAttrValues is the most values one attribute carries.
	MaxAttrValues = 32
	// MaxAttrValueChars is the longest value of an attribute, in characters (Unicode code points, what JSON Schema `maxLength` counts),
	// not bytes, so that every service that limits a value, whatever its language, counts it the same way.
	MaxAttrValueChars = 256
	// MaxAttrBytes is the largest `attrs` claim as raw JSON (what AttrsBytes returns). It keeps the whole signed token, with the
	// groups, under the 8 KiB that a verifier parses and that many app servers allow for one header line.
	MaxAttrBytes = 1024
	// AttrNamePattern is the shape of an attribute name: the name of a claim of the identity provider, namespaced ones included
	// (`https://example.com/department`).
	AttrNamePattern = `^[A-Za-z0-9_.:/-]{1,128}$`
)

var attrNameRegexp = regexp.MustCompile(AttrNamePattern)

// ErrInvalidAttrs means an `attrs` claim does not meet the contract. The text names the rule that was broken and never a claim
// name or a value, so it can be logged.
var ErrInvalidAttrs = errors.New("invalid attrs")

// ReservedAttrNames returns the names that cannot be attributes (compared without regard to case, see ValidAttrName). They are the
// claims of the token itself and of the OIDC ID token that apps-proxy has already interpreted: an attribute of such a name would be
// read as the verified email, a group or the subject of the user by a rule that selects by it, and so bypass the `email_verified`
// and the group-overage handling.
func ReservedAttrNames() []string {
	return []string{
		"sub", "email", "email_verified", "groups", "roles", "iss", "aud", "exp", "iat", "nbf", "nonce", "jti", "sid", "idp", "kbc",
		"azp", "acr", "amr", "auth_time", "at_hash", "c_hash", "attrs",
	}
}

// ValidAttrName reports whether name can be an attribute: it has the AttrNamePattern shape and is not reserved. The comparison with
// the reserved names ignores case, so `Email` is refused like `email`.
func ValidAttrName(name string) bool {
	return attrNameRegexp.MatchString(name) && !reservedAttrName(name)
}

func reservedAttrName(name string) bool {
	return slices.Contains(ReservedAttrNames(), strings.ToLower(name))
}

// AttrsBytes returns the size of attrs as raw JSON, the unit of MaxAttrBytes. A nil or empty map has size 0.
func AttrsBytes(attrs map[string][]string) int {
	if len(attrs) == 0 {
		return 0
	}
	raw, err := json.Marshal(attrs)
	if err != nil {
		return 0 // a map of strings always marshals
	}

	return len(raw)
}

// ValidateAttrs checks an `attrs` claim: at most MaxAttrNames names that ValidAttrName accepts, each with 1 to MaxAttrValues values
// of 1 to MaxAttrValueChars characters, valid UTF-8 and no control character, and at most MaxAttrBytes of JSON altogether. A nil or empty map is not an
// error: it is unknown. A value list may repeat a value; the signer removes duplicates, a verifier tolerates them. A verifier that
// gets an error treats every attribute as unknown, it does not use part of the claim.
func ValidateAttrs(attrs map[string][]string) error {
	if len(attrs) > MaxAttrNames {
		return invalidAttrs("too many names")
	}
	for name, values := range attrs {
		if !attrNameRegexp.MatchString(name) {
			return invalidAttrs("bad name")
		}
		if reservedAttrName(name) {
			return invalidAttrs("reserved name")
		}
		if len(values) == 0 {
			return invalidAttrs("no values")
		}
		if len(values) > MaxAttrValues {
			return invalidAttrs("too many values")
		}
		for _, value := range values {
			if err := validAttrValue(value); err != nil {
				return err
			}
		}
	}
	if AttrsBytes(attrs) > MaxAttrBytes {
		return invalidAttrs("too large")
	}

	return nil
}

func validAttrValue(value string) error {
	switch {
	case value == "":
		return invalidAttrs("empty value")
	case !utf8.ValidString(value) || strings.ContainsFunc(value, unicode.IsControl):
		return invalidAttrs("bad value")
	case utf8.RuneCountInString(value) > MaxAttrValueChars:
		return invalidAttrs("value too long")
	default:
		return nil
	}
}

func invalidAttrs(reason string) error {
	return fmt.Errorf("%w: %s", ErrInvalidAttrs, reason)
}
