// The output package is responsible for displaying outputs in a user's terminal
package output

import (
	"fmt"

	"github.com/chris-mulvi-data/jwt-decoder/internal/types"
)

// PrintItems iterates over a slice of key/value pairs and prints
// them to the terminal
func PrintItems(heading string, item any) error {
	fmt.Printf("\n%s%s%s:\n", Yellow, heading, Default)
	switch items := item.(type) {
	case []types.KV:
		for _, item := range items {
			fmt.Printf("\t%s: ", item.Key)
			PrintValueWithColorByType(item.Value)
			fmt.Print("\n")
		}

	case string:
		PrintValueWithColorByType(item)

	default:
		return fmt.Errorf("unsupported item type for PrintItems")
	}

	return nil
}

// PrintStringWithHeading prints a single string value with a heading
func PrintStringWithHeading(heading, value string) {
	fmt.Printf("\n%s%s%s:\n", Yellow, heading, Default)
	fmt.Print("\t")
	PrintValueWithColorByType(value)
}

// PrintValueWithColorByType will print an item to the terminal with
// color coding associated with that type.  Some types that contain
// other objects, like slices and lists, will be passed to specific
// functions to handle the separation of the nested data.  This function
// is then called from those functions to act on the contained objects.
func PrintValueWithColorByType(value any) {

	fmt.Print("\t") // tab indent before output

	switch val := value.(type) {
	case string:
		fmt.Printf("%s%s%s", Cyan, val, Default)

	case float64:
		// NOTE: float gets printed as an int since this is probably an epoch
		//       timestamp when found in a JWT may have to rethink this at some
		//       point to pre-process certain fields into the right context
		fmt.Printf("%s%d%s", Magenta, int(val), Default)

	case float32:
		// NOTE: float gets printed as an int since this is probably an epoch
		//       timestamp when found in a JWT may have to rethink this at some
		//       point to pre-process certain fields into the right context
		fmt.Printf("%s%d%s", Magenta, int(val), Default)

	case int:
		fmt.Printf("%s%d%s", Magenta, val, Default)

	case []any:
		PrintSliceItems(val)

	default:
		fmt.Printf("%v\n", val)

	}

}

// PrintSliceItems iterates over a slice and calls the PrintValuesWithColorByType function
// to print the items with the correct color coding for that type
func PrintSliceItems(items []any) {
	for i, item := range items {
		if i > 0 {
			fmt.Print("\t") // tab after first item for alignment
		}
		// fmt.Print("\t") // indent before each slice item
		PrintValueWithColorByType(item)
		fmt.Print("\n") // move to next line
	}
}

// PrintError prints a nice and pretty version of an error message
func PrintError(err error) {

	fmt.Printf("%sERROR:%s %s%s%s\n", Red, Default, Red, err, Default)
}
