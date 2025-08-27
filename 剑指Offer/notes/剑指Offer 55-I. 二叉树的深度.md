# 剑指Offer 55-I. 二叉树的深度

### BFS 层序遍历

层序遍历只需使用队列即可，但是如何知道每一层呢？

两种方法：

一种当某一层结束的时候往队列中加入一个标志（如null），在开始添加第一层时就增加null。因此只要从队列中拿出null就知道这一层结束了，因此判断是不是最后一层，如果不是则继续往队列中加一个null，否则不要加了，不然会造成死循环，不管是不是最后一层都到了这一层的结尾，深度加一。

第二种，用两个队列，一个queue保存当前层的元素，一个tmp保存下一层的元素，当访问完本层的元素，下一层的元素也全部添加到tmp中，此时只需让queue = tmp，即可访问下一层，然后让深度加一。



### 代码

**哨兵方法：**

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
    public int maxDepth(TreeNode root) {
        int depth = 0;
        if (root == null)
            return depth;
            
        LinkedList<TreeNode> queue = new LinkedList<TreeNode>();
        queue.addLast(root);
        queue.addLast(null);
        while (!queue.isEmpty()) {
            root = queue.removeFirst();
            if (root != null) {
                if (root.left != null)
                    queue.addLast(root.left);
                if (root.right != null)
                    queue.addLast(root.right);
            }
            else {
                if (!queue.isEmpty())
                    queue.addLast(null); // 到达这一层最后一个元素
                depth += 1;
            }
        }
        return depth;
    }
}
```

**双队列法**

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
    public int maxDepth(TreeNode root) {
        int depth = 0;
        if (root == null)
            return depth;

        LinkedList<TreeNode> queue = new LinkedList<TreeNode>(), tmp;
        queue.addLast(root);
        while (!queue.isEmpty()) {
            tmp = new LinkedList<>();
            for (TreeNode node : queue) {
                if (node.left != null) tmp.addLast(node.left);
                if (node.right != null) tmp.addLast(node.right);
            }
            queue = tmp;
            depth += 1;
        }
        return depth;
    }
}
```

时间复杂度`O(N)`： 因为要遍历整棵树

空间复杂度`O(N)`：最坏情况下，当树为满二叉树时，队列`queue`同时存储 N/2个节点

### DFS 后序递归遍历

以当前节点为根节点的子树的深度等于左右子树深度最大值+1， 如果当前节点为空则返回0

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
    public int maxDepth(TreeNode root) {
        if (root == null)  return 0;
        return Math.max(maxDepth(root.left), maxDepth(root.right)) + 1;
    }
}
```

时间复杂度`O(N)` ：需要遍历整棵树

空间复杂度`O(N)`：最坏情况，树退化成链表，递归深度为N