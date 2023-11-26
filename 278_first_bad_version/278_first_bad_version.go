package _278_first_bad_version

var SavedFirstBadVersion int

func FirstBadVersion(n int) int {
	leftIndex := 1
	rightIndex := n
	var firstBadVersion int

	for leftIndex <= rightIndex {
		midIndex := leftIndex + (rightIndex-leftIndex)/2
		if isBadVersion(midIndex) {
			firstBadVersion = midIndex
			rightIndex = midIndex - 1
		} else {
			leftIndex = midIndex + 1
		}
	}

	return firstBadVersion
}

func SetFirstBadVersion(version int) {
	SavedFirstBadVersion = version
}

func isBadVersion(version int) bool {
	return version == SavedFirstBadVersion
}
