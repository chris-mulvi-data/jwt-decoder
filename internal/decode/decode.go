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

func DecodeToken(rawToken string) (types.DecodedToken, error) {

	var err error = nil
	var decodedToken types.DecodedToken

	// split the raw token into the header, payload, and signature parts
	tokenParts := strings.Split(rawToken, ".")
	if len(tokenParts) != 3 {
		return decodedToken, errors.New("invalid token format")
	}

	headerBytes, err := base64.RawURLEncoding.DecodeString(tokenParts[0])
	if err != nil {
		return decodedToken, err
	}

	payloadBytes, err := base64.RawStdEncoding.DecodeString(tokenParts[1])
	if err != nil {
		return decodedToken, err
	}

	var decodedHeader []types.KV
	var decodedPaylod []types.KV

	if decodedHeader, err = DecodeOrderedNode(headerBytes); err != nil {
		return decodedToken, err
	}

	if decodedPaylod, err = DecodeOrderedNode(payloadBytes); err != nil {
		return decodedToken, err
	}

	decodedToken.Header = decodedHeader
	decodedToken.Payload = decodedPaylod
	decodedToken.Signature = tokenParts[2]

	return decodedToken, err
}

func DecodeOrderedNode(b []byte) ([]types.KV, error) {
	decoder := json.NewDecoder(bytes.NewReader(b))
	tok, err := decoder.Token()
	if err != nil {
		return nil, err
	}

	if d, ok := tok.(json.Delim); !ok || d != '{' {
		return nil, fmt.Errorf("expected '{'")
	}

	var out []types.KV
	for decoder.More() {
		keyTok, err := decoder.Token()
		if err != nil {
			return nil, err
		}

		key := keyTok.(string)

		var val any
		if err := decoder.Decode(&val); err != nil {
			return nil, err
		}

		out = append(out, types.KV{Key: key, Value: val})
	}
	_, _ = decoder.Token() // consume the last '}'
	return out, nil
}
