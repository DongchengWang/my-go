package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
	s1 := "anagram"
	s2 := "nagaram"

	got := isAnagram(s1, s2)
	assert.True(t, got)
}

func TestIsAnagram2(t *testing.T) {
	s1 := "rat"
	s2 := "car"

	got := isAnagram(s1, s2)
	assert.False(t, got)
}
