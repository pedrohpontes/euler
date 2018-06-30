package multiples

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumUpTo(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 0},
		{2, 1},
		{3, 3},
		{4, 6},
		{10, 45},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expected, sumUpTo(test.input))
	}
}

func TestSumOfMultiples(t *testing.T) {
	testCases := []struct {
		divisor  int
		input    int
		expected int
	}{
		{3, 0, 0},
		{3, 1, 0},
		{3, 2, 0},
		{3, 3, 0},
		{3, 4, 3},
		{3, 10, 18},
		{5, 5, 0},
		{5, 6, 5},
		{5, 11, 15},
		{5, 21, 50},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expected, sumOfMultiplesOf(test.divisor, test.input))
	}
}

func TestMultiplesOf3and5(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 3},
		{5, 3},
		{6, 8},
		{10, 23},
		{1000, 233168},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expected, SumOfMultiplesOf3and5(test.input))
	}
}
