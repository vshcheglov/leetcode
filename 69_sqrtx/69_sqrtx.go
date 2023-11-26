package _69_sqrtx

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	l, r := 0, x
	var ans int
	for l <= r {
		m := l + (r-l)/2
		if m*m == x {
			return m
		}
		if m*m < x {
			l = m + 1
			ans = m
		} else {
			r = m - 1
		}
	}
	return ans
}
