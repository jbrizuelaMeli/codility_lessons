package timecomplexity

type TapeEquilibrium struct{}

//Given a non-empty array a of N integers, returns the minimal difference that can be achieved.
//The difference between the two parts is the value of: |(A[0] + A[1] + ... + A[P − 1]) − (A[P] + A[P + 1] + ... + A[N − 1])|
//Such that P has the values 1,2,3...N
func (te TapeEquilibrium) Solution(a []int) int {
	l := len(a)
	if l == 0 {
		return 0
	}
	var a0, a1, minDiff int
	idxP := 0
	for p := range a {
		if p == l-1 {
			break
		}
		for i := idxP; i < p+1; i++ {
			a0 += a[i]
			if p > 0 {
				a1 -= a[i]
			}
			idxP++
		}
		for j := idxP; j < l && p == 0; j++ {
			a1 += a[j]
		}
		diff := abs(a0 - a1)
		if diff <= minDiff || p == 0 {
			minDiff = diff
		}
	}
	return minDiff
}

// Abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
