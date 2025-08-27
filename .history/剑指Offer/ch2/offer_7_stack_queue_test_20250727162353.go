package ch2

import "testing"

func TestQueue_Enqueue(t *testing.T) {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
}
