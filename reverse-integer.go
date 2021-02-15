import "math"

func reverse(x int) int {
	var reversedNum int
	for ; x != 0; x /= 10 {
		reversedNum = reversedNum*10 + x%10
	}
	if reversedNum > math.MaxInt32 || reversedNum < math.MinInt32 {
		return 0
	}
	return
}
