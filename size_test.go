package size

import (
	"fmt"
	"testing"
)

func TestParseSize(t *testing.T) {
	fmt.Println(ParseSize("1.2MB"))
	fmt.Println(ParseSize("2.3b"))
	fmt.Println(ParseSize("9.0123TB"))

	fmt.Println(ParseSize("abcd"))
	fmt.Println(ParseSize("MB"))
	fmt.Println(ParseSize("KB"))
	fmt.Println(ParseSize("Bk"))

	fmt.Println(ParseSize(""))
}

func TestParseTargetSize(t *testing.T) {
	fmt.Println(ParseTargetSize("1.221MB", KB))
	fmt.Println(ParseTargetSize("2.12gB", GB))
}

func TestSizeMeta_TosString(t *testing.T) {
	size := ParseTargetSize("1.221MB", KB)
	fmt.Println(size.TosString())
	size1 := ParseTargetSize("1231GB", MB)
	fmt.Println(size1.TosString())
}
