import "math"

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	i := 0
	sign := 1
	for ; i < len(s) && s[i] == ' '; i++ {
	}
	if i < len(s) && s[i] == '+' {
		i++
	} else if i < len(s) && s[i] == '-' {
		sign = -1
		i++
	}
	number := 0
	for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
		number = number*10 + int(s[i]) - '0'
		if number*sign >= math.MaxInt32 {
			return math.MaxInt32
		}
		if number*sign <= math.MinInt32 {
			return math.MinInt32
		}
	}
	return number * sign
}
