package ch2

func IsPowerOf2(n int) bool {
	return n&(n-1) == 0
}

func BinCount(n int) (res int) {
	x := uint(n)
	flag := uint(1)
	for i := 0; x != 0 && i < 64; i++ {
		if x&flag != 0 {
			res++
		}
		flag <<= 1
	}
	return res
}

func BinCount2(n int) (res int) {
	x := uint(n)
	for x != 0 {
		res++
		x = (x - 1) & x
	}
	return res
}
