package utils

import (
	"reflect"
	"testing"
)

func TestToTwoDimensionSlices(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{s: `[[5,4],[6,4],[6,7],[2,3]]`},
			want: [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}},
		},
		{
			args: args{s: `[[1,1],[1,1],[1,1]]`},
			want: [][]int{{1, 1}, {1, 1}, {1, 1}},
		},
		{
			args: args{s: `[]`},
			want: [][]int{},
		},
		{
			args: args{s: `[[]]`},
			want: [][]int{[]int{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToTwoDimensionSlices[int](tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToTwoDimensionSlices() = %v, want %v", got, tt.want)
			}
		})
	}
}
