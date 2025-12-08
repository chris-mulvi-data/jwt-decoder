package main

import (
	"fmt"

	"github.com/chris-mulvi-data/jwt-decoder/internal/decode"
	"github.com/chris-mulvi-data/jwt-decoder/internal/input"
	"github.com/chris-mulvi-data/jwt-decoder/internal/output"
)

func main() {
	token, err := input.GetTokenFromUser("Enter JWT")
	if err != nil {
		fmt.Println(err)
	}

	decoded, err := decode.DecodeToken(token)
	if err != nil {
		fmt.Println(err)
	}

	output.PrintItems("Header", decoded.Header)
	output.PrintItems("Payload", decoded.Payload)
	output.PrintStringWithHeading("Signature", decoded.Signature)

}
