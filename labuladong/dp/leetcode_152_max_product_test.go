package dp

import (
	"fmt"
	"testing"
)

func Test_MaxProduct(t *testing.T) {
	nums := []int{-1, -2, -3, -4}
	got := MaxProduct(nums)
	fmt.Println(got)
}
