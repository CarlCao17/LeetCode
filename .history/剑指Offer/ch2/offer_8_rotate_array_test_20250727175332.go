package ch2

import "testing"

func TestFindSmallestNumInRotateArray(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "corner case: no rotate",
			args: args{a: []int{1, 2, 3, 4, 5}},
			want: 1,
		},
		{
			name: "corner case: fully rotate",
			args: args{a: []int{5, 4, 3, 2, 1}},
			want: 1,
		},
		{
			name: "corner case: no rotate, non-strictly increasing",
			args: args{a: []int{1, 1, 2, 2, 3, 3}},
			want: 1,
		},
		{
			name: "corner case: fully rotate, non-strictly increasing",
			args: args{a: []int{3, 3, 2, 2, 1, 1}},
			want: 1,
		},
		{
			name: "rotate 1",
			args: args{a: []int{2, 3, 4, 5, 1}},
			want: 1,
		},
		{
			name: "rotate 2",
			args: args{a: []int{3, 4, 5, 1, 2}},
			want: 1,
		},
		{
			name: "rotate 3",
			args: args{a: []int{4, 5, 1, 2, 3}},
			want: 1,
		},
		{
			name: "rotate 4",
			args: args{a: []int{5, 1, 2, 3, 4}},
			want: 1,
		},
		{
			name: "corner case: rotate 1, non-strictly increasing",
			args: args{a: []int{1, 2, 2, 3, 3, 1}},
			want: 1,
		},
		{
			name: "corner case: rotate 2, non-strictly increasing",
			args: args{a: []int{2, 2, 3, 3, 1, 1}},
			want: 1,
		},
		{
			name: "corner case: rotate 3, non-strictly increasing",
			args: args{a: []int{2, 3, 3, 1, 1, 2}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindSmallestNumInRotateArray(tt.args.a); got != tt.want {
				t.Errorf("FindSmallestNumInRotateArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
