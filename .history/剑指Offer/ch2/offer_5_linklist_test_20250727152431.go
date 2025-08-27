package ch2

import (
	"reflect"
	"testing"
)

func ToList(a []int) *ListNode {
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

func TestToList(t *testing.T) {
	l := NewList()
	l.Append(1)
	if l.Len != 1 {
		t.Errorf("assert l.Len: want=1, got=%d\n", l.Len)
	}
	if l.head == nil || l.head.Val != 1 || l.head.Next != nil {
		t.Errorf("assert l.head: got=%v\n", *l.head)
	}
	if l.head != l.tail {
		t.Errorf("assert l.head == l.tail: tail=%v\n", l.tail)
		if l.tail != nil {
			t.Errorf("l.tail: %v\n", *l.tail)
		}
	}

	l.Append(2).Append(3)
	got := l.ToSlice()
	want := []int{1, 2, 3}
	if reflect.DeepEqual(got, want) {
		t.Errorf("assert l: want=%v, got=%v\n", want, got)
	}
	l.InsertHead(4)
	want = []int{4, 1, 2, 3}
	got = l.ToSlice()
	if l.Len != 4 {
		t.Errorf("assert l.Len: want=%d, got=%d\n", 4, got)
	}
	if reflect.DeepEqual(got, want) {
		t.Errorf("assert l: want=%v, got=%v\n", want, got)
	}
}
