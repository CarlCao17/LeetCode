# 剑指Offer 27.二叉树的镜像

### 方法一 递归法

要实现二叉树的镜像，也就是要交换二叉树的左右子树，一直交换到叶子节点为止。首先，我们可以使用递归：

- 终止条件： `root == null || (root.left == null && root.right == null)`，当然你也可以只写前面一部分，这样最多多递归一层，也问题不大，因为节点数<= 1000，二叉树最多10层
- 交换： 先暂存 `root.left`，再令`root.left`等于镜像翻转后的`root.right`即`root.left = mirrorTree(root.right)`，同样再令`root.right = mirrorTree(temp)`
  **为什么要暂存`root.left`呢？**
  因为当你令`root.left = mirrorTree(root.right)`时，`root.left`已经指向镜像翻转后的右子树了，如果再令`root.right = mirrorTree(root.left)`就会出问题。可以联想到交换两个变量，也需要借助第三个临时变量，这是一样的。

#### 代码

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
    public TreeNode mirrorTree(TreeNode root) {
        if (root == null || (root.left == null && root.right == null))
            return root;
        TreeNode temp = root.left;
        root.left = mirrorTree(root.right);
        root.right = mirrorTree(temp);
        return root;
    }
}
```

### 方法二 辅助栈

和方法一没什么区别，只是借助栈非递归遍历树，遍历树的同时交换左右子树

#### 代码

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
    public TreeNode mirrorTree(TreeNode root) {
        if (root == null)
            return root;
        LinkedList<TreeNode> queue = new LinkedList<>();
        queue.addLast(root);
        while (!queue.isEmpty()) {
            TreeNode p = queue.removeFirst();
            TreeNode temp = p.left; // 先交换再加入，使得访问顺序按照交换后的顺序，比较直观，其实先加入再交换也可以
            p.left = p.right;
            p.right = temp;
            if (p.left != null) // 注意一定要加入非空节点到队列中，否则取出时会有可能是null，然后取左右子树就会有NullPointerException
                queue.addLast(p.left);
            if (p.right != null)
                queue.addLast(p.right);
        }

        return root;
    }
}
```

**时间复杂度**: `O(n)`，n为树的节点数，因为要遍历树中的每一个节点

**空间复杂度**: `O(n)`，最坏情况下，二叉树退化成一条链表形状，且每个节点都有两个子节点，则最多队列中最多存放`(n+1)/2)`个节点