func moveZeroes(nums []int) {
	index := 0
	for i, num := range nums {
		if num != 0 {
			nums[index], nums[i] = num, nums[index]
			index++
		}
	}
}
