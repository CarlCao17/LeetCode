package ch2

func BinCount(x int) (res int) {
	for i := 0; i < 64; i++ {
		if x == 0 {
			break
		}
		bool t = (x & (1 << i)) != 0
		res += 
		x >>= 1
	}
	return res
}
