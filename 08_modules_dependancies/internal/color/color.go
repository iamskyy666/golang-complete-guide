package color

import (
	"fmt"
	"strconv"
	"strings"
)

// Color is the color type

type Color struct {
	Value int
}

var (
	// Black is for black-color
	Black   = Color{Value: 30}
	Green   = Color{Value: 32}
	Red     = Color{Value: 31}
	Yellow  = Color{Value: 33}
	Blue    = Color{Value: 34}
	Cyan    = Color{Value: 36}
	White   = Color{Value: 37}
	Magenta = Color{Value: 35}

	Bold      = Color{Value: 1}
	Underline = Color{Value: 4}
)

// Text accepts texts and colors to print..
func Text(txt string, colors ...Color) string {
	if len(colors) == 0 {
		return txt
	}

	var codes []string
	for _, color := range colors {
		codes = append(codes, strconv.Itoa(color.Value))
	}

	return fmt.Sprintf("\033[%sm%s\033[0m",strings.Join(codes,";"),txt)

}