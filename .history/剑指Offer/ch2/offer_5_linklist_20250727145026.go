package ch2

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	Head ListNode
	Len  int
}

func NewList() List {
	return List{}
}
