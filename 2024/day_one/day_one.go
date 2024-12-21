package main

import (
	"bufio"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/mxygem/advent-of-code/internal"
)

func main() {
	input := internal.Open("day_one_input.txt")

	left, right := splitInput(string(input))

	slices.Sort(left)
	slices.Sort(right)

	// part one
	distance := distances(left, right)
	fmt.Printf("distance total: %d\n", distance)

	// part two
	similarity := similarities(left, right)
	fmt.Printf("similarity total: %d\n", similarity)
}

func splitInput(input string) ([]int, []int) {
	left, right := []int{}, []int{}

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "   ")

		leftInt, _ := strconv.Atoi(split[0])
		rightInt, _ := strconv.Atoi(split[1])

		left = append(left, leftInt)
		right = append(right, rightInt)
	}

	return left, right
}

func distances(left, right []int) int {
	var n int

	for i, l := range left {
		n += int(math.Abs(float64(right[i] - l)))
	}

	return n
}

func similarities(left, right []int) int {
	var n int
	for _, l := range left {

		var c int
		for _, r := range right {
			if r < l {
				continue
			}

			if l != r {
				n += l * c
				break
			}

			c++
		}
	}

	return n
}
