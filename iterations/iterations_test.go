package iterations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	bg := &BinaryGap{}
	n := 328

	res := bg.Solution(n)
	assert.Equal(t, 2, res)
}

func TestSolution2(t *testing.T) {
	bg := &BinaryGap{}
	n := 123443

	res := bg.Solution(n)
	assert.Equal(t, 3, res)
}

func TestSolution3(t *testing.T) {
	bg := &BinaryGap{}
	n := 345365628

	res := bg.Solution(n)
	assert.Equal(t, 3, res)
}

func TestSolution4(t *testing.T) {
	bg := &BinaryGap{}
	n := 3099887828

	res := bg.Solution(n)
	assert.Equal(t, 3, res)
}
