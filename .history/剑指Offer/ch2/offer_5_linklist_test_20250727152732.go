package ch2

import (
	"reflect"
	"testing"
)

func FromArray(a []int) *ListNode {
	if len(a) == 0 {
		return nil
	}
	d := &ListNode{}
	prev := d
	for _, aa := range a {
		n := &ListNode{Val: aa}
		prev.Next = n
		prev = prev.Next
	}
	return d.Next
}

func (l *ListNode) ToArray() []int {
	if l == nil {
		return []int{}
	}
	res := make([]int, 0, 8)
	for p := l; p != nil; p = p.Next {
		res = append(res, p.Val)
	}
	return res
}

func TestListNodeHelper(t *testing.T) {
	list := FromArray([]int{1, 2, 3, 4, 5, 6, 7})
	got := list.ToArray()
	if !reflect.DeepEqual(got, []int{1, 2, 3, 4, 5, 6, 7}) {

	}
}
