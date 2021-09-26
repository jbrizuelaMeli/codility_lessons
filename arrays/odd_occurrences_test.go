package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOddOccurrences_Solution(t *testing.T) {
	oo := OddOccurrences{}

	a := []int{9, 3, 9, 3, 9, 7, 9}
	unpaired := oo.Solution(a)
	assert.Equal(t, 7, unpaired)
}

func TestOddOccurrences_Solution2(t *testing.T) {
	oo := OddOccurrences{}

	a := []int{9, 3, 9, 3, 9, 7, 9, 7}
	unpaired := oo.Solution(a)
	assert.Equal(t, 0, unpaired)
}

func TestOddOccurrences_Solution3(t *testing.T) {
	oo := OddOccurrences{}

	var a []int
	unpaired := oo.Solution(a)
	assert.Equal(t, 0, unpaired)
}
