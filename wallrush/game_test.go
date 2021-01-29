package wallrush

import (
	"testing"
)

func TestGameSetupOk(t *testing.T) {
	g, _ := NewGame(1, 1)
	if g.b == nil {
		t.Error("Board not created")
	}
	if len(g.c) != len(g.p) {
		t.Errorf("Length of players, \"%d\", and their channels, \"%d\", are not the same", len(g.p), len(g.c))
	}
}

func TestGameSetupSizeErr(t *testing.T) {
	_, err := NewGame(1, 0)
	if err == nil {
		t.Error("No errors detected with board size == 0")
	}
}

func TestGameSetup0PErr(t *testing.T) {
	_, err := NewGame(0, 1)
	if err == nil {
		t.Error("No errors detected with number of players == 0")
	}
}
