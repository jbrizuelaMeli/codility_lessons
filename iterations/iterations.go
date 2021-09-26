package iterations

import (
	"strconv"
)

//Given a positive integer N, returns the length of its longest binary gap.
//The function should return 0 if N doesn't contain a binary gap.
func Solution(n int) int {
	//Discard even numbers or 1
	if n == 1 {
		return 0
	}
	//to binary form -> 137 = 10001001
	nb := strconv.FormatInt(int64(n), 2)
	//To array of bytes
	nbb := []byte(nb)
	var gap int
	var gapAux int
	for idx, val := range nbb {
		if val == byte('1') && idx > 0 {
			if gapAux > gap {
				gap = gapAux
			}
			gapAux = 0
		}
		if val == byte('0') {
			gapAux++
		}
	}
	return gap
}
