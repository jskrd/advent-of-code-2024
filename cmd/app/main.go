package main

import (
	"fmt"

	"github.com/jskrd/advent-of-code-2024/internal/day01"
	"github.com/jskrd/advent-of-code-2024/internal/day02"
)

type Puzzle struct {
	name     string
	solution func() (int, int)
}

func main() {
	puzzles := []Puzzle{
		{"Day 1: Historian Hysteria", day01.Solve},
		{"Day 2: Red-Nosed Reports", day02.Solve},
	}

	for i, puzzle := range puzzles {
		one, two := puzzle.solution()
		fmt.Printf("%s\n  Part One: %d\n  Part Two: %d\n", puzzle.name, one, two)
		if i < len(puzzles)-1 {
			fmt.Println()
		}
	}
}
