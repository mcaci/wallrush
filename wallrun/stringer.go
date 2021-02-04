package wallrush

import (
	"fmt"
	"math"
	"strings"

	"github.com/fatih/color"
	"github.com/mcaci/othello/wallrun/board"
	"github.com/mcaci/othello/wallrun/char"
)

func (g *Game) String() string {
	l := int(math.Sqrt(float64(len(*g.b))))
	// Board
	sb := strings.Builder{}
	for i, c := range *g.b {
		var s = c.String()
		for _, pl := range g.p {
			if char.Pos(pl) != i {
				continue
			}
			s = pl.String()
		}
		sb.WriteString(s)
		if i%l != l-1 {
			continue
		}
		sb.WriteRune('\n')
	}
	sb.WriteRune('\n')
	// Player info
	m := board.Count(g.b)
	for k, v := range m {
		col := color.BlackString
		player := fmt.Sprintf("Player %d: %d\n", k, v)
		switch k {
		case 0:
			player = fmt.Sprintf("Free blocks: %d\n", v)
		case 1:
			col = color.BlueString
		case 2:
			col = color.GreenString
		default:
			col = color.New(color.FgRed, color.FgYellow).SprintfFunc()
		}
		sb.WriteString(col(player))
	}
	return sb.String()
}
