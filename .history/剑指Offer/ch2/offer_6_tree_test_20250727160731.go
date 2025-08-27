package ch2

import (
	"strconv"
	"strings"
	"testing"
)

func (p *BinaryTreeNode) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	q := []*BinaryTreeNode{p}
	for len(q) > 0 {
		sz := len(q)
		allNil := true
		for i := 0; i < sz; i++ {
			if q[i] != nil {
				allNil = false
				break
			}
		}
		if allNil {
			break
		}
		for i := 0; i < sz; i++ {
			if sb.Len() > 1 {
				sb.WriteRune(',')
			}
			sb.WriteString(nodeToStr(q[i]))
			if q[i] != nil {
				q = append(q, q[i].Left)
				q = append(q, q[i].Right)
			}
		}
		q = q[sz:]
	}
	sb.WriteString("]")
	return sb.String()
}

// 最后一层叶子节点，全是 nil
func isLastleaveLevel(q []*BinaryTreeNode) bool {

}

func nodeToStr(p *BinaryTreeNode) string {
	if p == nil {
		return "null"
	}
	return strconv.Itoa(p.Val)
}

func TestBuildBinaryTree(t *testing.T) {
	btn := BuildBinaryTree(nil, nil)
	if btn != nil {
		t.Errorf("BuildBinaryTree(nil, nil) should be nil, got=%v\n", btn)
	}
	pre := []int{1, 2, 4, 7, 3, 5, 6, 8}
	in := []int{4, 7, 2, 1, 5, 3, 8, 6}
	got := BuildBinaryTree(pre, in)
	gotStr := got.String()
	want := "[1,2,3,4,null,5,6,null,7,null,null,8,null]"
	if gotStr != want {
		t.Errorf("BuildBinaryTree(%v, %v) should be %s, got=%s\n", pre, in, want, gotStr)
	}
}
