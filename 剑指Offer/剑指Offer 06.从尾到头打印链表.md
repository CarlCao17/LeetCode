# 剑指Offer 06.从尾到头打印链表

### 使用栈来逆序

### 代码

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
    public int[] reversePrint(ListNode head) {
        LinkedList<Integer> stack = new LinkedList<>();
        ListNode p = head;

        while (p != null) {
            stack.push(p.val);
            p = p.next;
        }
        int[] res = new int[stack.size()];
        int i = 0;
        while (!stack.isEmpty()) {
            res[i] = stack.pop();
            i++;
        }
        return res;
    }
}
```

### 将链表逆序然后存储到数组中

```
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public int[] reversePrint(ListNode head) {
        if (head == null)
            return new int[0];
        int cnt = 1;
        ListNode pre = head, curr = head.next;
        head.next = null;
        while (curr != null) {
            cnt++;
            ListNode temp = curr.next;
            curr.next = pre;
            pre = curr;
            curr = temp;
        }
        int[] res = new int[cnt];
        int i = 0;
        while (pre != null) {
            res[i] = pre.val;
            pre = pre.next;
            i++;
        }
        return res;
    }
}
```