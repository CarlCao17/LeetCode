package ch2

type Queue struct {
	s1 []int
	s2 []int
}

func (q *Queue) Append(n int) {
	q.s1 = append(q.s1, n)
}

func (q *Queue) DeleteHead() int {

}
