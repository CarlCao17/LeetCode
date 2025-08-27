package dp

import "fmt"

func numDistinct(s string, t string) (ans int) {
	m, n := len(s), len(t)
	if m < n || m == 0 || n == 0 {
		return 0
	}
	dp := make([][]int, m)
	dp[0] = make([]int, n)
	if s[0] == t[0] {
		dp[0][0] = 1
	}
	for i := 1; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = dp[i-1][0]
		if s[i] == t[0] {
			dp[i][0] += 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if j > i {
				break
			}
			dp[i][j] = dp[i-1][j]
			if s[i] == t[j] {
				dp[i][j] += dp[i-1][j-1]
			}
		}
	}
	fmt.Println(dp)
	return dp[m-1][n-1]
}
