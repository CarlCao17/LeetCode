# [剑指 Offer 62. 圆圈中最后剩下的数字](https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/)

约瑟夫环问题很容易想到用循环链表实现，每次找到第`m`个元素就删去并跳转到下一个元素，继续找直到链表中仅剩一个元素。

```java
class Solution {
    private class Node {
        int val;
        Node next;
        public Node(int val) { this.val = val; }
    }
    
    public int lastRemaining(int n, int m) {
        Node head = new Node(0);
        Node p = head;
        Node prev;
        for (int i = 1; i < n; i++) {
            p.next = new Node(i);
            p = p.next;
        }
        p.next = head;
        prev = p;
        p = head;
        
        while (p.next != p) {
            int k = 1;
            while (k < m) { // 走m-1步
                p = p.next;
                prev = prev.next;
                k++;
            }
            p = removeNode(prev);
        }
        return p.val;
    }

    public Node removeNode(Node prev) {
        Node o = prev.next;
        prev.next = o.next;
        o.next = null;
        return prev.next;
    }
}
```

但很可惜超时了~，在`case: n = 70866 m = 116922`就超时了

**时间复杂度**$O(mn)$：因为每一次删除一个元素都需要走$m-1$步，而当且仅当删除了$n-1$个元素才会停下来，而

`m <= 10^5, n <= 10^6`当然可能会超时

因此应当寻找$O(n)$级别的算法，才不会超时，同时也可以用双端队列来实现：

### 双端队列

当然由于双端队列比循环链表还需要做更多的操作，因此肯定会超时

```java
 public int lastRemaining(int n, int m) {
        Deque<Integer> queue = new LinkedList<>();
        for (int i = 0; i < n; i++) {
            queue.addLast(i);
        }
        int number = 1;
        while (queue.size() > 1) {
            Integer item = queue.removeFirst();
            if (number % m == 0) {
                number = 1;
                continue;
            }
            queue.addLast(item);
            number++;
        }
        return queue.removeFirst();
    }
}
```





## 递推方法

要找长度为`n`的圈，每次删去第`m`个数剩下的是那个数字，可以记为问题`f(n, m)`。很显然它和更小的问题`f(n-1,m)`有关。第一次删去的元素`m%n`，剩下长为`n-1`的圈，则若假设`f(n-1, m)`最终剩下的是在`n-1`的圈中第`x`个元素，则第`x`个元素在原来的序列中应该是`f(n, m) = (x + m%n) % n = (x + m) % n` 

![fig1](https://assets.leetcode-cn.com/solution-static/jianzhi_62_fig1.gif)

```java
class Solution {
    public int lastRemaining(int n, int m) {
        int last = 0; // f(1,m) = 0
        int res = 0;
        for (int i = 2; i <= n; i++) {
            res = (last + m ) % i; // f(i, m) = (f(i-1, m) + m) % i
            last = res;
        }
        return res;
    }
}
```

时间复杂度$O(n)$