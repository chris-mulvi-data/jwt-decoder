package main

import (
	"fmt"

	"github.com/chris-mulvi-data/jwt-decoder/internal/decode"
	"github.com/chris-mulvi-data/jwt-decoder/internal/input"
	"github.com/chris-mulvi-data/jwt-decoder/internal/output"
	"github.com/chris-mulvi-data/jwt-decoder/internal/types"
)

func main() {
	var token string
	var err error
	var opts types.Options

	opts, err = input.ParseArgs()
	if err != nil {
		output.PrintError(err)
		return
	}

	// get the token string either from user prompt or command line argument
	if opts.ShouldPrompt {
		token, err = input.GetTokenFromUser("Enter JWT")
		if err != nil {
			output.PrintError(err)
			return
		}
	} else {
		token = opts.TokenString
	}

	// decode the token
	decoded, err := decode.DecodeToken(token)
	if err != nil {
		output.PrintError(err)
		return
	}

	// print the decoded token components
	output.PrintItems("Header", decoded.Header)
	output.PrintItems("Payload", decoded.Payload)
	output.PrintItems("Signature", decoded.Signature)
	fmt.Println() // add a space before the next terminal prompt

}
