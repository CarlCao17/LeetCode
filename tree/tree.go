package tree

import (
	"strconv"
	"strings"
)

func ToTree(s string) *TreeNode {
	return NewTree(ToSlice(s))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTree(s []string) *TreeNode {
	var res *TreeNode
	if len(s) == 0 {
		return res
	}
	root := NewTreeNode(s[0])
	queue := []*TreeNode{root}

	for j := 1; j < len(s); {
		n := len(queue)
		for i := 0; i < n; i++ {
			// process left
			if j >= len(s) {
				break
			}
			queue[i].Left = NewTreeNode(s[j])
			j++
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			// process right
			if j >= len(s) {
				break
			}
			queue[i].Right = NewTreeNode(s[j])
			j++
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[n:]
	}
	return root
}

func NewTreeNode(s string) *TreeNode {
	if s == "null" {
		return nil
	}
	return &TreeNode{Val: ToInt(s)}
}

func ToSlice(s string) []string {
	s = strings.ReplaceAll(s, " ", "") // del all space
	if s[0] != '[' || s[len(s)-1] != ']' {
		panic("invalid slice string, not start with '[' or end with ']': " + s)
	}
	if len(s) == 2 {
		return []string{}
	}
	return strings.Split(s[1:len(s)-1], ",")
}

func ToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
