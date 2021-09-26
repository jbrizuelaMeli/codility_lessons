package arrays

type CyclicRotation struct{}

//Given an array A consisting of N integers and an integer K, returns the array A rotated K times.
func (cr *CyclicRotation) Solution(a []int, k int) []int {
	l := len(a)
	if k == 0 || l == 0 || k%l == 0 {
		return a
	}
	b := make([]int, l)
	for idx, val := range a {
		n := idx + k
		b[n%l] = val
	}
	return b
}
