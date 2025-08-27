package ch2

import (
	"reflect"
	"testing"
)

func TestBuildBinaryTree(t *testing.T) {
	type args struct {
		pre []int
		in  []int
	}
	tests := []struct {
		name string
		args args
		want *BinaryTreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildBinaryTree(tt.args.pre, tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
