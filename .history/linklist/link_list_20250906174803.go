package linklist

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(list []int) *ListNode {
	if len(list) == 0 {
		return nil
	}
	head := &ListNode{Val: list[0]}
	prev := head
	for _, num := range list[1:] {
		p := &ListNode{Val: num}
		prev.Next = p
		prev = p
	}
	return head
}

func (n *ListNode) String() string {
	var b strings.Builder
	b.WriteRune('[')
	for n != nil {
		b.WriteString(strconv.Itoa(n.Val))
		if n.Next != nil {
			b.WriteRune(' ')
		}

		n = n.Next
	}
	b.WriteRune(']')
	return b.String()
}
