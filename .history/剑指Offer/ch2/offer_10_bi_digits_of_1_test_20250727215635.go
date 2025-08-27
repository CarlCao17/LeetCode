package ch2

import "testing"

func TestBinCount(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{
			args:    args{x: 0},
			wantRes: 0,
		},
		{
			args:    args{x: 1},
			wantRes: 1,
		},
		{
			args:    args{x: 2},
			wantRes: 1,
		},
		{
			args:    args{x: 11},
			wantRes: 3,
		},
		{
			args:    args{x: 100},
			wantRes: 3,
		},
		{
			args:    args{x: 99999},
			wantRes: 10,
		},
		{
			args:    args{x: -1},
			wantRes: 64,
		},
		{
			args:    args{x: 0x7FFFFFFFFFFFFFFF},
			wantRes: 63,
		},
		{
			args:    args{x: 0x800000000000000},
			wantRes: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := BinCount(tt.args.x); gotRes != tt.wantRes {
				t.Errorf("BinCount() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
