package countingelements

type MissingInteger struct{}

func (m MissingInteger) Solution(a []int) int {
	l := len(a)
	counters := make([]int, l)
	for _, val := range a {
		if val > 0 && val-1 < l {
			counters[val-1]++
		}
	}
	lastMiss := 0
	for val, q := range counters {
		lastMiss = val + 1
		if q == 0 {
			return lastMiss
		}
	}
	return lastMiss + 1
}
