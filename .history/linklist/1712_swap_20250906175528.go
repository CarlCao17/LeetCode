package linklist

func swapNodes(head *ListNode, k int) *ListNode {
	d := &ListNode{Next: head}
	pre := d
	for i := 0; i < k-1; i++ {
		pre = pre.Next
	}
	fast := head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	slow := head
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	swap(pre, slow)
	return d.Next
}

func swap(pPre, qPre *ListNode) {
	p, q := pPre.Next, qPre.Next
	t := q.Next

	q.Next = p.Next
	pPre.Next = q

	p.Next = t
	qPre.Next = p
}
