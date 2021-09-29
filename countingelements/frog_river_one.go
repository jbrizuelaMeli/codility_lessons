package countingelements

type FrogRiverOne struct{}

// given a non-empty array a consisting of N integers and integer x,
//returns the earliest time when the frog can jump to the other side of the river.
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
