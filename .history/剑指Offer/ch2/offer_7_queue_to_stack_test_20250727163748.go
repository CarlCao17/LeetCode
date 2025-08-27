package ch2

import "testing"

func TestStack_Push(t *testing.T) {
	type fields struct {
		q1 []int
		q2 []int
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				q1: tt.fields.q1,
				q2: tt.fields.q2,
			}
			s.Push(tt.args.n)
		})
	}
}
