package dp

func MaxProduct(nums []int) int {
	dp := nums[0]
	rdp := nums[0]

	res := dp
	for _, n := range nums[1:] {
		t := dp * n
		rt := rdp * n
		dp = max(t, rt, n)
		rdp = min(t, rt, n)
		res = max(res, dp)
	}
	return res
}
