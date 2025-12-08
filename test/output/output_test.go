package output

import (
	"testing"

	"github.com/chris-mulvi-data/jwt-decoder/internal/output"
	"github.com/chris-mulvi-data/jwt-decoder/internal/types"
)

func TestPrintValuesWithColorByType(t *testing.T) {

	testItems := []any{
		[]any{"hello", "foo", "bar", "baz", 42},
		12,
	}

	output.PrintValueWithColorByType(testItems)
}

func TestPrintItms(t *testing.T) {
	tests := []types.KV{
		{Key: "someKey", Value: "some value"},
		{Key: "anotherKey", Value: 12.34},
		{Key: "more", Value: 42},
		{Key: "aSlice", Value: []any{"one", "two", "three"}},
	}

	output.PrintItems(tests)

}
