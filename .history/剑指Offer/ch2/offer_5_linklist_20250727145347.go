package ch2

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	head *ListNode
	tail *ListNode
	Len  int
}

func NewList() *List {
	return &List{}
}

// func FromSlice(a []int) *List {
// 	// l := NewList()
// 	// for _
// }

func (l *List) Append(n int) *List {

}
