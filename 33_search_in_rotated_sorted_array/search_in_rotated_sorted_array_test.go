package _33_search_in_rotated_sorted_array

import "testing"

func TestSearch(t *testing.T) {
	testCases := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
	}

	for _, tc := range testCases {
		result := Search(tc.nums, tc.target)
		if result != tc.expected {
			t.Errorf("Search(%v, %d) = %d; expected %d", tc.nums, tc.target, result, tc.expected)
		}
	}
}
