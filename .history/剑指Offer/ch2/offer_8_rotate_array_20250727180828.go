package ch2

func FindSmallestNumInRotateArray(a []int) int {
	// len(a) == 0 ?
	if len(a) == 1 {
		return a[0]
	}
	l, r := 0, len(a)-1
	mid := l
	for a[l] >= a[r] {
		if r-l == 1 {
			mid = r
			break
		}
		mid = (l + r) / 2
		if a[mid] == a[r] && a[mid] == a[l] {
			return minInOrder(a, l, r)
		}
		if a[mid] >= a[l] {
			l = mid
		} else if a[mid] <= a[r] {
			r = mid
		}
	}
	return a[l]
}

func minInOrder(a []int, l, r int) int {

}
