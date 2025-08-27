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

}
