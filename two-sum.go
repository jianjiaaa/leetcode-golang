func twoSum(nums []int, target int) []int {
	store := map[int]int{}
	for idx, num := range nums {
		remaining := target - num
		if _, ok := store[remaining]; ok {
			return []int{store[remaining], idx}
		}
		store[num] = idx
	}
	return []int{}
}
