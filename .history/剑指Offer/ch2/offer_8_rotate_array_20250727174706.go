package ch2

func FindSmallestNumInRotateArray(a []int) int {
	// len(a) == 0 ?
	if len(a) == 1 {
		return a[0]
	}
	l, r := 0, len(a)-1
	for l < r {
		mid := l + (r-l)/2
		if a[mid] > a[r] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
}
