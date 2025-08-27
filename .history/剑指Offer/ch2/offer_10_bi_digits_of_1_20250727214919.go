package ch2

func BinCount(x int) (res int) {
	for x != 0 {
		if x&1 != 0 {
			res++
		}

	}
	return res
}
