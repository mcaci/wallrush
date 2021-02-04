package main

import (
	wallrush "github.com/mcaci/othello/wallrun"
)

const (
	n, l = 5, 10
)

func main() {
	wallrush.Run(n, l)
}
