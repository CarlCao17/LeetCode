package ch2

func BinarySearch(a []int, n int) int {
	return binarySearchL(a, 0, len(a)-1, n)
}

func binarySearchL(a []int, lo, hi int, n int) int {
	l, r := lo, hi
	for l < r {
		mid := l + (r-l)/2
		if n < a[mid] {
			r = mid - 1
		} else if n > a[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if l == hi && n > a[hi] {
		return hi + 1
	}
	return l
}

func binarySearchR(a []int, lo, hi int, n int) int {
	l, r := lo, hi
	for l < r {
		mid := l + (r-l)/2
		if n < a[mid] {
			r = mid - 1
		} else {
			l = mid
		}
	}
	return lo
}
