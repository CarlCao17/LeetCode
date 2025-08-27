package ch2

func MergeSortedInt(a1, a2 []int) []int {
	if cap(a1) < len(a1)+len(a2) {
		panic("")
	}
	a := a1[:cap(a1)]
	for i, j, k := len(a1)-1, len(a2)-1, cap(a1)-1; i >= 0 || j >= 0; k-- {
		if i < 0 {
			a[k] = a2[j]
			j--
		} else if j < 0 {
			a[k] = a1[i]
			i--
		} else {
			if a1[i] < a2[j] {
				a[k] = a2[j]
				j--
			} else {
				a[k] = a1[i]
				i--
			}
		}

	}
	return a
}
