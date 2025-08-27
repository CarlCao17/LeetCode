package ch2

func BinCount(n int) (res int) {
	x := uint(n)
	flag := uint(1)
	for i := 0; i < 64; i++ {
		if x&flag != 0 {
			res++
		}
		flag <<= 1
	}
	return res
}
