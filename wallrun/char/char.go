package char

import (
	"fmt"
	"math/rand"

	"github.com/mcaci/othello/wallrun/board"
)

// Dir is a direction
type Dir uint8

// Rand gives a random direction
func Rand() Dir { return Dir(rand.Intn(4)) }

const (
	// Right direction
	Right Dir = iota
	// Left direction
	Left
	// Down direction
	Down
	// Up direction
	Up
)

// Char is a character
type Char struct {
	id  uint8
	pos int
}

// NewChar creates a character
func NewChar(id uint8, pos int) *Char { return &Char{id: id, pos: pos} }

// Pos returns the character position
func Pos(c *Char) int { return c.pos }

// Move moves the character of one step
func Move(c *Char, d Dir, side int) error {
	leg := func(*Char, int) bool { return false }
	switch d {
	case Right:
		leg = func(c *Char, side int) bool { return c.pos%side != side-1 }
	case Left:
		leg = func(c *Char, side int) bool { return c.pos%side != 0 }
	case Down:
		leg = func(c *Char, side int) bool { return c.pos/side != side-1 }
	case Up:
		leg = func(c *Char, side int) bool { return c.pos/side != 0 }
	}
	if !leg(c, side) {
		return fmt.Errorf("cannot move in direction %d from position %d", d, c.pos)
	}
	mov := func(c *Char, delta int) { c.pos += delta }
	var delta int
	switch d {
	case Right:
		delta = 1
	case Left:
		delta = -1
	case Down:
		delta = side
	case Up:
		delta = -side
	}
	mov(c, delta)
	return nil
}

// Build makes the character build a wall cell
func Build(c *Char, b *board.Board) uint8 {
	return board.Build(b, c.pos, c.id)
}
