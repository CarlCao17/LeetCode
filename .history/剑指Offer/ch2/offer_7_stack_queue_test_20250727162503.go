package ch2

import "testing"

func TestQueue_Enqueue(t *testing.T) {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	if got := q.Dequeue(); got != 1 {
		t.Errorf("got=%d, should be 1", got)
	}
	if got := q.Dequeue(); got != 2 {
		t.Errorf("got=%d, should be 2", got)
	}
}
