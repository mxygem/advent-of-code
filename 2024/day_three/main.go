package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/mxygem/advent-of-code/internal"
)

func main() {
	input := internal.Open("day_three_input.txt")

	instructions := extractInstructions(input)
	pairs := numericPairs(instructions)
	sum := sumInstructions(pairs)

	// part one
	fmt.Printf("part one sum: %d\n", sum)
}

// extractInstructions looks through the provided input for sequences that
// match the form of 'mul(X,Y)' where X and Y are each 1-3 digit numbers.
// Any extra characters added, including spaces, render a particular sequence
// invalid and must be ignored.
func extractInstructions(input string) []string {
	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	return r.FindAllString(input, -1)
}

// numericPairs returns a collection of x,y pairs extracted from the provided
// input. For each instruction in the input, it:
//  1. trims the `mul()` text
//  2. splits the remaining text `x,y` at the comma and expects two values
//     to be found
//  3. attempts to convert the values into integers
//  4. adds valid pairs to the collection to be returned
func numericPairs(input []string) [][]int {
	pairs := make([][]int, 0, len(input))

	for _, in := range input {
		in = strings.TrimPrefix(in, "mul(")
		in = strings.TrimSuffix(in, ")")
		s := strings.Split(in, ",")

		if len(s) != 2 {
			fmt.Printf("did not find two digits: %q\n", s)
			continue
		}

		x, xErr := strconv.Atoi(s[0])
		y, yErr := strconv.Atoi(s[1])
		if xErr != nil || yErr != nil {
			fmt.Printf("x and/or y are not ints. x:%q, y:%q\n", s[0], s[1])
			continue
		}

		pairs = append(pairs, []int{x, y})
	}

	if len(pairs) == 0 {
		return nil
	}

	return pairs
}

// sumInstructions multiplies the digits of each instruction together and
// returns the total sum of all results.
func sumInstructions(input [][]int) int {
	var sum int

	for _, in := range input {
		if len(in) != 2 {
			fmt.Printf("skipping non-pair instruction")
			continue
		}

		x := in[0]
		y := in[1]

		sum += x * y
	}

	return sum
}
