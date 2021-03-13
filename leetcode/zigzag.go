package leetcode

import (
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var builder strings.Builder
	n := len(s)
	cycle := 2*numRows - 2
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < n; j += cycle {
			builder.WriteByte(s[j+i])
			if i != 0 && i != numRows-1 && j+cycle-i < n {
				builder.WriteByte(s[j+cycle-i])
			}
		}
	}

	return builder.String()
}
