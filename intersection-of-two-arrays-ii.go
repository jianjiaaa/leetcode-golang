func intersect(nums1 []int, nums2 []int) []int {
	map1 := map[int]int{}
	for _, num := range nums1 {
		map1[num]++
	}
	overlap := []int{}
	for _, num := range nums2 {
		if count, ok := map1[num]; ok && count > 0 {
			overlap = append(overlap, num)
			map1[num]--
		}
	}
	return overlap
}
