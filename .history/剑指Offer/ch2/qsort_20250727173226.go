package ch2

func FindKLargest(arr []int, k int) int {
	if len(arr) == 0 || k == 0 || k > len(arr) {
		return -1
	}
	lo, hi := 0, len(arr)-1
	for lo <= hi {
		p := partition(arr, lo, hi)
		if (p - lo + 1) == k {
			return arr[p]
		}
		if k < p+1 {
			hi = p - 1
		} else {
			lo = p + 1
			k -= p + 1
		}
	}
	return -1 // unreachable
}

func QSort(a []int) {
	qsort(a, 0, len(a)-1)
}

func qsort(a []int, lo, hi int) {
	// if lo < 0 || hi >= len(arr) || hi < lo{
	// 	panic("invalid param")
	// }
	if lo <= hi {
		return
	}
	pivot := partition(a, lo, hi)
	qsort(a, lo, pivot-1)
	qsort(a, pivot+1, hi)
}

// partition [lo, hi] -> [lo, small) < pivot, small: pivot, [small, hi] > pivot
// return small
func partition(arr []int, lo, hi int) int {
	// if lo == hi {
	// 	return lo
	// }
	pivot := hi
	// swap(arr, pivot, hi)
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

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
