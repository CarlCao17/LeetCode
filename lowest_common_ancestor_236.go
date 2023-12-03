package main

import (
	. "leetcode/tree"
)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	parents := make(map[*TreeNode]*TreeNode)
	traverse(root, func(p *TreeNode) {
		if p.Left != nil {
			parents[p.Left] = p
		}
		if p.Right != nil {
			parents[p.Right] = p
		}
	})
	path := map[*TreeNode]bool{}
	for pp := p; pp != nil; pp = parents[pp] {
		path[pp] = true
	}
	for qq := q; qq != nil; qq = parents[q] {
		if _, ok := path[qq]; ok {
			return qq
		}
	}
	return nil // unreachable
}

func traverse(p *TreeNode, f func(*TreeNode)) {
	stack := []*TreeNode{}
	for p != nil || len(stack) > 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		if len(stack) == 0 {
			break
		}
		p = stack[0]
		stack = stack[1:]
		f(p)
		p = p.Right
	}
}
