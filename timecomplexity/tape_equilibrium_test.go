package timecomplexity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTapeEquilibrium_Solution(t *testing.T) {
	te := &TapeEquilibrium{}

	a := []int{3, 1, 2, 4, 3}
	minDiff := te.Solution(a)
	assert.Equal(t, 1, minDiff)
}
