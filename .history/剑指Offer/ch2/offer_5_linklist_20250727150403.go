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
	l.Len++
	node := &ListNode{
		Val: n,
	}
	if l.tail != nil {
		l.tail.Next = node
	}
	l.tail = node
	if l.head == nil {
		l.head = node
	}
	return l
}

func (l *List) InsertHead(n int) *List {
	l.Len++
	node := &ListNode{Val: n}
	node.Next = l.head
	l.head = node
	if l.tail == nil {
		l.tail = node
	}
}

func (l *List) ToSlice() []int {
	res := make([]int, 0, l.Len)
	for p := l.head; p.Next != nil; p = p.Next {
		res = append(res, p.Val)
	}
	return res
}
