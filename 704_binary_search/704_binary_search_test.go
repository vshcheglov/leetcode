package _704_binary_search

import "testing"

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{"Test1", []int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{"Test2", []int{-1, 0, 3, 5, 9, 12}, 2, -1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := BinarySearch(tc.nums, tc.target)
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}
