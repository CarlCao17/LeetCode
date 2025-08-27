package dp

func maxValueOf01Package(wt []int, vt []int, W int) int {
	if len(wt) == 0 || W == 0 {
		return 0
	}
	dp := make([]int, W+1)
	res := make([]int, W+1)
	for i := 0; i < len(wt); i++ {
		for j := 1; j <= W; j++ {
			res[j] = dp[j]
			if wt[i] <= j {
				res[j] = max(res[j], dp[j-wt[i]]+vt[i])
			}
		}
		copy(dp, res)
	}
	return dp[W]
}
