package iterations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryGap_Solution(t *testing.T) {
	bg := &BinaryGap{}
	n := 328

	res := bg.Solution(n)
	assert.Equal(t, 2, res)
}

func TestBinaryGap_Solution2(t *testing.T) {
	bg := &BinaryGap{}
	n := 123443

	res := bg.Solution(n)
	assert.Equal(t, 3, res)
}

func TestBinaryGap_Solution3(t *testing.T) {
	bg := &BinaryGap{}
	n := 345365628

	res := bg.Solution(n)
	assert.Equal(t, 3, res)
}

func TestBinaryGap_Solution4(t *testing.T) {
	bg := &BinaryGap{}
	n := 3099887828

	res := bg.Solution(n)
	assert.Equal(t, 3, res)
}
