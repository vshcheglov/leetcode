package _852_peak_index_in_mountain_array

func PeakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		m := l + ((r - l) / 2)
		if arr[m] < arr[m+1] {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l
}
