package pairsum

func PairSum(numbers []int, targetSum int) []int {
	previous := make(map[int]int)
	for i, n := range numbers {
		complement := targetSum - n
		if hashMapIndex, ok := previous[complement]; ok {
			return []int{hashMapIndex, i}
		}
		previous[n] = i
	}
	return []int{}
}
