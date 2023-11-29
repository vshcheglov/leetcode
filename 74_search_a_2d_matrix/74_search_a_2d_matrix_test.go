package _74_search_a_2d_matrix

import "testing"

func TestSearchMatrix(t *testing.T) {
	testCases := []struct {
		matrix [][]int
		target int
		want   bool
	}{
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3, true},
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13, false},
		{[][]int{{1}}, 0, false},
		{[][]int{{1}}, 2, false},
		{[][]int{{1}}, 1, true},
	}

	for _, tc := range testCases {
		got := SearchMatrix(tc.matrix, tc.target)
		if got != tc.want {
			t.Errorf("SearchMatrix(%v, %d) = %v; want %v", tc.matrix, tc.target, got, tc.want)
		}
	}
}
