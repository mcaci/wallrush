package wallrush

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/gosuri/uilive"
	"github.com/mcaci/othello/wallrush/board"
	"github.com/mcaci/othello/wallrush/char"
)

const (
	n, l = 2, 8
)

// Run runs the game
func Run() {
	g, err := NewGame(n, l)
	if err != nil {
		log.Printf("error reported during the creation of the board: %v.", err)
		log.Fatal("exiting the game.")
	}

	var wg sync.WaitGroup
	wg.Add(int(n))
	player := func(i int) {
		Start(g, i)
		wg.Done()
	}
	for i := 0; i < int(n); i++ {
		go player(i)
	}

	clear()
	done := make(chan struct{})
	go Render(g, done)

	wg.Wait()
	Finish(g, done)
	RenderMap(g)
}

// Game is the main struct for the wallrush game
type Game struct {
	b *board.Board
	p []*char.Char
	c chan func()
}

// NewGame returns a new Game
func NewGame(n, s uint8) (*Game, error) {
	if n == 0 {
		return nil, fmt.Errorf("number of player should be greater than 0")
	}
	if s == 0 {
		return nil, fmt.Errorf("size of the board should be greater than 0")
	}
	l, a := int(s), int(s*s)
	g := Game{
		b: board.NewBoard(l),
		p: make([]*char.Char, n),
		c: make(chan func()),
	}
	for i := range g.p {
		g.p[i] = char.NewChar(uint8(i+1), rand.Intn(a))
	}
	return &g, nil
}

// Start starts the player actions.
// It stops when the board is complete.
func Start(g *Game, id int) {
	l := int(math.Sqrt(float64(len(*g.b))))
	for !board.Complete(g.b) {
		g.c <- func() { char.Build(g.p[id], g.b) }
		g.c <- func() { char.Move(g.p[id], char.Rand(), l) }
		time.Sleep(150 * time.Millisecond)
	}
}

// Finish closes the Game's channels
func Finish(g *Game, done chan<- struct{}) {
	done <- struct{}{}
	close(g.c)
}

// Render reads the events and prints the board
func Render(g *Game, done <-chan struct{}) {
	t := time.NewTicker(100 * time.Millisecond)
	writer := uilive.New()
	writer.Start()
	defer writer.Stop()
	for {
		select {
		case act := <-g.c:
			act()
		case <-t.C:
			fmt.Fprintln(writer, g)
		case <-done:
			return
		}
	}
}

// RenderMap prints the board
func RenderMap(g *Game) {
	fmt.Println(g.b)
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
