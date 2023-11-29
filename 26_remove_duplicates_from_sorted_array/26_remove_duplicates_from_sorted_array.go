package _26_remove_duplicates_from_sorted_array

func RemoveDuplicates(nums []int) int {
	count := len(nums)
	if count == 1 {
		return 1
	}
	length := 1
	l, r := 0, 1
	for r <= count-1 {
		if nums[l] == nums[r] {
			r = r + 1
		} else {
			l = l + 1
			nums[l] = nums[r]
			r = r + 1
			length = length + 1
		}
	}
	return length
}
