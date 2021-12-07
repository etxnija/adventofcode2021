package day1

func increasedSliding(mesurements []int, window int) int {
	i := 0

	size := len(mesurements) - window

	for j := 0; j < size; j++ {
		if sum(mesurements[j:j+window]) < sum(mesurements[j+1:j+1+window]) {
			i++
		}
	}
	return i
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
