# 剑指Offer 35. 复杂链表的复制

[复制复杂链表][https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof/]

在复杂链表中，每一个节点除了有一个`next`指针指向下一个节点，还要一个`random`指针指向链表中的任意节点或者`null`。

## 链表——初始解法

把它当做链表来复制，如果没有`random`指针很好处理。多了`random`需要多一遍遍历，由于`random`可能指向之后的节点，所以必须先将链表深拷贝一遍，然后再单独处理`random`域。对于`random`域最主要的问题在于，如何已知原链表的`random`域快速找到拷贝后的链表中对应的节点。由于原链表和拷贝后的链表一模一样，所以可以利用中间的索引建立对应关系：`oldNodeMap`将`Node`映射到它在链表中的索引，`newIndexMap`利用从`oldNodeMap`中查到的索引对应到新拷贝链表中的元素。

**关键步骤**

`newCurr.random = newIndexMap.get(oldNodeMap.get(oldCurr.random));`

#### 代码一

```java
/*
// Definition for a Node.
class Node {
    int val;
    Node next;
    Node random;

    public Node(int val) {
        this.val = val;
        this.next = null;
        this.random = null;
    }
}
*/
class Solution {
    public Node copyRandomList(Node head) {
        if (head == null)
            return null;
        Map<Integer, Node> newIndexMap = new HashMap<Integer, Node>();
        Map<Node, Integer> oldNodeMap = new HashMap<Node, Integer>();

        Node dumHead = new Node(0);
        Node pre = dumHead;
        Node newCurr, oldCurr = head;
        int index = 0;
        while (oldCurr != null) {
            newCurr = new Node(oldCurr.val);
            pre.next = newCurr;
            
            newIndexMap.put(index, newCurr);
            oldNodeMap.put(oldCurr, index);
            oldCurr = oldCurr.next;
            pre = newCurr;
            index++;
        }
        oldNodeMap.put(null, -1);
        newIndexMap.put(-1, null);

        oldCurr = head;
        newCurr = dumHead.next;
        while (oldCurr != null) {
            newCurr.random = newIndexMap.get(oldNodeMap.get(oldCurr.random));

            oldCurr = oldCurr.next;
            newCurr = newCurr.next;
        }
        return dumHead.next;
    }
}
```

时间复杂度`O(N)` : 两遍遍历链表

空间复杂度`O(N)` : 需要两个哈希表，每个哈希表大小和链表长度相同



## 链表——优化解法

其实我们完全不必要两个对应关系`Node (in primary list) -> index -> Node(in new copied list)`，可以直接建立`Node(old) -> Node(new)`的映射。

**继续优化**：可以不需要哈希表来建立映射，而是在拷贝后，直接将拷贝后的节点直接放在原链表原节点后面，只需最后解开两个链表即可。

### 代码二

```java
/*
// Definition for a Node.
class Node {
    int val;
    Node next;
    Node random;

    public Node(int val) {
        this.val = val;
        this.next = null;
        this.random = null;
    }
}
*/
class Solution {
    public Node copyRandomList(Node head) {
        if (head == null)
            return null;

        Map<Node, Node> map = new HashMap<>();
        Node dumHead = new Node(0);
        Node newCurr, oldCurr = head, pre = dumHead;
        while (oldCurr != null) {
            newCurr = new Node(oldCurr.val);
            pre.next = newCurr;

            map.put(oldCurr, newCurr);
            pre = newCurr;
            oldCurr = oldCurr.next;
        }
        map.put(null, null);

        oldCurr = head;
        newCurr = dumHead.next;
        while (newCurr != null) {
            newCurr.random = map.get(oldCurr.random);
            newCurr = newCurr.next;
            oldCurr = oldCurr.next;
        }
        return dumHead.next;
    }
}
```

### 代码三

