package leetcode

// Init left, right.
// Iterate numbers,
// if < sum, move right to middle;
// if > sum, move left to middle.
func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		sum := numbers[left] + numbers[right]
		if target < sum {
			right--
		} else if target > sum {
			left++
		} else {
			return []int{left + 1, right + 1}
		}
	}

	return nil
}
