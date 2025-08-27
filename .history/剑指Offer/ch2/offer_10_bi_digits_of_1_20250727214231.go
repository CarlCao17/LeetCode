package ch2

func BinCount(x int) (res int) {
	for i := 0; i < 64; i++ {
		if x == 0 {
			break
		}
		res += x & (1 << i)
	}
	return res
}
