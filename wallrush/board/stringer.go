package board

import (
	"math"
	"strings"

	"github.com/fatih/color"
)

const (
	// https://www.unicode.org/charts/ or https://unicode-search.net/unicode-namesearch.pl?term=CIRCLED
	zero  = 0x53E3
	one   = 0x4E00
	two   = 0x4E8C
	three = 0x4E09
	def   = 0x4E10
)

func (b *Board) String() string {
	l := int(math.Sqrt(float64(len(*b))))
	sb := strings.Builder{}
	for i, c := range *b {
		sb.WriteString(c.String())
		if i%l != l-1 {
			continue
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (c *Cell) String() string {
	lvl, own := c.lvl, c.own
	f := color.RedString
	switch own {
	case 0:
		f = color.BlackString
	case 1:
		f = color.BlueString
	case 2:
		f = color.GreenString
	default:
		f = color.RedString
	}
	var r rune
	switch lvl {
	case 0:
		r = zero
	case 1:
		r = one
	case 2:
		r = two
	case 3:
		r = three
	default:
		r = def
	}
	return f("%c", r)
}
