package ch2

import (
	"testing"
)

func TestBuildBinaryTree(t *testing.T) {
	btn := BuildBinaryTree(nil, nil)
	if btn != nil {
		t.Errorf("BuildBinaryTree(nil, nil) should be nil, got=%v\n", btn)
	}

}
