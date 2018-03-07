package main

import (
	"fmt"

	"github.com/silvajonathan/GreedySearchEightQueensGo/eightqueens"
)

func main() {
	initialBoard := [8]int32{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Eight Queens problem.")
	eightqueens.Execute(initialBoard)
}
