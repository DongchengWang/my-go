package leetcode

func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	longest := 1
	start := 0

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if i-j < 2 {
				dp[j][i] = s[i] == s[j]
			} else {
				dp[j][i] = s[i] == s[j] && dp[j+1][i-1]
			}

			distance := i - j + 1
			if dp[j][i] && longest < distance {
				longest = distance
				start = j
			}
		}
	}

	return s[start : start+longest]
}
