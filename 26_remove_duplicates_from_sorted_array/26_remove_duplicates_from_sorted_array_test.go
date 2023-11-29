package _26_remove_duplicates_from_sorted_array

import "testing"

func TestSearch(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}

	for _, tc := range testCases {
		result := RemoveDuplicates(tc.nums)
		if result != tc.expected {
			t.Errorf("RemoveDuplicates(%v) = %d; expected %d", tc.nums, result, tc.expected)
		}
	}
}
