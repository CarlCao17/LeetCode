package ch2

func partition(arr []int, lo, hi int) int {
	if len(arr) == 0 || lo < 0 || hi > len(arr) || hi > lo {
		panic("invalid param")
	}
	if lo == hi {
		return lo
	}
	idx := 
}
