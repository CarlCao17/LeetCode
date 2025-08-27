package ch2

func BinCount(x int) (res int) {
	for i := 0; i < 64; i++ {
		res += x & (1 << i)
	}
	return res
}
