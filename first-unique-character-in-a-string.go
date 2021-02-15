func firstUniqChar(s string) int {
	alphabets := make([]int, 26)
	for i, _ := range alphabets {
		alphabets[i] = -1
	}
	for idx, char := range s {
		if alphabets[char-'a'] == -1 {
			alphabets[char-'a'] = idx
		} else if alphabets[char-'a'] >= 0 {
			alphabets[char-'a'] = -2
		}
	}
	minIdx := len(s)
	for _, strIdx := range alphabets {
		if strIdx < minIdx && strIdx > -1 {
			minIdx = strIdx
		}
	}
	if minIdx == len(s) {
		return -1
	}
	return minIdx
}
