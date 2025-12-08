package decode

import (
	"fmt"
	"testing"

	"github.com/chris-mulvi-data/jwt-decoder/internal/decode"
)

var token string = " eyJhbGciOiJIUzI1NiJ9.eyJSb2xlIjoiQWRtaW4iLCJJc3N1ZXIiOiJJc3N1ZXIiLCJVc2VybmFtZSI6IkphdmFJblVzZSIsImV4cCI6MTc2NTIyOTQ1MCwiaWF0IjoxNzY1MjI5NDUwfQ.HHOlS0fPp8FDhfCdjk0cujYzQjSevLRd19TiTU8JwF8"

func TestDecodeToken(t *testing.T) {
	decoded, err := decode.DecodeToken(token)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(decoded.Header)
	fmt.Println(decoded.Payload)
}
