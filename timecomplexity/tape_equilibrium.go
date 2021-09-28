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
	for p := 1; p < l; p++ {
		for i := idxP; i < p; i++ {
			a0 += a[i]
			if p > 1 {
				a1 -= a[i]
			}
			idxP++
		}
		for j := idxP; j < l && p == 1; j++ {
			a1 += a[j]
		}
		diff := abs(a0 - a1)
		if diff <= minDiff || p == 1 {
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
