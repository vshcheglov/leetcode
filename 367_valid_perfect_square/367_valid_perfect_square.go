package _367_valid_perfect_square

func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}
	left, right := 0, num/2
	for left <= right {
		mid := left + (right-left)/2
		square := mid * mid
		if square == num {
			return true
		}
		if square < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
