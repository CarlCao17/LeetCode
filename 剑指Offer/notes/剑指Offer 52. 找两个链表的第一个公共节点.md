# 剑指Offer 52. 找两个链表的第一个公共节点

## 方法一. 常规做法

先确定两个链表的长度，然后让指向长链表的指针先走，走到某个节点，使得从该节点到长链表尾的距离等于锻炼表的长度，然后两指针再同步走，一直走到底，如果走到底则说明不想交，否则在之前应该已经相遇。

```java
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) {
 *         val = x;
 *         next = null;
 *     }
 * }
 */
public class Solution {
    public ListNode getIntersectionNode(ListNode headA, ListNode headB) {
        ListNode p = headA, q = headB;
        // 找链表长度
        int m = 0, n = 0;
        while (p != null) {
            m++;
            p = p.next;
        }
        while (q != null) {
            n++;
            q = q.next;
        }
        // 让长链表指针先走到预定位置，使之后的长度相同
        p = headA;
        q = headB;
        while (m > n) {
            p = p.next;
            m--;
        }
        while (n > m) {
            q = q.next;
            n--;
        }
        // 比较判断，要么一起走到null，要么找到相同节点，注意要用!= 不能用equals
        while (p != null  && p != q) {
            p = p.next;
            q = q.next;
        }
        if (p == null)
            return null;
        return p;
    }
}
```

时间复杂度 $O(m+n)$ 最多两个指针各走两遍整个链表

空间复杂度 $O(1)$

## 方法二 浪漫相遇

假设一个链表 $A$ 长为$m$，另一个链表$B$ 长为$n$，如果相交，假设相交部分长为$k\ (k <= m, k <= n)$ 。让两个链表指针同步跑，如果有某一个指针先走到尾部（即$null$)则让它再跑到对面的链表头部，再跑，这样不论相交或是不想交，总会找到某一个点使得两个指针 $p$ 和 $q$ 相遇（要么是在 $ null$处，要么是在相交点）。为方便叙述和理解，假设第一次赋值 $p = headA, q = headB$为第1步：

- **相交**：不妨第一个链表更长一些（即 $m > n$ )，则指针 $q$ 先走到链表尾部，然后转向第一个链表头部继续前进，同样到某个时候指针 $p$ 也走到链表尾部再转向第二个链表头部继续前进。当走到第 $m+n-k$ 步时，必定相遇。

   **原因:** （如果链表长度不相同）
  当走到第$n+1$步时，$q$本来指向$null$, 我们让它转向 $q = headA$，而$p$ 还在链表$A$上的第 $n+1$ 个位置。
  然后当走到第$m+1$步时，$p$要转向，$p = headB$ 即$p$在链表$B$的1号位置 ，$q$ 在链表$A$上的第$m-n+1$个位置 ($m+1 - (n+1) + 1 = m-n+1$)，$p$再往前走$n-k-1$步，$p$走到了链表$B$的第$n-k$个位置，即相交点。
  同理，$q$ 再往前走$n-k-1$步也走到链表A的第$m-k$个位置，即相交点。即$p$和$q$都走了 $m+n-k$步，到达相交点（如果还是很难理解，照着我的描述画个图即可）

  （如果链表长相同）

  那就简单了，只要走一遍，就会同时走到相交点

- 不想交：比较简单，两个指针分别走到链表尾，如果链表长相等则同时走到链表尾，如果不相等则先走一遍自己的链表再走一遍对面的链表，总共走$m+n$步同时再次来到链表尾

因此，它们不论链表长是否相同，不论是否会相交，终究会走到同时相交点或者链表尾，可以利用$p\space !=\space q$作为循环条件，$p == null$来判断是否相交

```java
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) {
 *         val = x;
 *         next = null;
 *     }
 * }
 */
public class Solution {
    public ListNode getIntersectionNode(ListNode headA, ListNode headB) {
        ListNode p = headA, q = headB;
        while (p != q) {
            p = p != null ? p.next : headB;
            q = q != null ? q.next : headA;
        }
        return p;
    }
}
```

时间复杂度$O(m+n)$ 最多走$m+n$步，但是方法二的系数（1）比方法一（2）要小，而且更优美

空间复杂度$O(1)$ 只使用了两个指针，占用常量空间

打败100%