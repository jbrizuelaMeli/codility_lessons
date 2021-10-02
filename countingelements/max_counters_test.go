package countingelements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxCounters_Solution(t *testing.T) {
	m := &MaxCounter{}
	a := []int{3, 4, 4, 6, 1, 4, 4}
	assert.Equal(t, []int{3, 2, 2, 4, 2}, m.SolutionCorrect(5, a))
	assert.Equal(t, []int{3, 2, 2, 4, 2}, m.Solution(5, a))
}

func TestMaxCounters_Solution2(t *testing.T) {
	m := &MaxCounter{}
	a := []int{3, 4, 4, 6, 1, 4, 4, 5, 4, 6, 4, 3, 4}
	assert.Equal(t, []int{5, 5, 6, 7, 5}, m.SolutionCorrect(5, a))
	assert.Equal(t, []int{5, 5, 6, 7, 5}, m.Solution(5, a))
}
