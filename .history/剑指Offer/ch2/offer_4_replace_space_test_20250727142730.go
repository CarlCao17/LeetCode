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
		{
			args: args{s: " aaa"},
			want: "%20aaa",
		},
		{
			args: args{s: "aaa  "},
			want: "aaa%20%20",
		},
		{
			args: args{s: " a  a  a "},
			want: "%%20a%20%20a%20%20a%20",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceSpace(tt.args.s); got != tt.want {
				t.Errorf("ReplaceSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
