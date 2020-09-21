package permutations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	// test input
	var testInput = []int{1, 2, 3}
	// expected
	var expected = [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}
	// assert equality
	assert.ElementsMatch(t, permute(testInput), expected, "they should be equal")
}

func Test2(t *testing.T) {
	// test input
	var testInput = []int{1, 2, 3}
	// expected
	var expected = [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}
	// assert equality
	assert.ElementsMatch(t, permute2(testInput), expected, "they should be equal")
}

func Test3(t *testing.T) {
	// test input
	var testInput = []int{1, 1, 2}
	// expected
	var expected = [][]int{
		{1, 1, 2},
		{1, 2, 1},
		{2, 1, 1},
	}
	// assert equality
	assert.ElementsMatch(t, permuteUnique(testInput), expected, "they should be equal")
}

func Test4(t *testing.T) {
	// test input
	var testInput = []int{1, 1, 2}
	// expected
	var expected = [][]int{
		{1, 1, 2},
		{1, 2, 1},
		{2, 1, 1},
	}
	// assert equality
	assert.ElementsMatch(t, permuteUnique2(testInput), expected, "they should be equal")
}
