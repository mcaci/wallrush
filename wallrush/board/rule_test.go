package board

import "testing"

func TestOthelloFlippingScenarios(t *testing.T) {
	for name, tc := range testcases() {
		bi := NewBIT(&tc.b, tc.start, tc.d.step, tc.d.limit)
		if tc.flipOk != canFlip(bi, tc.owner, tc.b.Built(tc.start)) {
			t.Errorf("%q, error", name)
		}
	}
}

func TestOthelloFlippingCell(t *testing.T) {
	for name, tc := range testcases() {
		Flip(&tc.b, tc.start, tc.owner, tc.b.Built(tc.start))
		for _, idx := range tc.indexes {
			if tc.b[idx].Own() != tc.owner {
				t.Errorf("%q, error", name)
			}
		}
	}
}

func TestOthelloCellNotToFlip(t *testing.T) {
	for name, tc := range testcases() {
		Flip(&tc.b, tc.start, tc.owner, tc.b.Built(tc.start))
		for _, idx := range tc.indexesNot {
			if tc.b[idx].Own() == tc.owner {
				t.Errorf("%q, error", name)
			}
		}
	}
}

func testcases() map[string]struct {
	b          Board
	start      int
	d          dir
	owner      uint8
	flipOk     bool
	indexes    []int
	indexesNot []int
} {
	const testBoardSize = 4
	dirs := allDirs(testBoardSize)
	c0 := &Cell{lvl: 3, own: 0}
	c1 := &Cell{lvl: 1, own: 1}
	c1Full := &Cell{lvl: 3, own: 1}
	testBoard := func(cs ...*Cell) Board {
		b := NewBoard(testBoardSize)
		for i, c := range cs {
			(*b)[i] = &Cell{lvl: c.lvl, own: c.own}
		}
		return *b
	}

	return map[string]struct {
		b          Board
		start      int
		d          dir
		owner      uint8
		flipOk     bool
		indexes    []int
		indexesNot []int
	}{
		"flip ok on horizontal line": {
			b:      testBoard(c0, c1, c1, &Cell{lvl: 1, own: 0}),
			d:      dirs["e"],
			flipOk: true,
		},
		"flip ok on horizontal line, start from right": {
			b:          testBoard(c1, c0, c0, c1Full),
			start:      testBoardSize - 1,
			d:          dirs["w"],
			owner:      1,
			flipOk:     true,
			indexes:    []int{1, 2},
			indexesNot: []int{5, 6, 11, 13},
		},
		"flip not ok on horizontal line: no owned cell at the end of the line": {
			b: testBoard(c0, c1, c1, c1),
			d: dirs["e"],
		},
		"flip not ok on horizontal line, start from right: no owned cell at the end of the line": {
			b:     testBoard(c1, c1, c1, c0),
			start: testBoardSize - 1,
			d:     dirs["w"],
		},
		"flip not ok on horizontal line: possible flip section is not longer than 1": {
			b:     testBoard(c0, c1, c1, c1),
			d:     dirs["e"],
			owner: uint8(1),
		},
		"flip not ok on horizontal line, start from right: possible flip section is not longer than 1": {
			b:     testBoard(c0, c1, c1, c1),
			owner: uint8(1),
			start: testBoardSize - 1,
			d:     dirs["w"],
		},
		"flip not ok on horizontal line: empty wall in between": {
			b: testBoard(c0, c1, &Cell{lvl: 0, own: 1}, c0),
			d: dirs["e"],
		},
		"flip not ok on horizontal line, start from right: empty wall in between": {
			b:     testBoard(c0, c1, &Cell{lvl: 0, own: 1}, c0),
			start: testBoardSize - 1,
			d:     dirs["w"],
		},
		"flip not ok on horizontal line: wall not complete yet": {
			b: testBoard(c1),
			d: dirs["e"],
		},
		"flip not ok on horizontal line: wall not complete yet, full line": {
			b:     testBoard(c1, c0, c0, c1),
			owner: uint8(1),
			d:     dirs["w"],
		},
		"flip ok: vertical going up": {
			b:      testBoard(c0, c0, c0, c0, c1, c0, c0, c0, c1, c0, c0, c0, c0, c0, c0, c0),
			d:      dirs["s"],
			flipOk: true,
		},
		"flip ok: vertical going down": {
			b:      testBoard(c0, c0, c0, c0, c1, c0, c0, c0, c1, c0, c0, c0, c0, c0, c0, c0),
			start:  testBoardSize * (testBoardSize - 1),
			d:      dirs["n"],
			flipOk: true,
		},
		"flip not ok: vertical going down": {
			b:     testBoard(c0, c0, c0, c0, &Cell{lvl: 0, own: 1}, c0, c0, c0, c1, c0, c0, c0, c0, c0, c0, c0),
			start: testBoardSize * (testBoardSize - 1),
			d:     dirs["n"],
		},
		"flip ok: diagonal going southeast": {
			b:          testBoard(c1Full, c0, c0, c0, c0, &Cell{lvl: 3, own: 0}, c0, c0, c0, c0, &Cell{lvl: 3, own: 0}, c0, c0, c0, c0, c1),
			d:          dirs["se"],
			owner:      1,
			flipOk:     true,
			indexes:    []int{5, 10},
			indexesNot: []int{1, 2, 3, 4, 6, 7},
		},
		"flip ok: diagonal going southeast, start from cell with i=1": {
			b:      testBoard(c0, c0, c0, c0, c0, c0, c1, c0, c0, c0, c0, c0, c0, c0, c0, c0),
			start:  1,
			d:      dirs["se"],
			flipOk: true,
		},
		"flip ok: diagonal going northwest, start from cell with i=11": {
			b:      testBoard(c0, c0, c0, c0, c0, c0, c1, c0, c0, c0, c0, c0, c0, c0, c0, c0),
			start:  11,
			d:      dirs["nw"],
			flipOk: true,
		},
		"flip ok: diagonal going southwest": {
			b:      testBoard(c0, c0, c0, c0, c0, c0, c1, c0, c0, c1, c0, c0, c0, c0, c0, c0),
			start:  3,
			d:      dirs["sw"],
			flipOk: true,
		},
		"flip ok: diagonal going northeast": {
			b:      testBoard(c0, c0, c0, c0, c0, c0, c1, c0, c0, c1, c0, c0, c0, c0, c0, c0),
			start:  12,
			d:      dirs["ne"],
			flipOk: true,
		},
	}
}
