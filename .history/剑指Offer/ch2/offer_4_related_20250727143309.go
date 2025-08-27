package ch2

func mergeSortedInt(a1, a2 []int) []int {
	if cap(a1) < len(a1)+len(a2) {
		panic("")
	}
	a := a1[:cap(a1)]
	for i, j, k := len(a1)-1, len(a2)-1, cap(a1)-1; i >= 0 || j >= 0; k-- {

	}
}
