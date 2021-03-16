package leetcode

import "testing"

func TestTwoSum(t *testing.T) {
	got := twoSum([]int{2, 7, 11, 15}, 9)
	want := []int{1, 2}

	for i, v := range got {
		if v != want[i] {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}

}
