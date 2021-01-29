package board

import (
	"testing"
)

func TestBoardCreation(t *testing.T) {
	testCases := map[string]struct {
		size int
	}{
		"board of size one": {
			size: 1,
		},
		"board of size two": {
			size: 2,
		},
		"board of size five": {
			size: 5,
		},
		"board of size eight": {
			size: 8,
		},
	}
	for name, tc := range testCases {
		b := NewBoard(tc.size)
		if tc.size*tc.size != len(*b) {
			t.Errorf("%q test case failed: could not create board of size %d, created board was %v", name, tc.size, b)
		}
	}
}

func TestFlipOnce(t *testing.T) {
	b := NewBoard(1)
	Build(b, 0, 0)
	if (*b)[0].lvl == 0 {
		t.Errorf("%q test case failed: could not flip a cell on a board of size %d, board is %v", "TestFlipCell", 1, b)
	}
}

func TestFlip4Times(t *testing.T) {
	b := NewBoard(1)
	for i := 0; i < 4; i++ {
		Build(b, 0, 0)
	}
	if (*b)[0].lvl != 3 {
		t.Errorf("%q test case failed: could not flip a cell on a board of size %d, board is %v", "TestFlipCell", 1, b)
	}
}

func TestOwnerChange(t *testing.T) {
	b := NewBoard(1)
	Build(b, 0, 1)
	if (*b)[0].own != 1 {
		t.Errorf("%q test case failed: could not flip and change owner on a cell on a board of size %d, board is %v", "TestFlipCell", 1, b)
	}
}

func TestNoOwnerChangeAfterWallIsBuilt(t *testing.T) {
	b := NewBoard(1)
	for i := 0; i < 4; i++ {
		Build(b, 0, 1)
	}
	Build(b, 0, 2)
	if (*b)[0].own != 1 {
		t.Errorf("%q test case failed: should not be able to flip and change owner on a fully built cell on a board of size %d, board is %v", "TestFlipCell", 1, b)
	}
}

func TestNotEndedGame(t *testing.T) {
	b := NewBoard(1)
	if Complete(b) {
		t.Log(b)
		t.Error("Game is not finished until all cells are built at least 3 times")
	}
}

func TestEndedGame(t *testing.T) {
	b := NewBoard(1)
	Build(b, 0, 0)
	Build(b, 0, 0)
	Build(b, 0, 0)
	if !Complete(b) {
		t.Log(b)
		t.Error("Game should be finished since all cells have been built at least 3 times")
	}
}
