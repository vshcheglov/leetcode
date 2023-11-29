package _206_reverse_linked_list

import (
	"reflect"
	"testing"
)

func fromArray(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head
	for _, val := range values[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

func toArray(head *ListNode) []int {
	values := []int{}
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	return values
}

func TestReverseLinkedList(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{}, []int{}},
	}

	for _, tc := range testCases {
		inputList := fromArray(tc.input)
		reversedList := ReverseList(inputList)
		result := toArray(reversedList)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("ReverseList(%v) = %v; expected %v", tc.input, result, tc.expected)
		}
	}
}
