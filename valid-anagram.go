func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	alphabetsCounter := make([]int, 26)
	for i := 0; i < len(s); i++ {
		alphabetsCounter[s[i]-'a']++
		alphabetsCounter[t[i]-'a']--
	}
	for _, count := range alphabetsCounter {
		if count != 0 {
			return false
		}
	}
	return true
}
