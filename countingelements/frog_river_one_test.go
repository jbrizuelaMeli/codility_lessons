package countingelements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrogRiverOne_Solution(t *testing.T) {
	te := &FrogRiverOne{}

	a := []int{1, 3, 1, 4, 2, 3, 5, 4}
	earliestTime := te.Solution(5, a)
	assert.Equal(t, 6, earliestTime)
}

func TestFrogRiverOne_Solution2(t *testing.T) {
	te := &FrogRiverOne{}

	a := []int{4, 23, 4, 1, 7, 4, 5, 7, 4, 5, 6, 7, 8, 34, 2, 3, 8, 9, 10, 14, 11, 12, 13, 15, 15, 15, 16, 19, 22, 21, 20, 4, 1, 3, 18, 17}
	earliestTime := te.Solution(4, a)
	assert.Equal(t, 15, earliestTime)
}
