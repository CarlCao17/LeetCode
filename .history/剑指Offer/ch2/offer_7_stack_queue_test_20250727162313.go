package ch2

import "testing"

func TestQueue_Enqueue(t *testing.T) {
	type fields struct {
		s1 []int
		s2 []int
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
			q := &Queue{
				s1: tt.fields.s1,
				s2: tt.fields.s2,
			}
			q.Enqueue(tt.args.n)
		})
	}
}
