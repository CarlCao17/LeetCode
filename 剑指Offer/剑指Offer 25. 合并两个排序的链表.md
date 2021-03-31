# 剑指Offer 25. 合并两个排序的链表

### 错误解法

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
    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        if (l1 == null && l2 == null)
            return null;
        if (l1 == null)
            return l2;
        if (l2 == null)
            return l1;
        ListNode head = getMinNotNull(l1, l2);
        ListNode p = head; // pre node

        while (l1 != null && l2 != null) {
            p.next = getMinNotNull(l1, l2);
            p = p.next;
        }
        if (l1 != null)
            p.next = l1;
        if (l2 != null)
            p.next = l2;

        return head;
    }

    public ListNode getMinNotNull(ListNode l1, ListNode l2) {
        ListNode next;
        if (l1.val <= l2.val) {
            next = l1;
            l1 = l1.next;
        }
        else {
            next = l2;
            l2 = l2.next;
        }
        return next;
    }
}
```

显示Time Limit

为什么会超时呢？很容易想到可能是哪里没链好，导致死循环；

调试之后才发现：`l1`是局部变量，你即使传入到`getMinNotNull`之后修改它的值为`l1 = l1.next`，但是回到`mergeTwoLists`之后，`l1`的值仍为原来的值，所以**传参到函数中，然后直接修改引用变量的值并没有用**。



### 正确解法

### 链表小技巧

链表对于头部的处理比较麻烦，难道提取头部的时候又要比较一次`l1.val`和`l2.val`嘛？然后和循环体一样？再往下走一步嘛？

**trick**：

- 增加一个头结点，初始化的时候把头结点搞定，然后进入循环，最后返回头结点之后的节点即可；
- 链表如何链起来，访问当前节点**p**时，一定要找到它的下一个节点在哪，然后提前赋值: `p.next = xxx`（对于树也是类似的）

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
    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        ListNode dum = new ListNode(0); // pesum head node
        ListNode p = dum; // current node

        while (l1 != null && l2 != null) {
            if (l1.val <= l2.val) {
                p.next = l1;
                l1 = l1.next;
            }
            else {
                p.next = l2;
                l2 = l2.next;
            }
            p = p.next;
        }
        if (l1 != null)
            p.next = l1;
        if (l2 != null)
            p.next = l2;

        return dum.next;
    }
}
```

### 分析

时间复杂度： `O(m+n)`，`m`为`l1`的长度，`n`为`l2`的长度

空间复杂度：`O(1)`