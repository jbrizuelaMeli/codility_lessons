package countingelements

type MaxCounter struct{}

func (m MaxCounter) Solution(n int, a []int) []int {
	maxCounter := 0
	base := 0
	counters := make([]int, n)
	b := n + 1
	var maxCounterFound bool
	for _, val := range a {
		if val == b {
			maxCounterFound = true
			base = maxCounter
		} else {
			if counters[val-1] < base {
				counters[val-1] = base
			}
			counters[val-1] += 1
			if counters[val-1] > maxCounter {
				maxCounter = counters[val-1]
			}
		}
	}
	if maxCounter > 0 && maxCounterFound {
		for idx := range counters {
			if counters[idx] < base {
				counters[idx] = base
			}
		}
	}
	return counters
}

func (m MaxCounter) SolutionCorrect(n int, a []int) []int {
	maxCounter := 0
	counters := make([]int, n)
	mapCounter := make(map[int]int)
	for _, val := range a {
		mapCounter[val]++
	}
	b := n + 1
	for _, val := range a {
		if val == b {
			for idx := range counters {
				counters[idx] = maxCounter
			}
		} else {
			counters[val-1] += 1
			if counters[val-1] > maxCounter {
				maxCounter = counters[val-1]
			}
		}
	}
	return counters
}
