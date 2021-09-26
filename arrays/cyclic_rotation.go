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
		//ex: a[1,2,3,4,5] k = 2 -> a[4,5,1,2,3]
		// a[idx] -> a[idx + k] only if (idx + k < len(a) - 1) else a[idx] -> a[(idx + k)%l]
		// (idx + k)%l = [0,...,l]
		b[n%l] = val
	}
	return b
}
