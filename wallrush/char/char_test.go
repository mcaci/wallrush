package char

import (
	"testing"
)

func TestLegalMove(t *testing.T) {
	testcases := map[string]struct {
		pos       int
		side      int
		direction Dir
		expected  int
	}{
		"move right": {
			pos:       0,
			side:      2,
			direction: Right,
			expected:  1,
		},
		"move left": {
			pos:       1,
			side:      2,
			direction: Left,
			expected:  0,
		},
		"move down": {
			pos:       0,
			side:      2,
			direction: Down,
			expected:  2,
		},
		"move up": {
			pos:       3,
			side:      2,
			direction: Up,
			expected:  1,
		},
		"move up with side > 2": {
			pos:       12,
			side:      5,
			direction: Up,
			expected:  7,
		},
	}
	for name, tc := range testcases {
		c := NewChar(0, tc.pos)
		Move(c, tc.direction, tc.side)
		if Pos(c) != tc.expected {
			t.Errorf("%q: expecting next position to be %d but was %d", name, tc.expected, Pos(c))
		}
	}
}

func TestIllegalMove(t *testing.T) {
	testcases := map[string]struct {
		pos       int
		side      int
		direction Dir
	}{
		"move right first row": {
			pos:       1,
			side:      2,
			direction: Right,
		},
		"move right second row": {
			pos:       3,
			side:      2,
			direction: Right,
		},
		"move left first row": {
			pos:       0,
			side:      2,
			direction: Left,
		},
		"move down last row": {
			pos:       2,
			side:      2,
			direction: Down,
		},
		"move up first row": {
			pos:       1,
			side:      2,
			direction: Up,
		},
	}
	for name, tc := range testcases {
		before := tc.pos
		c := NewChar(0, tc.pos)
		if err := Move(c, tc.direction, tc.side); err == nil {
			t.Errorf("%q: expecting error for moving in direction %d from position %d", name, tc.direction, before)
		}
	}
}
