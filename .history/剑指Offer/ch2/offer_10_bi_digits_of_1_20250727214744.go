package ch2

func BinCount(x int) (res int) {
	for i := 0; i < 64; i++ {
		if x == 0 {
			break
		}
		if (x & 1) != 0 {
			res += 1
		}
		x >>= 1
	}
	return res
}
