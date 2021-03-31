# 剑指Offer 55-II. 平衡二叉树

只要分别判断左子树和右子树是不是平衡二叉树，然后判断根节点是不是平衡的（左子树深度和右子树深度之差不超过1）即可。

这里利用DFS对二叉树做后序遍历，从底至顶返回子树的深度，如果该子树不平衡，则直接返回-1.

算法流程如下：

- **`dfs(root)`函数**

  1. 当`root`左右子树有一个不平衡，或者左右子树深度之差 >= 2，则返回-1，代表**这颗子树不是平衡树**（注意这里有做剪枝，只要一棵树的子树不是平衡树，那么这颗树必然不是平衡树，因此只要递归过程中有某一部分返回-1，则无须继续遍历直接返回-1，一直到顶层返回false)
  2. 只有当左右子树均平衡，且左右子树深度之差 <= 1，才返回这棵树的深度(因为递归返回到上一层时还需要利用子树的深度继续判断)，同时这里利用了如果返回的是深度(>=0)则代表了子树必然平衡——用一个深度代表了两种状态：子树深度值，以及子树平衡状态

  3. 终止条件：

         - 当`root`为空：说明越过叶子节点，因此返回高度0；
         - 当左右子树深度为-1：代表此树的左右子树不是平衡树，因此剪枝，直接返回-1

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
    public boolean isBalanced(TreeNode root) {
        if (root == null)
            return true;
        return dfs(root) < 0 ? false : true;
    }

    public int dfs(TreeNode root) {
        if (root == null) return 0;
        int depthLeft, depthRight;
        if ((depthLeft = dfs(root.left)) < 0 || (depthRight = dfs(root.right)) < 0 || Math.abs(depthLeft - depthRight) > 1)
            return -1;
        return Math.max(depthLeft, depthRight) + 1;
    }
}
```

时间复杂度 `O(N)`：最差情况下，需遍历树所有的节点

空间复杂度`O(N)`：最差情况下，树退化到链表，此时递归深度为`O(N)`