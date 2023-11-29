package _33_search_in_rotated_sorted_array

func Search(nums []int, target int) int {
	count := len(nums)
	minIndex := 0
	l, r := 0, count-1
	for l <= r {
		m := l + ((r - l) / 2)
		if nums[m] < nums[minIndex] {
			minIndex = m
			r = m - 1
		} else {
			l = m + 1
		}
	}

	l, r = 0, count-1
	for l <= r {
		m := l + ((r - l) / 2)
		rotIndex := (m + minIndex) % count
		v := nums[rotIndex]
		if v == target {
			return rotIndex
		}
		if v < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}
