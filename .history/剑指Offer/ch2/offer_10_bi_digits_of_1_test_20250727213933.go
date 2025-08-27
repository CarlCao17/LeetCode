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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := BinCount(tt.args.x); gotRes != tt.wantRes {
				t.Errorf("BinCount() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
