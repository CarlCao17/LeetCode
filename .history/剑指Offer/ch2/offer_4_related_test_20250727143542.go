package ch2

import (
	"reflect"
	"testing"
)

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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSortedInt(tt.args.a1, tt.args.a2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSortedInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
