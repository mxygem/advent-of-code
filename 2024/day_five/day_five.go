package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("day_five_input.txt")
	if err != nil {
		log.Fatalf("reading input: %v", err)
	}

	rules, updates, err := parseInput(string(input))
	if err != nil {
		log.Fatalf("parsing input: %v", err)
	}

	sum := middlePagesSum(rules, updates)
	fmt.Printf("day one middle pages sum: %v\n", sum)
}

func parseInput(input string) (map[int][]int, [][]int, error) {
	rulesInput := strings.Split(string(input), "\n\n")[0]
	updatesInput := strings.Split(string(input), "\n\n")[1]

	rules, err := rules(strings.Split(rulesInput, "\n"))
	if err != nil {
		return nil, nil, fmt.Errorf("parsing rules: %v", err)
	}

	updates, err := updates(strings.Split(updatesInput, "\n"))
	if err != nil {
		return nil, nil, fmt.Errorf("parsing updates: %v", err)
	}

	return rules, updates, nil
}

func rules(rulesIn []string) (map[int][]int, error) {
	rules := make(map[int][]int, len(rulesIn))

	for _, rule := range rulesIn {
		parts := strings.Split(rule, "|")

		first, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("parsing left side of rule: %v", err)
		}
		second, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("parsing right side of rule: %v", err)
		}
		rules[first] = append(rules[first], second)
	}
	return rules, nil
}

func updates(updatesIn []string) ([][]int, error) {
	updates := make([][]int, 0, len(updatesIn))

	for _, updateIn := range updatesIn {
		parts := strings.Split(updateIn, ",")
		update := make([]int, 0, len(parts))

		for _, part := range parts {
			if part == "" {
				continue
			}

			page, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("parsing update page: %v", err)
			}
			update = append(update, page)
		}
		if len(update) == 0 {
			continue
		}

		updates = append(updates, update)
	}
	return updates, nil
}

func middlePagesSum(rules map[int][]int, updates [][]int) int {
	var sum int

	for _, update := range updates {
		fr := filterRules(update, rules)
		if _, valid := validUpdate(update, fr); !valid {
			continue
		}

		sum += middlePage(update)
	}

	return sum
}

func filterRules(update []int, rules map[int][]int) map[int][]int {
	filtered := map[int][]int{}

	for _, page := range update {
		if _, ok := rules[page]; ok {
			filtered[page] = append(filtered[page], rules[page]...)
		}
	}

	return filtered
}

func validUpdate(update []int, rules map[int][]int) (int, bool) {
	for key, rule := range rules {
		if !inOrder(update, key, rule) {
			return key, false
		}
	}

	return 0, true
}

func reorderUpdate(update []int, rules map[int][]int) []int {
	invalidKey, _ := validUpdate(update, rules)
	fmt.Printf("invalidKey: %v\n", invalidKey)

	return nil
}

func inOrder(update []int, key int, rule []int) bool {
	keyIdx := slices.Index(update, key)

	for i, u := range update {
		for _, r := range rule {
			if u == r && i < keyIdx {
				return false
			}
		}
	}

	return true
}

func middlePage(update []int) int {
	return update[len(update)/2]
}
