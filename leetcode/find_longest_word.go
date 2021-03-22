package leetcode

func findLongestWord(s string, dictionary []string) string {
	longest := ""

	for _, word := range dictionary {
		l1 := len(longest)
		l2 := len(word)
		if l1 > l2 || (l1 == l2 && longest < word) {
			continue
		}
		if isSubstring(s, word) {
			longest = word
		}
	}

	return longest
}

func isSubstring(s string, word string) bool {
	i := 0
	j := 0
	for i < len(s) && j < len(word) {
		if s[i] == word[j] {
			j++
		}
		i++
	}

	return j == len(word)
}
