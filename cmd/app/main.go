package main

import (
	"fmt"
)

type Puzzle struct {
	name     string
	solution func() (int, int)
}

func main() {
	puzzles := []Puzzle{}

	for i, puzzle := range puzzles {
		one, two := puzzle.solution()
		fmt.Printf("%s\n  Part One: %d\n  Part Two: %d\n", puzzle.name, one, two)
		if i < len(puzzles)-1 {
			fmt.Println()
		}
	}
}
