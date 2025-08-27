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
	swap(arr, pivot, hi)
	small := lo - 1
	for i := lo; i < hi; i++ {
		if arr[i] < arr[pivot] {
			small++
			swap(arr, small, i)
		}
	}
	small++
	swap(arr, small, pivot)
	return small
}
