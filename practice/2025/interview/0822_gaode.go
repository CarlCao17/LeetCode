package interview

func FindLargestProduct(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	dp, rdp := nums[0], nums[0]
	for _, n := range nums[1:] {
		t := dp * n
		rt := rdp * n
		dp = max(t, n) // dp 没有计算 rdp
		rdp = min(rt, n)
	}
	return max(dp, rdp) // ❌：res 没有更新
}

func maxProduct(nums []int) int {
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

func main() {
	cases := [][]float64{
		{}
	}
	expect := []float64{

	}
	for i, case := range cases {
		got := FindLaFindLargestProduct(case)
		if got != expect[i] {
			fmt.Println()
		}
	}
}
