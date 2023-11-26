package _1_two_sum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{"Test1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"Test2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"Test3", []int{3, 3}, 6, []int{0, 1}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := twoSum(tc.nums, tc.target)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("%s: Expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}
