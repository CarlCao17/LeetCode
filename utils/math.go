package utils

func Max[T Integer | Float](a, b T) T {
	return MaxN(a, b)
}

func Max3[T Integer | Float](a, b, c T) T {
	return MaxN(a, b, c)
}

func Min[T Integer | Float](a, b T) T {
	return MinN(a, b)
}

func Min3[T Integer | Float](a, b, c T) T {
	return MinN(a, b, c)
}

func MaxN[T Integer | Float](nums ...T) T {
	if len(nums) == 0 {
		panic("MaxN should have nums")
	}
	res := nums[0]
	for _, num := range nums[1:] {
		if num > res {
			res = num
		}
	}
	return res
}

func MinN[T Integer | Float](nums ...T) T {
	if len(nums) == 0 {
		panic("MinN should have nums")
	}
	res := nums[0]
	for _, num := range nums[1:] {
		if num < res {
			res = num
		}
	}
	return res
}
