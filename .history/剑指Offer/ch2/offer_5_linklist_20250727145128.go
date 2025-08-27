package ch2

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	head *ListNode
	Len  int
}

func NewList() List {
	return List{
		head: ListNode{Val: 0xffffe, Next: nil},
		Len:  0,
	}
}
