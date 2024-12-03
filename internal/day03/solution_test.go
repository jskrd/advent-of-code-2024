package day03

import (
	"reflect"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	memory := []string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"}

	expected := 161
	actual := solvePartOne(memory)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestSolvePartTwo(t *testing.T) {
	memory := []string{"mul(2,4)", "don't()", "do()", "mul(8,5)"}

	expected := 48
	actual := solvePartTwo(memory)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestParseInput(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			[]string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"},
		},
		{
			"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			[]string{"mul(2,4)", "don't()", "mul(5,5)", "mul(11,8)", "do()", "mul(8,5)"},
		},
	}

	for _, test := range tests {
		actual := parseInput(test.input)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("Expected %v but got %v", test.expected, actual)
		}
	}
}
