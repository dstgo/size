package size

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
	"strings"
)

type Size = uint64

const (
	B  Size = 1
	KB      = B * 1024
	MB      = KB * 1024
	GB      = MB * 1024
	TB      = GB * 1024
	PB      = TB * 1024
)

var UnitMap = map[Size][]string{
	PB: {"PB", "Pb", "pb", "pB", "P", "p"},
	TB: {"TB", "Tb", "tb", "tB", "T", "t"},
	GB: {"GB", "Gb", "gb", "gB", "G", "g"},
	MB: {"MB", "Mb", "mb", "mB", "M", "m"},
	KB: {"KB", "Kb", "kb", "kB", "K", "k"},
	B:  {"B", "b"},
}

type SizeMeta struct {
	Data float64
	Unit Size
}

func NewSize(size float64, unit Size) SizeMeta {
	return SizeMeta{
		Data: size,
		Unit: unit,
	}
}

func (s SizeMeta) Round(n int32) float64 {
	return decimal.NewFromFloat(s.Data).Round(n).InexactFloat64()
}

func (s SizeMeta) String() string {
	if _, ok := UnitMap[s.Unit]; !ok {
		return ""
	}
	return fmt.Sprintf(`%.2f%s`, decimal.NewFromFloat(s.Data).Round(2).InexactFloat64(), UnitMap[s.Unit][0])
}

// ParseSize
// @Date 2023-06-25 16:37:42
// @Param str string
// @Return SizeMeta
// @Return error
// @Description: 将字符串转换成对应的size描述
func ParseSize(str string) SizeMeta {
	var size SizeMeta
	if len(str) == 0 {
		return size
	}

	for unit, aliasSlice := range UnitMap {
		for _, alias := range aliasSlice {
			if before, found := strings.CutSuffix(str, alias); found {
				if float, err := strconv.ParseFloat(before, 64); err == nil {
					return SizeMeta{
						Data: float,
						Unit: unit,
					}
				}
			}
		}
	}

	return size
}

// ParseTargetSize
// @Date 2023-06-25 16:43:06
// @Param str string
// @Param size Size
// @Return SizeMeta
// @Return error
// @Description: 转换成指定的大小
func ParseTargetSize(str string, size Size) SizeMeta {
	parseSize := ParseSize(str)

	bytes := parseSize.Data * float64(parseSize.Unit)

	data := bytes / float64(size)

	return SizeMeta{
		Data: data,
		Unit: size,
	}
}
