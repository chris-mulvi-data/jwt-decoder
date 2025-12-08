package main

import (
	"fmt"

	"github.com/chris-mulvi-data/jwt-decoder/internal/input"
)

func main() {
	token, err := input.GetTokenFromUser("Enter JWT")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token)
}
