package leetcode

func reverseVowels(s string) string {
	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}
	sArr := []rune(s)
	left := 0
	right := len(sArr) - 1

	for left < right {
		if _, ok := vowels[sArr[left]]; !ok {
			left++
		} else if _, ok := vowels[sArr[right]]; !ok {
			right--
		} else {
			sArr[left], sArr[right] = sArr[right], sArr[left]
			left++
			right--
		}
	}

	return string(sArr)
}
