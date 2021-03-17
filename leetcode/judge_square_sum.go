package leetcode

import "math"

// 0...sqrt(c)
// left can equal to right
func judgeSquareSum(c int) bool {
	if c < 0 {
		return false
	}

	left := 0
	right := int(math.Sqrt(float64(c)))

	for left <= right {
		sum := left*left + right*right
		if sum == c {
			return true
		} else if sum < c {
			left++
		} else {
			right--
		}
	}

	return false
}
