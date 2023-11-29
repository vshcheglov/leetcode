package _74_search_a_2d_matrix

func SearchMatrix(matrix [][]int, target int) bool {
	cols := len(matrix[0])
	rows := len(matrix)
	length := len(matrix[len(matrix)-1]) + cols*(rows-1)
	l, r := 0, length-1
	for l <= r {
		m := l + ((r - l) / 2)
		rowInd := m / cols
		colInd := m - m/cols*cols
		if matrix[rowInd][colInd] == target {
			return true
		}
		if matrix[rowInd][colInd] >= target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return false
}
