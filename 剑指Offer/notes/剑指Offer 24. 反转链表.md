# 剑指Offer 24. 反转链表

### 方法一 头插法

```java
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode reverseList(ListNode head) {
        ListNode dum = new ListNode(0);
        dum.next = null;

        // 头插法
        ListNode p = head;
        while (p != null) {
            ListNode temp = p.next;
            p.next = dum.next;
            dum.next = p;
            p = temp;
        }
        return dum.next;
    }
    
}
```


### 方法二 利用双指针反转

```java
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode reverseList(ListNode head) {
        ListNode pre = null, curr = head;

        while (curr != null) {
            ListNode temp = curr.next;
            curr.next = pre;
            pre = curr;
            curr = temp;
        }
        return pre;
    }
    
}
```

### 方法三 递归

递归版本稍微复杂一些，其关键在于反向工作。假设链表的其余部分已经被反转，现在应该如何反转它前面的部分？

假设链表为：

$n_1\rightarrow \ldots \rightarrow n_{k-1} \rightarrow n_k \rightarrow n_{k+1} \rightarrow \ldots \rightarrow n_m \rightarrow \varnothing$

若从节点$ n_{k+1}$到 $n_m$已经被反转，而我们正处于 $n_k$,

$n_1\rightarrow \ldots \rightarrow n_{k-1} \rightarrow n_k \rightarrow n_{k+1} \leftarrow \ldots \leftarrow n_m$
我们希望 $n_{k+1}$的下一个节点指向 $n_k$
所以，$n_k.\textit{next}.\textit{next} = n_k $

需要注意的是 $n_1$的下一个节点必须指向 $\varnothing$。如果忽略了这一点，链表中可能会产生环。

```java
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode reverseList(ListNode head) {
        if (head == null || head.next == null)
            return head;
        ListNode newHead = reverseList(head.next);
        head.next.next = head;
        head.next = null;
        return newHead;
    }
```

