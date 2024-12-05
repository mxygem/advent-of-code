package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("day_two_input.txt")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	levels, err := parseInput(string(input))
	if err != nil {
		log.Fatalf("error parsing input: %v", err)
	}

	// part one
	safe := safeLevels(levels)
	fmt.Printf("safe levels: %d\n", safe)

	// part two
	safe = safeLevelsDampened(levels)
	fmt.Printf("dampened safe levels: %d\n", safe)
}

func parseInput(input string) ([][]int, error) {
	lines := strings.Split(input, "\n")

	levels := make([][]int, 0, len(lines))
	for i, line := range lines {
		ls := strings.Split(line, " ")

		level := make([]int, 0, len(ls))
		for _, l := range ls {
			if l == "" {
				continue
			}

			li, err := strconv.Atoi(l)
			if err != nil {
				return nil, fmt.Errorf("error parsing line %d: %v", i, err)
			}

			level = append(level, li)
		}

		levels = append(levels, level)
	}

	return levels, nil
}

func safeLevels(input [][]int) int {
	var safe int
	for _, levels := range input {
		if len(levels) == 0 {
			continue
		}

		if isSafe(levels) {
			safe++
			continue
		}
	}

	return safe
}

func isSafe(levels []int) bool {
	var last, dir int
	for i, level := range levels {
		if i == 0 {
			last = level
			continue
		}

		if dir == 0 {
			if level < last {
				dir = -1
			} else if level > last {
				dir = 1
			} else {

				return false
			}
		}

		if dir == -1 {
			if last-3 > level {
				return false
			}

			if level > last || level == last {
				return false
			}
		} else if dir == 1 {
			if last+3 < level {
				return false
			}

			if level < last || level == last {
				return false
			}
		}
		last = level
	}

	return true
}

func safeLevelsDampened(input [][]int) int {
	var safe int
	for _, levels := range input {
		lc := len(levels)
		if lc == 0 {
			continue
		}

		if isSafe(levels) {
			safe++
			continue
		}

		for i := 0; i < lc; i++ {
			ls := append([]int{}, levels[:i]...)
			ls = append(ls, levels[i+1:]...)
			ok := isSafe(ls)
			if ok {
				safe++
				break
			}

		}

	}

	return safe
}
