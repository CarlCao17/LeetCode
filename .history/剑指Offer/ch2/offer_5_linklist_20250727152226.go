package ch2

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReversePrint(p *ListNode) []int {
	if p == nil {
		return []int{}
	}
	s := make([]int, 0, 8)
	for p != nil {
		s = append(s, p.Val)
		p = p.Next
	}
	res := make([]int, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		res[len(s)-1-i] = s[i]
	}
	return res
}
