package contiguoussubarrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {

	input := []int{1, 2, 1, 2, 3}

	var expected int = 1

	// assert equality
	assert.Equal(t, expected, duplicatesOnSegement(input), "they should be equal")
}

func Test2(t *testing.T) {

	input := []int{0, 0, 0}

	var expected int = 3

	// assert equality
	assert.Equal(t, expected, duplicatesOnSegement(input), "they should be equal")
}
