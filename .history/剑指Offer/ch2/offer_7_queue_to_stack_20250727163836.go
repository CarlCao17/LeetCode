package ch2

type Stack struct {
	q1 []int // 约定 q1 总是可以插入值
	q2 []int // q2 总是空，用于弹出使用
}

func (s *Stack) Push(n int) {
	s.q1 = append(s.q1, n)
}

func (s *Stack) Pop() int {
	if len(s.q1) == 0 {
		return 0
	}
	// assert len(s.q2) == 0
	s.q2 = append(s.q2, s.q1[:len(s.q1)-1]...)
	res := s.q1[len(s.q1)-1]
	s.q1 = s.q2
	s.q2 = nil
	return res
}
