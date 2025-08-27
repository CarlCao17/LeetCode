package dp

func MaxProduct(nums []int) int {
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
