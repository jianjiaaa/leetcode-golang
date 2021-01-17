func singleNumber(nums []int) int {
	numbers := map[int]int{}
	for _, num := range nums {
		numbers[num]++
	}
	for k, v := range numbers {
		if v == 1 {
			return k
		}
	}
	return 0
}