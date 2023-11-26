package _367_valid_perfect_square

import "testing"

func TestIsPerfectSquare(t *testing.T) {
	testCases := []struct {
		name     string
		num      int
		expected bool
	}{
		{"Test1", 16, true},
		{"Test2", 14, false},
		{"Test3", 0, true},
		{"Test4", 1, true},
		{"Test5", 2, false},
		{"Test6", 4, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isPerfectSquare(tc.num)
			if result != tc.expected {
				t.Errorf("%s: Expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}
