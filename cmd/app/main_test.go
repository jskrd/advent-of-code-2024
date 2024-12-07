package main

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		day    int
		part   int
		input  string
		expect int
	}{
		{1, 1, "../../test/data/day01/input.txt", 3508942},
		{1, 2, "../../test/data/day01/input.txt", 26593248},
		{2, 1, "../../test/data/day02/input.txt", 279},
		{2, 2, "../../test/data/day02/input.txt", 343},
		{3, 1, "../../test/data/day03/input.txt", 182780583},
		{3, 2, "../../test/data/day03/input.txt", 90772405},
		{4, 1, "../../test/data/day04/input.txt", 2642},
		{4, 2, "../../test/data/day04/input.txt", 1974},
		{5, 1, "../../test/data/day05/input.txt", 4637},
		{5, 2, "../../test/data/day05/input.txt", 6370},
	}

	for _, test := range tests {
		actual := solve(test.day, test.part, test.input)
		if test.expect != actual {
			t.Errorf("Expected %d but got %d", test.expect, actual)
		}
	}
}
