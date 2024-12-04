package day04

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := "..X...\n.SAMX.\n.A..A.\nXMAS.S\n.X...."

	expected := [][]rune{
		{'.', '.', 'X', '.', '.', '.'},
		{'.', 'S', 'A', 'M', 'X', '.'},
		{'.', 'A', '.', '.', 'A', '.'},
		{'X', 'M', 'A', 'S', '.', 'S'},
		{'.', 'X', '.', '.', '.', '.'},
	}
	actual := parseInput(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestSolvePartOne(t *testing.T) {
	grid := [][]rune{
		{'.', '.', '.', '.', 'X', 'X', 'M', 'A', 'S', '.'},
		{'.', 'S', 'A', 'M', 'X', 'M', 'S', '.', '.', '.'},
		{'.', '.', '.', 'S', '.', '.', 'A', '.', '.', '.'},
		{'.', '.', 'A', '.', 'A', '.', 'M', 'S', '.', 'X'},
		{'X', 'M', 'A', 'S', 'A', 'M', 'X', '.', 'M', 'M'},
		{'X', '.', '.', '.', '.', '.', 'X', 'A', '.', 'A'},
		{'S', '.', 'S', '.', 'S', '.', 'S', '.', 'S', 'S'},
		{'.', 'A', '.', 'A', '.', 'A', '.', 'A', '.', 'A'},
		{'.', '.', 'M', '.', 'M', '.', 'M', '.', 'M', 'M'},
		{'.', 'X', '.', 'X', '.', 'X', 'M', 'A', 'S', 'X'},
	}

	expected := 18
	actual := solvePartOne(grid)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestSolvePartTwo(t *testing.T) {
	grid := [][]rune{
		{'.', 'M', '.', 'S', '.', '.', 'M', 'A', 'S', '.'},
		{'.', 'S', 'A', 'M', '.', 'M', 'S', 'M', 'S', '.'},
		{'.', 'M', '.', 'S', '.', 'M', 'A', 'A', '.', '.'},
		{'.', '.', 'A', '.', 'A', 'S', 'M', 'S', 'M', '.'},
		{'.', 'M', 'A', 'S', 'A', 'M', '.', '.', 'M', 'M'},
		{'.', '.', '.', '.', '.', '.', '.', 'A', '.', 'A'},
		{'S', '.', 'S', '.', 'S', '.', 'S', '.', 'S', 'S'},
		{'.', 'A', '.', 'A', '.', 'A', '.', 'A', '.', 'A'},
		{'M', '.', 'M', '.', 'M', '.', 'M', '.', 'M', 'M'},
		{'.', '.', '.', '.', '.', '.', 'M', 'A', 'S', '.'},
	}

	expected := 9
	actual := solvePartTwo(grid)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestSearchUp(t *testing.T) {
	grid := [][]rune{
		{'.', '.', '.', '.', 'X', 'X', 'M', 'A', 'S', '.'},
		{'.', 'S', 'A', 'M', 'X', 'M', 'S', '.', '.', '.'},
		{'.', '.', '.', 'S', '.', '.', 'A', '.', '.', '.'},
		{'.', '.', 'A', '.', 'A', '.', 'M', 'S', '.', 'X'},
		{'X', 'M', 'A', 'S', 'A', 'M', 'X', '.', 'M', 'M'},
		{'X', '.', '.', '.', '.', '.', 'X', 'A', '.', 'A'},
		{'S', '.', 'S', '.', 'S', '.', 'S', '.', 'S', 'S'},
		{'.', 'A', '.', 'A', '.', 'A', '.', 'A', '.', 'A'},
		{'.', '.', 'M', '.', 'M', '.', 'M', '.', 'M', 'M'},
		{'.', 'X', '.', 'X', '.', 'X', 'M', 'A', 'S', 'X'},
	}

	tests := []struct {
		row, col int
		expected bool
	}{
		{0, 4, false},
		{0, 5, false},
		{1, 4, false},
		{3, 9, false},
		{4, 0, false},
		{4, 6, true},
		{5, 0, false},
		{5, 6, false},
		{9, 1, false},
		{9, 3, false},
		{9, 5, false},
		{9, 9, true},
	}

	for _, test := range tests {
		actual := searchUp(&grid, test.row, test.col)
		if test.expected != actual {
			t.Errorf("Expected %t but got %t", test.expected, actual)
		}
	}
}

func TestSearchUpRight(t *testing.T) {
	grid := [][]rune{
		{'.', '.', '.', '.', 'X', 'X', 'M', 'A', 'S', '.'},
		{'.', 'S', 'A', 'M', 'X', 'M', 'S', '.', '.', '.'},
		{'.', '.', '.', 'S', '.', '.', 'A', '.', '.', '.'},
		{'.', '.', 'A', '.', 'A', '.', 'M', 'S', '.', 'X'},
		{'X', 'M', 'A', 'S', 'A', 'M', 'X', '.', 'M', 'M'},
		{'X', '.', '.', '.', '.', '.', 'X', 'A', '.', 'A'},
		{'S', '.', 'S', '.', 'S', '.', 'S', '.', 'S', 'S'},
		{'.', 'A', '.', 'A', '.', 'A', '.', 'A', '.', 'A'},
		{'.', '.', 'M', '.', 'M', '.', 'M', '.', 'M', 'M'},
		{'.', 'X', '.', 'X', '.', 'X', 'M', 'A', 'S', 'X'},
	}

	tests := []struct {
		row, col int
		expected bool
	}{
		{0, 4, false},
		{0, 5, false},
		{1, 4, false},
		{3, 9, false},
		{4, 0, false},
		{4, 6, false},
		{5, 0, true},
		{5, 6, false},
		{9, 1, true},
		{9, 3, true},
		{9, 5, true},
		{9, 9, false},
	}

	for _, test := range tests {
		actual := searchUpRight(&grid, test.row, test.col)
		if test.expected != actual {
			t.Errorf("Expected %t but got %t", test.expected, actual)
		}
	}
}

func TestSearchRight(t *testing.T) {
	grid := [][]rune{
		{'.', '.', '.', '.', 'X', 'X', 'M', 'A', 'S', '.'},
		{'.', 'S', 'A', 'M', 'X', 'M', 'S', '.', '.', '.'},
		{'.', '.', '.', 'S', '.', '.', 'A', '.', '.', '.'},
		{'.', '.', 'A', '.', 'A', '.', 'M', 'S', '.', 'X'},
		{'X', 'M', 'A', 'S', 'A', 'M', 'X', '.', 'M', 'M'},
		{'X', '.', '.', '.', '.', '.', 'X', 'A', '.', 'A'},
		{'S', '.', 'S', '.', 'S', '.', 'S', '.', 'S', 'S'},
		{'.', 'A', '.', 'A', '.', 'A', '.', 'A', '.', 'A'},
		{'.', '.', 'M', '.', 'M', '.', 'M', '.', 'M', 'M'},
		{'.', 'X', '.', 'X', '.', 'X', 'M', 'A', 'S', 'X'},
	}

	tests := []struct {
		row, col int
		expected bool
	}{
		{0, 4, false},
		{0, 5, true},
		{1, 4, false},
		{3, 9, false},
		{4, 0, true},
		{4, 6, false},
		{5, 0, false},
		{5, 6, false},
		{9, 1, false},
		{9, 3, false},
		{9, 5, true},
		{9, 9, false},
	}

	for _, test := range tests {
		actual := searchRight(&grid, test.row, test.col)
		if test.expected != actual {
			t.Errorf("Expected %t but got %t", test.expected, actual)
		}
	}
}

func TestSearchDownRight(t *testing.T) {
	grid := [][]rune{
		{'.', '.', '.', '.', 'X', 'X', 'M', 'A', 'S', '.'},
		{'.', 'S', 'A', 'M', 'X', 'M', 'S', '.', '.', '.'},
		{'.', '.', '.', 'S', '.', '.', 'A', '.', '.', '.'},
		{'.', '.', 'A', '.', 'A', '.', 'M', 'S', '.', 'X'},
		{'X', 'M', 'A', 'S', 'A', 'M', 'X', '.', 'M', 'M'},
		{'X', '.', '.', '.', '.', '.', 'X', 'A', '.', 'A'},
		{'S', '.', 'S', '.', 'S', '.', 'S', '.', 'S', 'S'},
		{'.', 'A', '.', 'A', '.', 'A', '.', 'A', '.', 'A'},
		{'.', '.', 'M', '.', 'M', '.', 'M', '.', 'M', 'M'},
		{'.', 'X', '.', 'X', '.', 'X', 'M', 'A', 'S', 'X'},
	}

	tests := []struct {
		row, col int
		expected bool
	}{
		{0, 4, true},
		{0, 5, false},
		{1, 4, false},
		{3, 9, false},
		{4, 0, false},
		{4, 6, false},
		{5, 0, false},
		{5, 6, false},
		{9, 1, false},
		{9, 3, false},
		{9, 5, false},
		{9, 9, false},
	}

	for _, test := range tests {
		actual := searchDownRight(&grid, test.row, test.col)
		if test.expected != actual {
			t.Errorf("Expected %t but got %t", test.expected, actual)
		}
	}
}

func TestSearchDown(t *testing.T) {
	grid := [][]rune{
		{'.', '.', '.', '.', 'X', 'X', 'M', 'A', 'S', '.'},
		{'.', 'S', 'A', 'M', 'X', 'M', 'S', '.', '.', '.'},
		{'.', '.', '.', 'S', '.', '.', 'A', '.', '.', '.'},
		{'.', '.', 'A', '.', 'A', '.', 'M', 'S', '.', 'X'},
		{'X', 'M', 'A', 'S', 'A', 'M', 'X', '.', 'M', 'M'},
		{'X', '.', '.', '.', '.', '.', 'X', 'A', '.', 'A'},
		{'S', '.', 'S', '.', 'S', '.', 'S', '.', 'S', 'S'},
		{'.', 'A', '.', 'A', '.', 'A', '.', 'A', '.', 'A'},
		{'.', '.', 'M', '.', 'M', '.', 'M', '.', 'M', 'M'},
		{'.', 'X', '.', 'X', '.', 'X', 'M', 'A', 'S', 'X'},
	}

	tests := []struct {
		row, col int
		expected bool
	}{
		{0, 4, false},
		{0, 5, false},
		{1, 4, false},
		{3, 9, true},
		{4, 0, false},
		{4, 6, false},
		{5, 0, false},
		{5, 6, false},
		{9, 1, false},
		{9, 3, false},
		{9, 5, false},
		{9, 9, false},
	}

	for _, test := range tests {
		actual := searchDown(&grid, test.row, test.col)
		if test.expected != actual {
			t.Errorf("Expected %t but got %t", test.expected, actual)
		}
	}
}

func TestSearchDownLeft(t *testing.T) {
	grid := [][]rune{
		{'.', '.', '.', '.', 'X', 'X', 'M', 'A', 'S', '.'},
		{'.', 'S', 'A', 'M', 'X', 'M', 'S', '.', '.', '.'},
		{'.', '.', '.', 'S', '.', '.', 'A', '.', '.', '.'},
		{'.', '.', 'A', '.', 'A', '.', 'M', 'S', '.', 'X'},
		{'X', 'M', 'A', 'S', 'A', 'M', 'X', '.', 'M', 'M'},
		{'X', '.', '.', '.', '.', '.', 'X', 'A', '.', 'A'},
		{'S', '.', 'S', '.', 'S', '.', 'S', '.', 'S', 'S'},
		{'.', 'A', '.', 'A', '.', 'A', '.', 'A', '.', 'A'},
		{'.', '.', 'M', '.', 'M', '.', 'M', '.', 'M', 'M'},
		{'.', 'X', '.', 'X', '.', 'X', 'M', 'A', 'S', 'X'},
	}

	tests := []struct {
		row, col int
		expected bool
	}{
		{0, 4, false},
		{0, 5, false},
		{1, 4, false},
		{3, 9, true},
		{4, 0, false},
		{4, 6, false},
		{5, 0, false},
		{5, 6, false},
		{9, 1, false},
		{9, 3, false},
		{9, 5, false},
		{9, 9, false},
	}

	for _, test := range tests {
		actual := searchDownLeft(&grid, test.row, test.col)
		if test.expected != actual {
			t.Errorf("Expected %t but got %t", test.expected, actual)
		}
	}
}

func TestSearchLeft(t *testing.T) {
	grid := [][]rune{
		{'.', '.', '.', '.', 'X', 'X', 'M', 'A', 'S', '.'},
		{'.', 'S', 'A', 'M', 'X', 'M', 'S', '.', '.', '.'},
		{'.', '.', '.', 'S', '.', '.', 'A', '.', '.', '.'},
		{'.', '.', 'A', '.', 'A', '.', 'M', 'S', '.', 'X'},
		{'X', 'M', 'A', 'S', 'A', 'M', 'X', '.', 'M', 'M'},
		{'X', '.', '.', '.', '.', '.', 'X', 'A', '.', 'A'},
		{'S', '.', 'S', '.', 'S', '.', 'S', '.', 'S', 'S'},
		{'.', 'A', '.', 'A', '.', 'A', '.', 'A', '.', 'A'},
		{'.', '.', 'M', '.', 'M', '.', 'M', '.', 'M', 'M'},
		{'.', 'X', '.', 'X', '.', 'X', 'M', 'A', 'S', 'X'},
	}

	tests := []struct {
		row, col int
		expected bool
	}{
		{0, 4, false},
		{0, 5, false},
		{1, 4, true},
		{3, 9, false},
		{4, 0, false},
		{4, 6, true},
		{5, 0, false},
		{5, 6, false},
		{9, 1, false},
		{9, 3, false},
		{9, 5, false},
		{9, 9, false},
	}

	for _, test := range tests {
		actual := searchLeft(&grid, test.row, test.col)
		if test.expected != actual {
			t.Errorf("Expected %t but got %t", test.expected, actual)
		}
	}
}

func TestSearchUpLeft(t *testing.T) {
	grid := [][]rune{
		{'.', '.', '.', '.', 'X', 'X', 'M', 'A', 'S', '.'},
		{'.', 'S', 'A', 'M', 'X', 'M', 'S', '.', '.', '.'},
		{'.', '.', '.', 'S', '.', '.', 'A', '.', '.', '.'},
		{'.', '.', 'A', '.', 'A', '.', 'M', 'S', '.', 'X'},
		{'X', 'M', 'A', 'S', 'A', 'M', 'X', '.', 'M', 'M'},
		{'X', '.', '.', '.', '.', '.', 'X', 'A', '.', 'A'},
		{'S', '.', 'S', '.', 'S', '.', 'S', '.', 'S', 'S'},
		{'.', 'A', '.', 'A', '.', 'A', '.', 'A', '.', 'A'},
		{'.', '.', 'M', '.', 'M', '.', 'M', '.', 'M', 'M'},
		{'.', 'X', '.', 'X', '.', 'X', 'M', 'A', 'S', 'X'},
	}

	tests := []struct {
		row, col int
		expected bool
	}{
		{0, 4, false},
		{0, 5, false},
		{1, 4, false},
		{3, 9, false},
		{4, 0, false},
		{4, 6, false},
		{5, 0, false},
		{5, 6, true},
		{9, 1, false},
		{9, 3, true},
		{9, 5, true},
		{9, 9, true},
	}

	for _, test := range tests {
		actual := searchUpLeft(&grid, test.row, test.col)
		if test.expected != actual {
			t.Errorf("Expected %t but got %t", test.expected, actual)
		}
	}
}

func TestSearchX(t *testing.T) {
	grid := [][]rune{
		{'.', 'M', '.', 'S', '.', '.', 'M', 'A', 'S', '.'},
		{'.', 'S', 'A', 'M', '.', 'M', 'S', 'M', 'S', '.'},
		{'.', 'M', '.', 'S', '.', 'M', 'A', 'A', '.', '.'},
		{'.', '.', 'A', '.', 'A', 'S', 'M', 'S', 'M', '.'},
		{'.', 'M', 'A', 'S', 'A', 'M', '.', '.', 'M', 'M'},
		{'.', '.', '.', '.', '.', '.', '.', 'A', '.', 'A'},
		{'S', '.', 'S', '.', 'S', '.', 'S', '.', 'S', 'S'},
		{'.', 'A', '.', 'A', '.', 'A', '.', 'A', '.', 'A'},
		{'M', '.', 'M', '.', 'M', '.', 'M', '.', 'M', 'M'},
		{'.', '.', '.', '.', '.', '.', 'M', 'A', 'S', '.'},
	}

	tests := []struct {
		row, col int
		expected bool
	}{
		{0, 7, false},
		{1, 2, true},
		{2, 6, true},
		{2, 7, true},
		{3, 2, true},
		{3, 4, true},
		{4, 2, false},
		{4, 4, false},
		{5, 7, false},
		{5, 9, false},
		{7, 1, true},
		{7, 3, true},
		{7, 5, true},
		{7, 7, true},
		{7, 9, false},
		{9, 7, false},
	}

	for _, test := range tests {
		actual := searchX(&grid, test.row, test.col)
		if test.expected != actual {
			t.Errorf("Expected %t but got %t", test.expected, actual)
		}
	}
}
