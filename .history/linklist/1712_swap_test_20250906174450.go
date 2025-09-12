package linklist

import "testing"

func Test_swapNodes(t *testing.T) {
	type args struct {
		head *ListNode
		k int
	}
	type testCase struct {
		args args
		expect *ListNode
	}
	cases := []testCase {
		{
			args: args{
				head: arrayToLinkList([]int{1,2,3,4,5}),
			k: 2,
			}
		},
	}
}

func arrayToLinkList(nums []int) *ListNode {
	d := &ListNode{}
	pre := d
	for _, num := range nums {
		pre.Next = &ListNode{Val: num}
		pre = pre.Next
	}
	return d.Next
}
