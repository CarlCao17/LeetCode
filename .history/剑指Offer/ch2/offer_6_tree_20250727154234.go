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
}
