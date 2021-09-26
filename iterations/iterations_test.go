package iterations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	n := 328

	res := Solution(n)
	assert.Equal(t, 2, res)
}

func TestSolution2(t *testing.T) {
	n := 123443

	res := Solution(n)
	assert.Equal(t, 3, res)
}

func TestSolution3(t *testing.T) {
	n := 345365628

	res := Solution(n)
	assert.Equal(t, 3, res)
}

func TestSolution4(t *testing.T) {
	n := 3099887828

	res := Solution(n)
	assert.Equal(t, 3, res)
}
