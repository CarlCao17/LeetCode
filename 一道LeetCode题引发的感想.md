

# 一道LeetCode题引发的感想

有时候并不是将所有情况概述在一起，代码优美更好；反而是直接了当解决，速度/性能更好

## LeetCode 2 Add Two Numbers

You are given two **non-empty** linked lists representing two non-negative integers. The digits are stored in **reverse order** and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

**Example:**

```
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
```

Solution 1(更优美)

创建新链表，保存结果； 使用原来链表

```python
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        carry, sums = divmod(l1.val + l2.val, 10)
        head = p = ListNode(sums)
        pre = p
        p1, p2 = l1.next, l2.next
        while p1 or p2:
            sums = 0
            if p1:
                sums += p1.val
            if p2:
                sums += p2.val
            sums += carry
            carry, sums = divmod(sums, 10)
            p = ListNode(sums)
            pre.next = p
            pre = p
            if p1:
                p1 = p1.next
            if p2:
                p2 = p2.next
        if carry:
            p = ListNode(carry)
            pre.next = p
            pre = p
            
        return head
  
```

```python
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        """
        carry: represents the carry number of the addition
        sums: the first-place of the sum
        For example:5+7=12, then, carry = 1, sums = 2
        head: head node of the new linked list 
        pre: the pre-node or the node before the current processing node
        We will save the result in l1's node.
        """
        # preprocess for the first node
    
        carry, sums = divmod(l1.val + l2.val, 10)
        l1.val = sums
        head = pre = l1
        l1, l2 = l1.next, l2.next
        
        while l1 or l2:
            sums = 0
            # If l1 exists, then save the sums in the l1, otherwise, turn to another list's current node l2 (l1 = l2)
            # And there are 3 situations: l1 exists, l2 exists; l1 exists, l2 not; l1 not, l2 exists
            if (l1 and l2) and (l1 is not l2): # the first situation
                sums += l2.val
            if not l1:
                l1 = l2
            sums += l1.val
            
            sums += carry
            carry, sums = divmod(sums, 10)
            l1.val = sums
            pre.next = l1
            pre = l1
            
            l1 = l1.next
            if l2:
                l2 = l2.next
        if carry:
            p = ListNode(carry)
            pre.next = p
            pre = p
            
        return head
            
```

![image-20200825202207558](/Users/caozhengcheng/Documents/knowledge base/pic/image-20200825202207558.png)

Solution 2(代码更长，复情况)

```python
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        
        head = l1
        carry = 0
        while l1 and l2:
            p = l1 # the current working node
            carry, sums = divmod(l1.val + l2.val + carry, 10)
            p.val = sums
            l1, l2 = l1.next, l2.next # l1 be the next node of p
        while l1:
            p = l1
            carry, sums = divmod(l1.val + carry, 10)
            p.val = sums
            l1 = l1.next
        while l2:
            p.next = l2
            p = l2
            carry, sums = divmod(l2.val + carry, 10)
            p.val = sums
            l2 = l2.next
        if carry:
            p.next = ListNode(carry)
            
        return head
     
```

![image-20200825201919402](/Users/caozhengcheng/Documents/knowledge base/pic/image-20200825201919402.png)