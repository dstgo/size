package size

import (
	"fmt"
	"github.com/shopspring/decimal"
	"regexp"
	"slices"
	"strconv"
	"unicode"
)

type Unit = uint64

const (
	B Unit = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
	EB
)

var unitsTable = map[Unit][]string{
	EB: {"EB", "Eb", "eb", "eB", "E", "e"},
	PB: {"PB", "Pb", "pb", "pB", "P", "p"},
	TB: {"TB", "Tb", "tb", "tB", "T", "t"},
	GB: {"GB", "Gb", "gb", "gB", "G", "g"},
	MB: {"MB", "Mb", "mb", "mB", "M", "m"},
	KB: {"KB", "Kb", "kb", "kB", "K", "k"},
	B:  {"B", "b"},
}

// New returns a new size
func New(size float64, unit Unit) Size {
	if size < 0 {
		size = -1
	}

	return Size{
		Data: size,
		Unit: unit,
	}
}

// Size represents a bytes size in Unit
type Size struct {
	Data float64
	Unit Unit
}

func (s Size) Round(n int32) float64 {
	return decimal.NewFromFloat(s.Data).Round(n).InexactFloat64()
}

// StringRound returns the string representation of the size
// if n < 0, use the original data as the float representation.
func (s Size) StringRound(n int32) string {
	if _, ok := unitsTable[s.Unit]; !ok {
		return "unknown unit"
	}
	if n < 0 {
		return fmt.Sprintf(`%.2f%s`, s.Data, unitsTable[s.Unit][0])
	}
	return fmt.Sprintf(`%.2f%s`, s.Round(n), unitsTable[s.Unit][0])
}

// String By default, only three decimal places are retained
func (s Size) String() string {
	return s.StringRound(3)
}

func (s Size) To(to Unit) Size {
	return Size{
		Data: (s.Data * float64(s.Unit)) / float64(to),
		Unit: to,
	}
}

// match the string that starting with non-negative float and ending with letters
var pattern = regexp.MustCompile("^\\d+(\\.\\d+)?[a-zA-Z]*$")

// Lookup returns a new size from the given string, return false if parse failed,
// use B unit if str has no unit string.
func Lookup(str string) (Size, bool) {
	var size Size
	if len(str) == 0 || !pattern.MatchString(str) {
		return size, false
	}

	var (
		floatStr string
		unitStr  string
	)

	for i := len(str) - 1; i >= 0; i-- {
		if !unicode.IsLetter(rune(str[i])) {
			floatStr = str[:i+1]
			unitStr = str[i+1:]
			break
		}
	}

	if parseFloat, err := strconv.ParseFloat(floatStr, 64); err != nil {
		return size, false
	} else {
		size.Data = parseFloat
	}

	// use the min unit if str has no unit,
	if len(unitStr) == 0 {
		size.Unit = B
		return size, true
	}

	// find the unit
	for unit, alias := range unitsTable {
		if index := slices.Index(alias, unitStr); index != -1 {
			size.Unit = unit
			return size, true
		}
	}

	// unit not found
	return size, false
}

// LookupTo returns a new size from the given string that convert to given unit, return false if failed
func LookupTo(str string, to Unit) (Size, bool) {
	parseSize, ok := Lookup(str)
	if !ok {
		return parseSize, ok
	}
	return parseSize.To(to), true
}
