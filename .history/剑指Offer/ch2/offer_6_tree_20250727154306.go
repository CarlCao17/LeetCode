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

func index(a []int, b int) int {
	for i, aa := range a {
		if b == aa {
			return i
		}
	}
	return -1 // unreachable
}
