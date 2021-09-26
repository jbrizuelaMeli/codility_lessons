package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCyclicRotation_Solution(t *testing.T) {
	cr := &CyclicRotation{}
	a := []int{1, 2, 3, 4, 5, 6}
	k := 2

	b := cr.Solution(a, k)
	assert.Equal(t, []int{5, 6, 1, 2, 3, 4}, b)
}

func TestCyclicRotation_Solution2(t *testing.T) {
	cr := &CyclicRotation{}
	a := []int{0, 0, 0, 0}
	k := 2

	b := cr.Solution(a, k)
	assert.Equal(t, []int{0, 0, 0, 0}, b)
}

func TestCyclicRotation_Solution3(t *testing.T) {
	cr := &CyclicRotation{}
	a := []int{1, 2, 3, 4}
	k := 8

	b := cr.Solution(a, k)
	assert.Equal(t, []int{1, 2, 3, 4}, b)
}

func TestCyclicRotation_Solution4(t *testing.T) {
	cr := &CyclicRotation{}
	a := []int{-1000, -500, 2, 5}
	k := 13

	b := cr.Solution(a, k)
	assert.Equal(t, []int{5, -1000, -500, 2}, b)
}
