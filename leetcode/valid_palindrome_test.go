package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidPalindrome(t *testing.T) {
	got := validPalindrome("aba")
	assert.True(t, got)

	got2 := validPalindrome("abca")
	assert.True(t, got2)
}
