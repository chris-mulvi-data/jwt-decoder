package input

import (
	"fmt"
	"testing"

	"github.com/chris-mulvi-data/jwt-decoder/internal/input"
)

func TestGetTokenFromUser(t *testing.T) {
	result, err := input.GetTokenFromUser("Enter JWT")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}
