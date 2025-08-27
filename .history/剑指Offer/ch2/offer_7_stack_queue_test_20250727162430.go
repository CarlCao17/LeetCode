package ch2

import "testing"

func TestQueue_Enqueue(t *testing.T) {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	if t := q.Dequeue(); t != 3 {
		t.Errorf("got=%d, should be 3", t)
	}
}
