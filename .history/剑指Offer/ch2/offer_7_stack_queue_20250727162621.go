package ch2

type Queue struct {
	s1 []int
	s2 []int
}

func (q *Queue) Enqueue(n int) {
	q.s1 = append(q.s1, n)
}

func (q *Queue) Dequeue() int {
	if len(q.s2) == 0 {
		q.s2 = append(q.s2, q.s1...)
		q.s1 = nil
	}
	res := q.s2[len(q.s2)-1]
	q.s2 = q.s2[:len(q.s2)-1]
	return res
}