```java
/*
// Definition for a Node.
class Node {
    int val;
    Node next;
    Node random;

    public Node(int val) {
        this.val = val;
        this.next = null;
        this.random = null;
    }
}
*/
class Solution {
    public Node copyRandomList(Node head) {
        if (head == null)
            return null;

        Node newCurr, oldCurr = head;
        while (oldCurr != null) {
            newCurr = new Node(oldCurr.val);
            newCurr.next = oldCurr.next;
            oldCurr.next = newCurr;

            oldCurr = newCurr.next;
        }

        Node newHead = head.next;
        oldCurr = head;
        while (oldCurr != null) {
            newCurr = oldCurr.next;
            newCurr.random = oldCurr.random == null ? null : oldCurr.random.next;

            oldCurr = newCurr.next;
        }
        
        oldCurr = head;
        newCurr = newHead;
        while (oldCurr != null) {
            oldCurr.next = newCurr.next;
            oldCurr = oldCurr.next;
            if (oldCurr == null) 
                break;
            newCurr.next = oldCurr.next;
            newCurr = newCurr.next;  
        }
        return newHead;
    }
}
```

## 图——DFS

参考自LeetCode，由于`random`域的存在，可以将复杂链表看成图，null也可以看成一个图节点。对于深拷贝问题，其实就是图的遍历问题，可以利用深度优先遍历 DFS和广度优先搜索 BFS。

**算法：深度优先搜索**

1. 从头结点`head`开始拷贝
2. 由于一个节点可能被多个指针指向，因此如果该节点已被拷贝，则不需要重复拷贝
3. 如果还没拷贝该节点，则创建一个新的节点进行拷贝，并将拷贝过的节点存到哈希表中
4. 使用递归拷贝所有的`next`节点，再递归拷贝所有的`random`节点

![2.png](https://pic.leetcode-cn.com/166afb3c11f82e09fdf3dd5e01731f12d73ae21c328b5981957a86b109e52c14-2.png)



```java
/*
// Definition for a Node.
class Node {
    int val;
    Node next;
    Node random;

    public Node(int val) {
        this.val = val;
        this.next = null;
        this.random = null;
    }
}
*/
class Solution {
    private HashMap<Node, Node> visited;
    public Node copyRandomList(Node head) {
        visited = new HashMap<Node, Node>();
        return copyNode(head);
    }

    public Node copyNode(Node node) {
        if (node == null) return null;
        if (visited.containsKey(node)) return visited.get(node);

        Node copy = new Node(node.val);
        visited.put(node, copy);
        copy.next = copyNode(node.next);
        copy.random = copyNode(node.random);
        return copy;
    }
}
```



## 图——BFS

广度优先搜索：维护一个队列，队列中保存的是当前广度搜索得到的节点，但还未访问，每次从队列中取元素，访问，然后添加它的相邻元素到队列中。

```java
/*
// Definition for a Node.
class Node {
    int val;
    Node next;
    Node random;

    public Node(int val) {
        this.val = val;
        this.next = null;
        this.random = null;
    }
}
*/
class Solution {
    public Node copyRandomList(Node head) {
        if (head == null) return null;

        HashMap<Node,Node> visited = new HashMap<Node, Node>();
        LinkedList<Node> queue = new LinkedList<>();
        Node aHead = new Node(head.val);
        visited.put(head, aHead);
        // visited.put(null, null);
        queue.addLast(head);

        Node p, q;
        while (!queue.isEmpty()) {
            p = queue.removeFirst();
            if (p.next != null && !visited.containsKey(p.next)) { // 如果p.next == null 那么copy之后的q节点next域不用赋值，如果p.next已经访问过直接取值即可
                q = new Node(p.next.val);
                visited.put(p.next, q);
                queue.addLast(p.next);
            }
            if (p.random != null && !visited.containsKey(p.random)) {
                q = new Node(p.random.val);
                visited.put(p.random, q);
                queue.addLast(p.random);
            }

            q = visited.get(p);
            q.next = visited.get(p.next);
            q.random = visited.get(p.random);
        }
        return aHead;
    }
}
```





