package leetcode

func isAnagram(s string, t string) bool {
	count := [26]int{}
	for _, val := range s {
		count[val-'a']++
	}
	for _, val := range t {
		count[val-'a']--
	}

	for _, val := range count {
		if val != 0 {
			return false
		}
	}

	return true
}
