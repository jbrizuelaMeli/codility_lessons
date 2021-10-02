package countingelements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingInteger_Solution(t *testing.T) {
	m := &MissingInteger{}
	a := []int{1, 3, 6, 4, 1, 2}
	assert.Equal(t, 5, m.Solution(a))
}

func TestMissingInteger_Solution1(t *testing.T) {
	m := &MissingInteger{}
	a := []int{1, 3, 6, 4, 1, 2, 7, 16}
	assert.Equal(t, 5, m.Solution(a))
}

func TestMissingInteger_Solution2(t *testing.T) {
	m := &MissingInteger{}
	a := []int{1, 2, 3}
	assert.Equal(t, 4, m.Solution(a))
}

func TestMissingInteger_Solution3(t *testing.T) {
	m := &MissingInteger{}
	a := []int{-1, -3}
	assert.Equal(t, 1, m.Solution(a))
}

func TestMissingInteger_Solution4(t *testing.T) {
	m := &MissingInteger{}
	a := []int{1}
	assert.Equal(t, 2, m.Solution(a))
}

func TestMissingInteger_Solution5(t *testing.T) {
	m := &MissingInteger{}
	a := []int{-1}
	assert.Equal(t, 1, m.Solution(a))
}

func TestMissingInteger_Solution6(t *testing.T) {
	m := &MissingInteger{}
	a := []int{0}
	assert.Equal(t, 1, m.Solution(a))
}
