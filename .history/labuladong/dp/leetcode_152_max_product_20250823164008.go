package dp

import "fmt"

func maxProduct(nums []int) int {
	dp := nums[0]
	rdp := nums[0]

	res := dp
	for _, n := range nums[1:] {
		dp = max(dp*n, max(rdp*n, n))
		rdp = min(dp*n, min(rdp*n, n))
		res = max(res, dp)
	}
	return res
}

func main() {
	nums := []int{-1, -2, -3, -4}
	got := maxProduct(nums)
	fmt.Println(got)
}
