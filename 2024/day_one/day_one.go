package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("day_one_input.txt")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	n := dayOne(string(content))

	fmt.Printf("distance total: %d\n", n)
}

func dayOne(input string) int {
	// split input into two lists
	left, right := splitInput(input)
	// fmt.Printf("left: %v\nright: %v\n", left, right)

	// sort the lists
	slices.Sort(left)
	slices.Sort(right)
	// fmt.Printf("left: %v\nright: %v\n", left, right)

	// calculate the distances
	distance := distances(left, right)

	return distance
}

func splitInput(input string) ([]int, []int) {
	left, right := []int{}, []int{}

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "   ")

		leftInt, leftErr := strconv.Atoi(split[0])
		rightInt, rightErr := strconv.Atoi(split[1])
		if leftErr != nil || rightErr != nil {
			continue
		}

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
