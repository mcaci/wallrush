package char

import "github.com/fatih/color"

const (
	// https://www.unicode.org/charts/ or https://unicode-search.net/unicode-namesearch.pl?term=CIRCLED
	water  = 0x328C
	ground = 0x328F
	fire   = 0x328B
)

func (c *Char) String() string {
	f := color.New(color.FgRed, color.FgYellow).SprintfFunc()
	switch c.id {
	case 1:
		f = color.BlueString
	case 2:
		f = color.GreenString
	default:
		f = color.New(color.FgRed, color.FgYellow).SprintfFunc()
	}
	var r rune
	switch c.id {
	case 1:
		r = water
	case 2:
		r = ground
	default:
		r = fire
	}
	return f("%c", r)
}
