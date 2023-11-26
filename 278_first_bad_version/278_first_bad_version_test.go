package _278_first_bad_version

import "testing"

func TestFirstBadVersion(t *testing.T) {
	testCases := []struct {
		name               string
		firstBadVersion    int
		versionToTest      int
		expectedBadVersion int
	}{
		{"FirstBadVersion4", 4, 5, 4},
		{"FirstBadVersion1", 1, 1, 1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			SetFirstBadVersion(tc.firstBadVersion)
			result := FirstBadVersion(tc.versionToTest)
			if result != tc.expectedBadVersion {
				t.Errorf("Expected %d, got %d", tc.expectedBadVersion, result)
			}
		})
	}
}
