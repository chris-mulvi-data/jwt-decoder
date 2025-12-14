// Package types defines the custom types used across the application
package types

// Options are the CLI options parsed from user input
type Options struct {
	TokenString  string
	ShouldPrompt bool
}

// KV is a structured key value pair
type KV struct {
	Key   string
	Value any
}

type DecodedToken struct {
	Header    []KV
	Payload   []KV
	Signature string
}
