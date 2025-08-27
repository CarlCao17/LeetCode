package ch2

// partition [lo, hi)
func partition(arr []int, lo, hi int) int {
	if len(arr) == 0 || lo < 0 || hi >= len(arr) || hi <= lo {
		panic("invalid param")
	}
	if lo+1 == hi {
		return lo
	}
	idx := hi
	arr[lo], arr[idx] = arr[idx], arr[lo]

}
