package main

import "testing"

func Test_backspaceDel(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				"b#",
			},
			want: -1,
		},
		{
			args: args{
				"b###",
			},
			want: -1,
		},
		{
			args: args{
				"b###b##",
			},
			want: -1,
		},
		{
			args: args{
				"bxo#j##",
			},
			want: 0,
		},
		{
			args: args{
				"bxj##",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceDel(tt.args.s); got != tt.want {
				t.Errorf("backspaceDel() = %v, want %v", got, tt.want)
			}
		})
	}
}
