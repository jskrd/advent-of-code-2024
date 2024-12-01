package day01_test

import (
	"strings"
	"testing"

	"github.com/jskrd/advent-of-code-2024/internal/day01"
)

const input = "3   4\n4   3\n2   5\n1   3\n3   9\n3   3"

func TestSolvePartOne(t *testing.T) {
	lines := strings.Split(input, "\n")

	expected := 11
	actual := day01.SolvePartOne(lines)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestSolvePartTwo(t *testing.T) {
	lines := strings.Split(input, "\n")

	expected := 31
	actual := day01.SolvePartTwo(lines)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
