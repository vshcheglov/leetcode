package _34_find_first_and_last_position_in_sorted_array

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	testCases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{}, 0, []int{-1, -1}},
	}

	for _, tc := range testCases {
		result := SearchRange(tc.nums, tc.target)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("SearchRange(%v, %d) = %v; expected %v", tc.nums, tc.target, result, tc.expected)
		}
	}
}
