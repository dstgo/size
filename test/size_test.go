package test

import (
	"fmt"
	"github.com/dstgo/size"
	"testing"
)

func TestParseSize(t *testing.T) {
	fmt.Println(size.ParseSize("1.2MB"))
	fmt.Println(size.ParseSize("2.3b"))
	fmt.Println(size.ParseSize("9.0123TB"))

	fmt.Println(size.ParseSize("abcd"))
	fmt.Println(size.ParseSize("MB"))
	fmt.Println(size.ParseSize("KB"))
	fmt.Println(size.ParseSize("Bk"))

	fmt.Println(size.ParseSize(""))

	newSize := size.NewSize(1, size.B)
	fmt.Println(newSize.Round(0))
}

func TestParseTargetSize(t *testing.T) {
	fmt.Println(size.ParseTargetSize("1.221MB", size.KB))
	fmt.Println(size.ParseTargetSize("2.12gB", size.GB))
}

func TestSizeMeta_TosString(t *testing.T) {
	s := size.ParseTargetSize("1.221MB", size.KB)
	fmt.Println(s.String())
	size1 := size.ParseTargetSize("1231GB", size.MB)
	fmt.Println(size1.String())
}
