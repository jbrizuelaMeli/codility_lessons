package arrays

type OddOccurrences struct{}

//The array contains an odd number of elements,
//and each element of the array can be paired with another element that has the same value,
//except for one element that is left unpaired.
//Given an array A consisting of N integers fulfilling the above conditions, returns the value of the unpaired element.
func (o *OddOccurrences) Solution(a []int) int {
	l := len(a)
	//Discard if array contains an even number of elements or is empty
	if l == 0 || l%2 == 0 {
		return 0
	}
	occurrences := make(map[int]int)
	for _, val := range a {
		occurrences[val]++
	}
	for idx, c := range occurrences {
		//if the number of occurrences is odd it means that it contains the unpaired value
		if c%2 != 0 {
			return idx
		}
	}
	return 0
}
