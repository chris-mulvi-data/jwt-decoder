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

	if opts.ShouldPrompt {
		token, err = input.GetTokenFromUser("Enter JWT")
		if err != nil {
			output.PrintError(err)
			return
		}
	} else {
		token = opts.TokenString
	}

	decoded, err := decode.DecodeToken(token)
	if err != nil {
		output.PrintError(err)
		return
	}

	output.PrintItems("Header", decoded.Header)
	output.PrintItems("Payload", decoded.Payload)
	output.PrintStringWithHeading("Signature", decoded.Signature)
	fmt.Print("\n") // add a space before the next terminal prompt

}
