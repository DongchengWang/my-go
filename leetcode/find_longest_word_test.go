package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLongestWord(t *testing.T) {
	s := "abpcplea"
	dict := []string{"ale", "apple", "monkey", "plea"}

	got := findLongestWord(s, dict)
	want := "apple"
	assert.Equal(t, want, got)
}

func TestFindLongestWord2(t *testing.T) {
	s := "abpcplea"
	dict := []string{"a", "b", "c"}

	got := findLongestWord(s, dict)
	want := "a"
	assert.Equal(t, want, got)
}
