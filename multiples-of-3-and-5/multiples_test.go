package multiples

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	}

	for _, test := range testCases {
		assert.Equal(t, test.expected, SumOfMultiplesOf3and5(test.input))
	}
}
