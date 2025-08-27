package ch2

type BinaryTreeNode struct {
	Val         int
	Left, Right *BinaryTreeNode
}

func BuildBinaryTree(pre, in []int) *BinaryTreeNode {
	if len(pre) == 0 {
		return nil
	}
	rv := pre[0]
	btn := &BinaryTreeNode{
		Val: rv,
	}
	if len(pre) == 1 {
		return btn
	}
	idx := index(in, rv)
	btn.Left = BuildBinaryTree(pre[1:idx+1], in[:idx])
	btn.Right = BuildBinaryTree(pre[idx+1:], in[idx+1:])
	return btn
}

func index(a []int, b int) int {
	for i, aa := range a {
		if b == aa {
			return i
		}
	}
	return -1 // unreachable
}
