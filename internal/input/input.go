// The input package is responsible for handling inputs from the user
package input

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/chris-mulvi-data/jwt-decoder/internal/types"
)

// ParseArgs gets the arguments supplied by the CLI
func ParseArgs() (types.Options, error) {
	var opts types.Options
	var err error = nil

	// check if the user supplied any arguments
	switch len(os.Args) {
	case 1:
		opts.ShouldPrompt = true
	case 2:
		opts.ShouldPrompt = false
		opts.TokenString = os.Args[1]
	default:
		err = errors.New("to many arguments supplied")
	}

	return opts, err
}

// GetTokenFromUser gets the token string from the user
func GetTokenFromUser(prompt string) (string, error) {
	var tokenString string

	fmt.Printf("%s: ", prompt)

	fmt.Scanf("%s", &tokenString)
	// clean the string to not have any quotes
	tokenString = strings.ReplaceAll(tokenString, "\"", "")

	return tokenString, nil
}
