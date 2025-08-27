package ch2

import (
	"strings"
	"testing"
)

func (p *BinaryTreeNode) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	q := []*BinaryTreeNode{p}
	for len(q) > 0 {

	}
}

func TestBuildBinaryTree(t *testing.T) {
	btn := BuildBinaryTree(nil, nil)
	if btn != nil {
		t.Errorf("BuildBinaryTree(nil, nil) should be nil, got=%v\n", btn)
	}
	want := &BinaryTreeNode{
		Val: 1,
		Left: &BinaryTreeNode{
			Val: 2,
			Left: &BinaryTreeNode{
				Val: 4,
				Right: &BinaryTreeNode{
					Val: 7,
				},
			},
		},
		Right: &BinaryTreeNode{
			Val: 3,
			Left: &BinaryTreeNode{
				Val: 5,
			},
			Right: &BinaryTreeNode{
				Val: 6,
				Left: &BinaryTreeNode{
					Val: 8,
				},
			},
		},
	}
	got := BuildBinaryTree([]int{1, 2, 4, 7, 3, 5, 6, 8}, []int{4, 7, 2, 1, 5, 3, 8, 6})
	if 
}
