func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	pointer := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[pointer] {
			pointer++
			nums[pointer] = nums[i]
		}
	}
	return pointer + 1
}
