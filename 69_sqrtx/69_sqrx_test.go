package _69_sqrtx

import "testing"

func TestMySqrt(t *testing.T) {
	testCases := []struct {
		x        int
		expected int
	}{
		{4, 2},
		{8, 2},
		{9, 3},
	}

	for _, tc := range testCases {
		result := mySqrt(tc.x)
		if result != tc.expected {
			t.Errorf("mySqrt(%d) = %d; expected %d", tc.x, result, tc.expected)
		}
	}
}
