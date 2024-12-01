package day01

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

func Solve() (int, int) {
	data, _ := os.ReadFile("internal/day01/input.txt")

	lines := strings.Split(string(data), "\n")

	return SolvePartOne(lines), SolvePartTwo(lines)
}

func SolvePartOne(lines []string) int {
	leftList, rightList := parseLists(lines)

	sort.Ints(leftList)
	sort.Ints(rightList)

	distances := []int{}
	for i, left := range leftList {
		right := rightList[i]
		distance := 0
		if left > right {
			distance = left - right
		} else {
			distance = right - left
		}
		distances = append(distances, distance)
	}

	sum := 0
	for _, distance := range distances {
		sum += distance
	}
	return sum
}

func SolvePartTwo(lines []string) int {
	leftList, rightList := parseLists(lines)

	similarities := []int{}
	for _, left := range leftList {
		count := 0
		for _, right := range rightList {
			if left == right {
				count++
			}
		}
		similarities = append(similarities, left*count)
	}

	sum := 0
	for _, similarity := range similarities {
		sum += similarity
	}
	return sum
}

func parseLists(lines []string) ([]int, []int) {
	leftList := []int{}
	rightList := []int{}

	for _, line := range lines {
		parts := strings.Split(line, "   ")
		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1])
		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	return leftList, rightList
}
