package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	got := twoSum([]int{2, 7, 11, 15}, 9)
	want := []int{1, 2}

	assert.Equal(t, got, want)
}
