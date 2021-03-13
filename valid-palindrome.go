import "strings"

func isPalindrome(s string) bool {
	lower := strings.ToLower(s)
	for start, end := 0, len(lower)-1; start < end; {
		if !(lower[start] >= '0' && lower[start] <= '9' || lower[start] >= 'a' && lower[start] <= 'z') {
			start++
			continue
		}
		if !(lower[end] >= '0' && lower[end] <= '9' || lower[end] >= 'a' && lower[end] <= 'z') {
			end--
			continue
		}
		if lower[start] != lower[end] {
			return false
		}
		start++
		end--
	}
	return true
}
