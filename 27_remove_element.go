package main

func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)
	for left < right {
		if nums[left] == val {
			right--
			swap(nums, left, right)
		} else {
			left++
		}
	}
	return left
}

func swap(nums []int, a, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}

//
//func main() {
//	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
//	for _, v := range nums[:removeElement(nums, 2)] {
//		fmt.Printf("%d ", v)
//	}
//}
