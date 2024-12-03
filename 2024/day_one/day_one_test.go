package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitInput(t *testing.T) {
	testCases := []struct {
		name            string
		input           string
		expectedListOne []int
		expectedListTwo []int
	}{
		{
			name:            "basic",
			input:           "1   2\n3   4",
			expectedListOne: []int{1, 3},
			expectedListTwo: []int{2, 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			listOne, listTwo := splitInput(tc.input)

			assert.Equal(t, tc.expectedListOne, listOne)
			assert.Equal(t, tc.expectedListTwo, listTwo)
		})
	}
}

func TestDistances(t *testing.T) {
	testCases := []struct {
		name     string
		left     []int
		right    []int
		expected int
	}{
		{
			name:     "basic",
			left:     []int{1, 3},
			right:    []int{2, 4},
			expected: 2,
		},
		{
			name:     "larger & more numbers",
			left:     []int{10, 154, 200},
			right:    []int{20, 400, 450},
			expected: 506,
		},
		{
			name:     "numbers in right smaller than left",
			left:     []int{10, 15, 20},
			right:    []int{5, 10, 15},
			expected: 15,
		},
		{
			name:     "duplicates in left",
			left:     []int{10, 10, 10},
			right:    []int{5, 10, 15},
			expected: 10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := distances(tc.left, tc.right)

			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestSimilarities(t *testing.T) {
	testCases := []struct {
		name     string
		left     []int
		right    []int
		expected int
	}{
		{
			name:     "example data",
			left:     []int{1, 2, 3, 3, 3, 4},
			right:    []int{3, 3, 3, 4, 5, 9},
			expected: 31,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := similarities(tc.left, tc.right)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
