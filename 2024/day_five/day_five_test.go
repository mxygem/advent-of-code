package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testRules = map[int][]int{
	47: {53, 13, 61, 29},
	97: {13, 61, 47, 29, 53, 75},
	75: {29, 53, 47, 61, 13},
	61: {13, 53, 29},
	29: {13},
	53: {29, 13},
}

func TestMiddlePagesSum(t *testing.T) {
	testCases := []struct {
		name     string
		rules    map[int][]int
		updates  [][]int
		expected int
	}{
		{
			name: "example",
			rules: map[int][]int{
				47: {53, 13, 61, 29},
				97: {13, 61, 47, 29, 53, 75},
				75: {29, 53, 47, 61, 13},
				61: {13, 53, 29},
				29: {13},
				53: {29, 13},
			},
			updates: [][]int{
				{75, 47, 61, 53, 29},
				{97, 61, 53, 29, 13},
				{75, 29, 13},
				{75, 97, 47, 61, 53},
				{61, 13, 29},
				{97, 13, 75, 29, 47},
			},
			expected: 143,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := middlePagesSum(tc.rules, tc.updates)

			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestRules(t *testing.T) {
	testCases := []struct {
		name        string
		rules       []string
		expected    map[int][]int
		expectedErr error
	}{
		{
			name: "basic",
			rules: []string{
				"47|53",
			},
			expected: map[int][]int{
				47: {53},
			},
		},

		{
			name: "example",
			rules: []string{
				"47|53",
				"97|13",
				"97|61",
				"97|47",
				"75|29",
				"61|13",
				"75|53",
				"29|13",
				"97|29",
				"53|29",
				"61|53",
				"97|53",
				"61|29",
				"47|13",
				"75|47",
				"97|75",
				"47|61",
				"75|61",
				"47|29",
				"75|13",
				"53|13",
			},
			expected: map[int][]int{
				47: {53, 13, 61, 29},
				97: {13, 61, 47, 29, 53, 75},
				75: {29, 53, 47, 61, 13},
				61: {13, 53, 29},
				29: {13},
				53: {29, 13},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := rules(tc.rules)

			if tc.expectedErr != nil {
				assert.EqualError(t, err, tc.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestUpdates(t *testing.T) {
	testCases := []struct {
		name        string
		updates     []string
		expected    [][]int
		expectedErr error
	}{
		{
			name: "basic",
			updates: []string{
				"75,47,61,53,29",
			},
			expected: [][]int{
				{75, 47, 61, 53, 29},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := updates(tc.updates)

			if tc.expectedErr != nil {
				assert.EqualError(t, err, tc.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestFilterRules(t *testing.T) {
	testCases := []struct {
		name     string
		update   []int
		rules    map[int][]int
		expected map[int][]int
	}{
		{
			name: "example",
			update: []int{
				75, 47, 61, 53, 29,
			},
			rules: map[int][]int{
				47: {53, 13, 61, 29},
				97: {13, 61, 47, 29, 53, 75},
				75: {29, 53, 47, 61, 13},
				61: {13, 53, 29},
				29: {13},
				53: {29, 13},
			},
			expected: map[int][]int{
				75: {29, 53, 47, 61, 13},
				47: {53, 13, 61, 29},
				61: {13, 53, 29},
				53: {29, 13},
				29: {13},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := filterRules(tc.update, tc.rules)

			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestValidUpdate(t *testing.T) {
	testCases := []struct {
		name          string
		update        []int
		expectedRule  int
		expectedValid bool
	}{
		{
			name:          "example one - valid",
			update:        []int{75, 47, 61, 53, 29},
			expectedValid: true,
		},
		{
			name:          "example two - valid",
			update:        []int{97, 61, 53, 29, 13},
			expectedValid: true,
		},
		{
			name:          "example three - valid",
			update:        []int{75, 29, 13},
			expectedValid: true,
		},
		{
			name:          "example four - invalid",
			update:        []int{75, 97, 47, 61, 53},
			expectedRule:  97,
			expectedValid: false,
		},
		{
			name:          "example five - invalid",
			update:        []int{61, 13, 29},
			expectedRule:  29,
			expectedValid: false,
		},
		{
			name:          "example six - invalid",
			update:        []int{97, 13, 75, 29, 47},
			expectedRule:  75,
			expectedValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualRule, actualValid := validUpdate(tc.update, testRules)

			assert.Equal(t, tc.expectedRule, actualRule)
			assert.Equal(t, tc.expectedValid, actualValid)
		})
	}
}

func TestInOrder(t *testing.T) {
	testCases := []struct {
		name     string
		update   []int
		key      int
		rule     []int
		expected bool
	}{
		{
			name:     "example - in order",
			update:   []int{75, 47, 61, 53, 29},
			key:      75,
			rule:     []int{29, 53, 47, 61, 13},
			expected: true,
		},
		{
			name:     "example - not in order",
			update:   []int{75, 97, 47, 61, 53},
			key:      97,
			rule:     []int{13, 61, 47, 29, 53, 75},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, inOrder(tc.update, tc.key, tc.rule))
		})
	}
}

func TestMiddlePage(t *testing.T) {
	testCases := []struct {
		name     string
		update   []int
		expected int
	}{
		{
			name:     "example one",
			update:   []int{75, 47, 61, 53, 29},
			expected: 61,
		},
		{
			name:     "example two",
			update:   []int{97, 61, 53, 29, 13},
			expected: 53,
		},
		{
			name:     "example three",
			update:   []int{75, 29, 13},
			expected: 29,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, middlePage(tc.update))
		})
	}
}

// var testRules = map[int][]int{
// 	47: {53, 13, 61, 29},
// 	97: {13, 61, 47, 29, 53, 75},
// 	75: {29, 53, 47, 61, 13},
// 	61: {13, 53, 29},
// 	29: {13},
// 	53: {29, 13},
// }

func TestReorderInvalidUpdate(t *testing.T) {
	testCases := []struct {
		name     string
		update   []int
		expected []int
	}{
		{
			name:     "basic",
			update:   []int{53, 47},
			expected: []int{47, 53},
		},
		// {
		// 	name:     "example one",
		// 	update:   []int{75, 97, 47, 61, 53},
		// 	expected: []int{97, 75, 47, 61, 53},
		// },
		// {
		// 	name:     "example two",
		// 	update:   []int{61, 13, 29},
		// 	expected: []int{61, 29, 13},
		// },
		// {
		// 	name:     "example three",
		// 	update:   []int{97, 13, 75, 29, 47},
		// 	expected: []int{97, 75, 47, 29, 13},
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, reorderUpdate(tc.update, testRules))
		})
	}
}
