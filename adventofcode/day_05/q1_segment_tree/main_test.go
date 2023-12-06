package main

import (
	"testing"
)

// dst = src + 3 if in tree, otherwise dst = src
//
//	                        {*, 98, 2}
//	                        /
//	            {*, 50, 48}
//	            /
//	{*, 10, 25}
//	/           \
//
// {*, 0, 5}          {*, 35, 12}
//
//	\
//	 {*, 5, 3}
func initTree() *TreeNode {
	root := &TreeNode{srcToDst: srcToDst{98 + 3, 98, 2}}
	root.Insert(&TreeNode{srcToDst: srcToDst{50 + 3, 50, 48}})
	root.Insert(&TreeNode{srcToDst: srcToDst{10 + 3, 10, 25}})
	root.Insert(&TreeNode{srcToDst: srcToDst{35 + 3, 35, 12}})
	root.Insert(&TreeNode{srcToDst: srcToDst{0 + 3, 0, 5}})
	root.Insert(&TreeNode{srcToDst: srcToDst{5 + 3, 5, 3}})
	return root
}

func Test_treeNode_Insert(t *testing.T) {
	root := initTree()

	expect := &TreeNode{
		srcToDst: srcToDst{101, 98, 2},
	}
	expect.lChild = &TreeNode{srcToDst: srcToDst{53, 50, 48}}
	expect.lChild.lChild = &TreeNode{srcToDst: srcToDst{13, 10, 25}}
	expect.lChild.lChild.rChild = &TreeNode{srcToDst: srcToDst{38, 35, 12}}
	expect.lChild.lChild.lChild = &TreeNode{srcToDst: srcToDst{3, 0, 5}}
	expect.lChild.lChild.lChild.rChild = &TreeNode{srcToDst: srcToDst{8, 5, 3}}
	if root.String() != expect.String() {
		t.Errorf("got=%s, expect=%s", root.String(), expect.String())
	}
}

func Test_treeNode_search(t *testing.T) {
	// [0,4],[5,7],[10,34],[35,46],[50,97],[98,99]
	root := initTree()
	nilNode := (*TreeNode)(nil)
	ninetySevenNode := root.lChild
	fortyNode := root.lChild.lChild.rChild
	twentyTwoNode := root.lChild.lChild
	fourNode := root.lChild.lChild.lChild
	sevenNode := root.lChild.lChild.lChild.rChild

	if got := root.search(98); got != root {
		t.Errorf("got=%s, expect=%s", got, root)
	}
	if got := root.search(100); got != nilNode {
		t.Errorf("got=%s, expect=%s", got, nilNode)
	}
	if got := root.search(97); got != ninetySevenNode {
		t.Errorf("got=%s, expect=%s", got, ninetySevenNode)
	}
	if got := root.search(47); got != nilNode {
		t.Errorf("got=%s, expect=%s", got, nilNode)
	}
	if got := root.search(40); got != fortyNode {
		t.Errorf("got=%s, expect=%s", got, fortyNode)
	}
	if got := root.search(22); got != twentyTwoNode {
		t.Errorf("got=%s, expect=%s", got, twentyTwoNode)
	}
	if got := root.search(8); got != nilNode {
		t.Errorf("got=%s, expect=%s", got, nilNode)
	}
	if got := root.search(4); got != fourNode {
		t.Errorf("got=%s, expect=%s", got, fourNode)
	}
	if got := root.search(7); got != sevenNode {
		t.Errorf("got=%s, expect=%s", got, sevenNode)
	}
}

func Test_treeNode_Search(t *testing.T) {
	root := initTree()

	if got := root.Search(98); got != 101 {
		t.Errorf("got=%d, expect=%d", got, 101)
	}
	if got := root.Search(100); got != 100 {
		t.Errorf("got=%d, expect=%d", got, 100)
	}
	if got := root.Search(97); got != 100 {
		t.Errorf("got=%d, expect=%d", got, 100)
	}
	if got := root.Search(47); got != 47 {
		t.Errorf("got=%d, expect=%d", got, 47)
	}
	if got := root.Search(40); got != 43 {
		t.Errorf("got=%d, expect=%d", got, 43)
	}
	if got := root.Search(22); got != 25 {
		t.Errorf("got=%d, expect=%d", got, 25)
	}
	if got := root.Search(8); got != 8 {
		t.Errorf("got=%d, expect=%d", got, 8)
	}
	if got := root.Search(4); got != 7 {
		t.Errorf("got=%d, expect=%d", got, 7)
	}
	if got := root.Search(7); got != 10 {
		t.Errorf("got=%d, expect=%d", got, 10)
	}
}
