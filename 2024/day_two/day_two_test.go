package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expected    [][]int
		expectedErr error
	}{
		{
			name:  "example data",
			input: "7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9",
			expected: [][]int{
				{7, 6, 4, 2, 1},
				{1, 2, 7, 8, 9},
				{9, 7, 6, 2, 1},
				{1, 3, 2, 4, 5},
				{8, 6, 4, 4, 1},
				{1, 3, 6, 7, 9},
			},
		},
		{
			name:        "invalid input - non-integer",
			input:       "7 6 a 3 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9",
			expectedErr: fmt.Errorf("error parsing line 0: strconv.Atoi: parsing \"a\": invalid syntax"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := parseInput(tc.input)

			if tc.expectedErr != nil {
				assert.EqualError(t, err, tc.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestSafeLevels(t *testing.T) {
	testCases := []struct {
		name     string
		input    [][]int
		expected int
	}{
		{
			name: "example data",
			input: [][]int{
				{7, 6, 4, 2, 1},
				{1, 2, 7, 8, 9},
				{9, 7, 6, 2, 1},
				{1, 3, 2, 4, 5},
				{8, 6, 4, 4, 1},
				{1, 3, 6, 7, 9},
			},
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := safeLevels(tc.input)

			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestIsSafe(t *testing.T) {
	testCases := []struct {
		name     string
		levels   []int
		expected bool
	}{
		{
			name:     "safe - all levels are decreasing by 1 or 2",
			levels:   []int{7, 6, 4, 2, 1},
			expected: true,
		},
		{
			name:     "safe - all levels are increasing by 1, 2, or 3",
			levels:   []int{1, 3, 6, 7, 9},
			expected: true,
		},
		{
			name:     "unsafe - 2, 7 increase too large - 5",
			levels:   []int{1, 2, 7, 8, 9},
			expected: false,
		},
		{
			name:     "unsafe - 6, 2 decrease too large - 4",
			levels:   []int{9, 7, 6, 2, 1},
			expected: false,
		},
		{
			name:     "unsafe - 1, 3 is increasing but 3, 2 is decreasing",
			levels:   []int{1, 3, 2, 4, 5},
			expected: false,
		},
		{
			name:     "unsafe - 4, 4 is neither increasing nor decreasing",
			levels:   []int{8, 6, 4, 4, 1},
			expected: false,
		},
		{
			name:     "unsafe - direction changes 27, 29, 27",
			levels:   []int{20, 21, 24, 25, 27, 29, 27},
			expected: false,
		},
		{
			name:     "unsafe - no change for first two",
			levels:   []int{62, 62, 65, 68, 69, 72, 75},
			expected: false,
		},
		{
			name:     "unsafe - no change for first three",
			levels:   []int{68, 68, 68, 69, 72},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := isSafe(tc.levels)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
