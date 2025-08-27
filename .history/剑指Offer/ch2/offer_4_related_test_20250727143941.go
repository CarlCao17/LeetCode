package ch2

import (
	"reflect"
	"testing"
)

func genA1(l, c int) []int {
	res := make([]int, l, c)
	for i := 0; i < l; i++ {
		res[i] = i+1
	}
	return res
}

func TestMergeSortedInt(t *testing.T) {
	type args struct {
		a1 []int
		a2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty a1",
			args: args{
				a1: genA1(0, 4),
				a2: []int{1,2,3,4},
			},
			want: []int{1,2,3,4},
		},
		{
			name: "empty a2",
			args: args{
				a1: genA1(3, 3),
				a2: []int{},
			},
			want: []int{1,2,3},
		},
		{
			name: "intersect: a1 shorter than a2",
			args: args{
				a1: genA1(2,5),
				a2: []int{1,2,3},
			},
			want: []int{1,1,2,2,3},
		},
		{
			name: ""
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSortedInt(tt.args.a1, tt.args.a2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSortedInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
