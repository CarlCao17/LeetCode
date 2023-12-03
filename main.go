package main

//
//import (
//	"fmt"
//
//	. "leetcode/linklist"
//)
//
//func main() {
//	node := NewListNode([]int{1, 2, 6, 3, 4, 5, 6})
//	fmt.Println(node)
//	fmt.Println(removeElements(node, 6))
//}
//
//func removeElements(head *ListNode, val int) *ListNode {
//	prev := &ListNode{Next: head} // dummy head node
//	p := prev
//	for p.Next != nil {
//		if p.Next.Val == val {
//			next := p.Next.Next
//			p.Next = next
//		} else {
//			p = p.Next
//		}
//	}
//	return prev.Next
//}
