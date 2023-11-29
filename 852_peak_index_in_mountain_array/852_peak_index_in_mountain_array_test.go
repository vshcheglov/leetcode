package _852_peak_index_in_mountain_array

import "testing"

func TestPeakIndexInMountainArray(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{0, 1, 0}, 1},
		{[]int{0, 2, 1, 0}, 1},
		{[]int{0, 10, 5, 2}, 1},
		{[]int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}, 2},
		{[]int{3, 5, 3, 2, 0}, 1},
	}

	for _, tc := range testCases {
		result := PeakIndexInMountainArray(tc.nums)
		if result != tc.expected {
			t.Errorf("PeakIndexInMountainArray(%v) = %d; expected %d", tc.nums, result, tc.expected)
		}
	}
}
