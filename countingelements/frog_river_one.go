package countingelements

type FrogRiverOne struct{}

// given a non-empty array a consisting of N integers and integer x,
//returns the earliest time when the frog can jump to the other side of the river.
//The frog can cross only when leaves appear at every position across the river from 1 to x (that is,
//we want to find the earliest moment when all the positions from 1 to x are covered by leaves)
//If the frog is never able to jump to the other side of the river, the function should return −1.
func (f FrogRiverOne) Solution(x int, a []int) int {
	l := len(a)
	if l == 0 || x < 1 {
		return -1
	}
	steps := make(map[int]bool)
	for idx, val := range a {
		if exist, _ := steps[val]; !exist && val <= x {
			steps[val] = true
		}
		if len(steps) == x {
			return idx
		}
	}
	return -1
}
