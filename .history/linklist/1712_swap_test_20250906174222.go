package linklist

import "testing"

func Test_swapNodes(t *testing.T) {

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
