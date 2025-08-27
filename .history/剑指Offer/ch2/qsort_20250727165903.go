package ch2

// partition [lo, hi]
func partition(arr []int, lo, hi int) int {
	if len(arr) == 0 || lo < 0 || hi >= len(arr) || hi <= lo {
		panic("invalid param")
	}
	if lo+1 == hi {
		return lo
	}
	pivot := hi
	swap(arr, pivot, hi-1)
	small := lo - 1
	for i := lo; i < hi; i++ {
		if arr[i] < pivot {
			small++
			swap(arr, small, i)
		}
	}
}
