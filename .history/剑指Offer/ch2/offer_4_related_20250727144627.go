package ch2

func MergeSortedInt(a1, a2 []int) []int {
	if cap(a1) < len(a1)+len(a2) {
		panic("")
	}
	a := a1[:cap(a1)]
	i, j := len(a1)-1, len(a2)-1
	k := len(a) - 1
	for i >= 0 && j >= 0 {
		if a1[i] < a2[j] {
			a[k] = a2[j]
			j--
		} else {
			a[k] = a1[i]
			i--
		}
		k--
	}
	for i >= 0 {
		a[k] = a1[i]
		i--
		k--
	}
	for j >= 0 {
		a[k] = a2[j]
		j--
		k--
	}

	return a
}
