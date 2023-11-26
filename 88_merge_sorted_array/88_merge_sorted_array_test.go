package _88_merge_sorted_array

import (
	"reflect"
	"testing"
)

func TestMergeSortedArray1(t *testing.T) {
	arr := []int{1, 2, 3, 0, 0, 0}
	Merge(arr, 3, []int{2, 5, 6}, 3)
	expected := []int{1, 2, 2, 3, 5, 6}

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Expected %v, got %v", expected, arr)
	}
}

func TestMergeSortedArray2(t *testing.T) {
	arr := []int{1}
	Merge(arr, 1, []int{}, 0)
	expected := []int{1}

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Expected %v, got %v", expected, arr)
	}
}

func TestMergeSortedArray3(t *testing.T) {
	arr := []int{0}
	Merge(arr, 0, []int{1}, 1)
	expected := []int{1}

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Expected %v, got %v", expected, arr)
	}
}
