package ch2

import (
	"testing"
)

func TestList_Append(t *testing.T) {
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

	l.Append(2)

}
