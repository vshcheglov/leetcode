package _26_remove_duplicates_from_sorted_array

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		nums         []int
		expected     int
		expectedNums []int
	}{
		{[]int{1, 1, 2}, 2, []int{1, 2}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
	}

	for _, tc := range testCases {
		k := RemoveDuplicates(tc.nums)
		if k != tc.expected {
			t.Errorf("RemoveDuplicates(%v) = %d; expected %d", tc.nums, k, tc.expected)
		}
		for i := 0; i < k; i++ {
			if tc.nums[i] != tc.expectedNums[i] {
				t.Errorf("After RemoveDuplicates, nums[%d] = %d; expected %d", i, tc.nums[i], tc.expectedNums[i])
			}
		}
	}
}
