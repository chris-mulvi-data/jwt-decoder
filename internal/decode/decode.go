// Package decode provides functionality to decode JWT tokens.
package decode

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/chris-mulvi-data/jwt-decoder/internal/types"
)

// DecodeToken decodes a JWT token string into its header, payload, and signature components.
//
// Usage:
//
//	decodedToken, err := DecodeToken("<jwt-token-string>")
func DecodeToken(jwtString string) (types.DecodedToken, error) {

	var err error = nil
	var decodedToken types.DecodedToken

	// split the token into its components
	tokenParts, err := splitToken(jwtString)
	if err != nil {
		return decodedToken, err
	}

	// decode the header and payload from base64url encoding
	headerBytes, err := base64.RawURLEncoding.DecodeString(tokenParts.Header)
	if err != nil {
		return decodedToken, err
	}
	payloadBytes, err := base64.RawStdEncoding.DecodeString(tokenParts.Payload)
	if err != nil {
		return decodedToken, err
	}

	var decodedHeader []types.KV
	var decodedPaylod []types.KV

	// parse the header and payload JSON into ordered key-value pairs
	if decodedHeader, err = DecodeOrderedNode(headerBytes); err != nil {
		return decodedToken, err
	}
	if decodedPaylod, err = DecodeOrderedNode(payloadBytes); err != nil {
		return decodedToken, err
	}

	// construct the DecodedToken struct
	decodedToken.Header = decodedHeader
	decodedToken.Payload = decodedPaylod
	decodedToken.Signature = tokenParts.Signature

	return decodedToken, err
}

// splitToken splits a JWT token string into its header, payload, and signature components.
func splitToken(tokenString string) (types.RawToken, error) {

	tokenString = strings.Trim(tokenString, " ")
	tokenParts := strings.Split(tokenString, ".")
	var tokenPartsTyped types.RawToken
	if len(tokenParts) != 3 {
		return tokenPartsTyped, errors.New("invalid token format")
	}
	// Assign the split parts to the RawToken struct
	tokenPartsTyped.Header = tokenParts[0]    // header is always first element
	tokenPartsTyped.Payload = tokenParts[1]   // payload is always second element
	tokenPartsTyped.Signature = tokenParts[2] // signature is always third element

	return tokenPartsTyped, nil
}

// DecodeOrderedNode decodes a JSON byte slice into an ordered list of key-value pairs.
func DecodeOrderedNode(b []byte) ([]types.KV, error) {
	decoder := json.NewDecoder(bytes.NewReader(b))
	tok, err := decoder.Token()
	if err != nil {
		return nil, err
	}

	// ensure the JSON starts with an object delimiter '{'
	if d, ok := tok.(json.Delim); !ok || d != '{' {
		return nil, fmt.Errorf("expected '{'")
	}

	// iterate over the JSON object and collect key-value pairs
	var out []types.KV
	for decoder.More() {
		keyTok, err := decoder.Token()
		if err != nil {
			return nil, err
		}

		// ensure the key is a string
		key := keyTok.(string)

		// decode the corresponding value
		var val any
		if err := decoder.Decode(&val); err != nil {
			return nil, err
		}

		out = append(out, types.KV{Key: key, Value: val})
	}
	// consume the closing object delimiter '}'
	_, _ = decoder.Token()

	return out, nil
}
