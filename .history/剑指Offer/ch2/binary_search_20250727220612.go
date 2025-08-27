package ch2

func BinarySearch(a []int, n int) int {
	return binarySearchL(a, 0, len(a)-1, n)
}

func binarySearchL(a []int, lo, hi int, n int) int {
	if lo == hi {
		return lo
	}

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if n > a[mid] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
