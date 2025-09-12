package linklist

import "testing"

func Test_swapNodes(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	type testCase struct {
		args   args
		expect *ListNode
	}
	cases := []testCase{
		{
			args: args{
				head: arrayToLinkList([]int{1, 2, 3, 4, 5}),
				k:    2,
			},
			expect: arrayToLinkList([]int{1, 4, 3, 2, 5}),
		},
	}
	for i, cas := range cases {
		got := swapNodes(cas.args.head, cas.args.k)
		if !matchList(got, cas.expect) {
			t.Errorf("case %d: got=%s, expect=%s", i, got, cas.expect)
		}
	}
}

func matchList(p, q *ListNode) bool {
	if p == q || p.Val == q.Val {
		return matchList(p.Next, q.Next)
	}
	return false
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
