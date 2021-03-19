package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseVowels(t *testing.T) {
	got := reverseVowels("leetcode")
	want := "leotcede"
	assert.Equal(t, want, got)
}
