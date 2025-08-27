package ch2

import "testing"

func TestReplaceSpace(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceSpace(tt.args.s); got != tt.want {
				t.Errorf("ReplaceSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
