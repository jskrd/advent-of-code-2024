package day03

import (
	"os"
	"regexp"
	"strconv"
)

func Solve() (int, int) {
	input, _ := os.ReadFile("internal/day03/input.txt")

	instructions := parseInput(string(input))

	return solvePartOne(instructions), solvePartTwo(instructions)
}

func solvePartOne(instructions []string) int {
	sum := 0
	for _, instruction := range instructions {
		if instruction[:3] == "mul" {
			re := regexp.MustCompile(`([0-9]+),([0-9]+)`)
			matches := re.FindStringSubmatch(instruction)
			a, _ := strconv.Atoi(matches[1])
			b, _ := strconv.Atoi(matches[2])
			sum += a * b
		}
	}
	return sum
}

func solvePartTwo(instructions []string) int {
	enabled := true
	sum := 0
	for _, instruction := range instructions {
		if instruction == "do()" {
			enabled = true
		}
		if instruction == "don't()" {
			enabled = false
		}
		if instruction[:3] == "mul" {
			if !enabled {
				continue
			}

			re := regexp.MustCompile(`([0-9]+),([0-9]+)`)
			matches := re.FindStringSubmatch(instruction)
			a, _ := strconv.Atoi(matches[1])
			b, _ := strconv.Atoi(matches[2])
			sum += a * b
		}
	}
	return sum
}

func parseInput(input string) []string {
	re := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)|do\(\)|don't\(\)`)
	return re.FindAllString(input, -1)
}
