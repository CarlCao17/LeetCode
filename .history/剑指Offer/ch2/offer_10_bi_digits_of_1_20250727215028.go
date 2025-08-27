package ch2

func BinCount(n int) (res int) {
	x := uint(n)
	for x != 0 {
		if x&1 != 0 {
			res++
		}
		x >>= 1
	}
	return res
}
