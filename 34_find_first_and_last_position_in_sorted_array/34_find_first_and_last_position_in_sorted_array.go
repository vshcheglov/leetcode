package _34_find_first_and_last_position_in_sorted_array

func SearchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	leftIndex := -1
	for l <= r {
		m := l + ((r - l) / 2)
		if nums[m] == target {
			leftIndex = m
		}
		if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	if leftIndex == -1 {
		return []int{-1, -1}
	}
	rightIndex := leftIndex
	r = len(nums) - 1
	for l <= r {
		m := l + ((r - l) / 2)
		if nums[m] == target {
			rightIndex = m
		}
		if nums[m] <= target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return []int{leftIndex, rightIndex}
}
