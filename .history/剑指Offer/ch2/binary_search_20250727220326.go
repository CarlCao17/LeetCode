package ch2

func BinarySearch(a []int, n int) int {
	return binarySearch(a, 0, len(a)-1, n)
}

func binarySearch(a []int, lo, hi int, n int) int {
	if lo == hi {
		return lo
	}
}
