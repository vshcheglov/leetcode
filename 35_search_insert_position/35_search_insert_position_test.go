package _35_search_insert_position

import "testing"

func TestSearchInsertPosition(t *testing.T) {
	testCases := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
	}

	for _, tc := range testCases {
		result := SearchInsert(tc.nums, tc.target)
		if result != tc.expected {
			t.Errorf("SearchInsert(%v, %d) = %d; expected %d", tc.nums, tc.target, result, tc.expected)
		}
	}
}
