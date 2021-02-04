package board

import (
	"errors"
	"math"
)

// Cell is a board cell
type Cell struct {
	lvl, own uint8
}

func built(c *Cell) bool      { return c.lvl >= 3 }
func inc(c *Cell)             { c.lvl++ }
func assign(c *Cell, o uint8) { c.own = o }

// Board is the game board
type Board []*Cell

// NewBoard creates a board
func NewBoard(l int) *Board {
	b := make(Board, l*l)
	for i := range b {
		b[i] = new(Cell)
	}
	return &b
}

// At returns a pointer to the Cell at index i
func (b *Board) At(i int) *Cell {
	return (*b)[i]
}

// Build builds the wall cell and gives its ownership to the player who
// did the walling, after the 3rd build action the wall is built
// and no more direct changes are permitted on the cell
func Build(b *Board, i int, o uint8) uint8 {
	c := b.At(i)
	if built(c) {
		return c.lvl
	}
	inc(c)
	assign(c, o)
	Flip(b, i, o, b.Built(i))
	return c.lvl
}

// Complete returns true if the board only contains full built cells (cell.lvl == 3)
func Complete(b *Board) bool {
	for _, c := range *b {
		if !built(c) {
			return false
		}
	}
	return true
}

// Count returns a map for each player of the currently owned cells (including a count on the free ones)
func Count(b *Board) map[uint8]int {
	m := make(map[uint8]int)
	for _, c := range *b {
		m[c.own]++
	}
	return m
}

func (c Cell) Lvl() uint8                   { return c.lvl }
func (c Cell) Own() uint8                   { return c.own }
func Built(c interface{ Lvl() uint8 }) bool { return c.Lvl() >= 3 }
func Inc(c *Cell)                           { c.lvl++ }
func Assign(c *Cell, o uint8)               { c.own = o }

func Side(b *Board) int            { return int(math.Sqrt(float64(len(*b)))) }
func (b *Board) Built(i int) bool  { return (*b)[i].Lvl() >= 3 }
func (b *Board) Ground(i int) bool { return (*b)[i].Lvl() == 0 }
func (b *Board) Owned(owner uint8) func(int) bool {
	return func(i int) bool { return (*b)[i].Own() == owner }
}

type BIterator struct {
	b    *Board
	curr int
	inc  int
	limF func(int) bool
}

func NewBIT(b *Board, curr int, inc int, lim func(int) bool) *BIterator {
	return &BIterator{b: b, curr: curr, inc: inc, limF: lim}
}

var ErrEndOfLine error = errors.New("end of line reached")

func (bi *BIterator) Next() (*Cell, error) {
	if bi.limF(bi.curr) {
		return nil, ErrEndOfLine
	}
	bi.curr += bi.inc
	return (*bi.b)[bi.curr], nil
}
