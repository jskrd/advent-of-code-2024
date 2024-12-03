package day01

import (
	"reflect"
	"testing"
)

const input = "3   4\n4   3\n2   5\n1   3\n3   9\n3   3"

func TestSolvePartOne(t *testing.T) {
	lists := [][]int{
		{3, 4, 2, 1, 3, 3},
		{4, 3, 5, 3, 9, 3},
	}

	expected := 11
	actual := solvePartOne(lists)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestSolvePartTwo(t *testing.T) {
	lists := [][]int{
		{3, 4, 2, 1, 3, 3},
		{4, 3, 5, 3, 9, 3},
	}

	expected := 31
	actual := solvePartTwo(lists)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestParseInput(t *testing.T) {
	input := "3   4\n4   3\n2   5\n1   3\n3   9\n3   3"

	expected := [][]int{
		{3, 4, 2, 1, 3, 3},
		{4, 3, 5, 3, 9, 3},
	}
	actual := parseInput(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}
