# 剑指Offer 32 - III. 从上到下打印二叉树 III（层序遍历 BFS + flag变量+链表反转）

### 解题思路

层序遍历很简单，只需要维护一个队列记录下一层要访问的节点即可。对于需要区分当前层和下一层，常见做法有两种：
1)每一层结束的时候插入一个标志如`null`，则当访问到标志时就可以确定当前层已访问到该层最后了，同时可以继续插入一个标志（这种方法的问题在于对于最后一层的特殊处理，可能造成死循环——原本有一个标志，然后检测到一个标志又插入一个...）；
2）利用一个新的辅助队列，保存下一层的节点。一次性遍历完当前层的所有节点，将下一层的节点记录到新的辅助队列中，然后将记录下一层节点的新的辅助队列赋值给队列

```java
queue.addLast(root);
while (!queue.isEmpty()) {
    LinkedList<TreeNode> nextLevel = new LinkedList<>();
    for (TreeNode p : queue) {
        // 遍历 p
        if （p.left != null)
            nextLevel.addLast(p.left);
        if (p.right != null)
            nextLevel.addLast(p.right);
    }
    queue = nextLevel;
}
```

这里提供一个新思路，只需要一个辅助队列即可，而不必反复创建新的队列。在遍历当前层的节点时仍然将下一层的节点添加到队列`queue`中，那么如何区分当前层和下一层节点呢，在`for`循环初始化时记录当前层的节点数，只需要循环这么多次即可取出全部节点。

```java
queue.addLast(root);
while (!queue.isEmpty()) {
    for (int i = queue.size(); i > 0; i--) {
        TreeNode p = queue.removeFirst();
        // 遍历 p 
        if (p.left != null)
            queue.addLast(p.left);
        if (p.right != null)
            queue.addLast(p.right);
    }    

}
```

那么如何实现每隔一层就反转记录的顺序呢？很容易想到用一个变量来控制顺序，每经过一层就改变值（比如用boolean变量）。那么如何控制逆序记录呢？
想到链表没？链表有尾插和头插，尾插得到的链表顺序就和插入顺序保持一致，头插则反序，具体到`Java API` 就是`LinkedList.addLast()`和`LinkedList.addFirst()`
最后说一句，这个变量可以用一个更有趣的做法：利用一个`int`变量，1 表示正序插入， 0 表示反序插入， 用异或操作控制翻转（正是因为`boolean`变量不好翻转，只能用条件判断加赋值，`urgly~`）.

```java
int flag = 1;

// 当需要翻转时
flag ^= 1;
```

### 代码

```java
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
class Solution {
    public List<List<Integer>> levelOrder(TreeNode root) {
        List<List<Integer>> res = new ArrayList<>();
        if (root == null)
            return res;
        
        LinkedList<TreeNode> queue = new LinkedList<>();
        queue.addLast(root);
        int flag = 1; // 正向打印
        while (!queue.isEmpty()) {
            LinkedList<Integer> level = new LinkedList<>();
            for (int i = queue.size(); i > 0; i--) {
                TreeNode p = queue.removeFirst();
                if (flag == 1) 
                    level.addLast(p.val);
                else 
                    level.addFirst(p.val);
                if (p.left != null) 
                    queue.addLast(p.left);
                if (p.right != null)
                    queue.addLast(p.right);
            }
            res.add(level);
            flag ^= 1; // 反转打印
        }
        return res;
    }
}
```

时间复杂度`O(N)`： N为树中的节点数，层序遍历需要访问所有的节点
空间复杂度`O(N)`： 最坏情况下（满二叉树）队列中需要存放 `N/2`的节点，但是常数比之前的双队列小一些
从II的实际运行结果也可以看到确实比双队列要快的多，双队列 2ms 只能打败 12%，而该方法 1ms 打败了99.85%用户

