package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractInstructions(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "bare instruction with two single digits",
			input:    "mul(1,1)",
			expected: []string{"mul(1,1)"},
		},
		{
			name:     "readme example",
			input:    "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			expected: []string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"},
		},
		{
			name:     "no match - spaces around numbers",
			input:    "mul( 1,2 )",
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := extractInstructions(tc.input)

			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestNumericPairs(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected [][]int
	}{
		{
			name:     "single, single digit pair",
			input:    []string{"mul(1,1)"},
			expected: [][]int{{1, 1}},
		},
		{
			name:     "readme example",
			input:    []string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"},
			expected: [][]int{{2, 4}, {5, 5}, {11, 8}, {8, 5}},
		},
		{
			name:     "invalid instruction",
			input:    []string{"foobar"},
			expected: nil,
		},
		{
			name:     "invalid instruction - x,y not ints",
			input:    []string{"mul(a,b)"},
			expected: nil,
		},
		{
			name:     "invalid instruction - missing second digit",
			input:    []string{"mul(1)"},
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := numericPairs(tc.input)

			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestSumInstructions(t *testing.T) {
	testCases := []struct {
		name     string
		input    [][]int
		expected int
	}{
		{
			name:     "zeros in input return zero",
			input:    [][]int{{10, 0}, {0, 2}},
			expected: 0,
		},
		{
			name:     "single non-zero instruction",
			input:    [][]int{{2, 1}},
			expected: 2,
		},
		{
			name:     "non-pair instructions are skipped",
			input:    [][]int{{1000}, {2, 5}, {50}},
			expected: 10,
		},
		{
			name:     "readme example",
			input:    [][]int{{2, 4}, {5, 5}, {11, 8}, {8, 5}},
			expected: 161,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := sumInstructions(tc.input)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
