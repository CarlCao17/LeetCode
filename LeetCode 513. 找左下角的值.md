# LeetCode 513. 找左下角的值

## BFS 层序遍历

很容易想到，我每次遍历的时候保存上一层，当访问到最后一层的时候就可以最后一层啦。然后取这一层的第一个元素即可（利用队列）

```java
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    public int findBottomLeftValue(TreeNode root) {
        LinkedList<TreeNode> queue = new LinkedList<>();
        LinkedList<Integer> lastLevel = new LinkedList<Integer>();
        queue.addLast(root);

        while (!queue.isEmpty()) {
            lastLevel = new LinkedList<Integer>();
            for (int i = queue.size(); i > 0; i--) {
                TreeNode p = queue.removeFirst();
                lastLevel.addLast(p.val);
                if (p.left != null)
                    queue.addLast(p.left);
                if (p.right != null)
                    queue.addLast(p.right);
            }
        }
        return lastLevel.removeFirst();
    }
}
```

时间复杂度$O(n)$: 遍历树需要$O(n)$，超过69%

空间复杂度$O(n)$: 层序遍历需要一个辅助队列，我们另外用一个队列来保存上一层，在最坏情况下队列需要保存$n/2$个元素

## DFS 先序遍历

保存一个当前最大深度，和当前最大深度对应的值，在DFS遍历过程中，如果当前树的深度大于最大深度，说明正在访问的节点比之前的节点都深，因此更新最大深度和最大深度对应的值。由于树的先序遍历保证一定是从左到右的顺序，所以保存的最大深度的值一定是当前深度最左的元素的值。

```java
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    private int maxDepth;
    private int res;
    
    public int findBottomLeftValue(TreeNode root) {
        maxDepth = 0;
        res = -1;
        dfs(root, 1);
        return res;
    }

    public void dfs(TreeNode root, int currDepth) {
        if (root == null) 
            return;
        if (currDepth > maxDepth) {
            maxDepth = currDepth;
            res = root.val;
        }
        dfs(root.left, currDepth+1);
        dfs(root.right, currDepth+1);
    }
}
```

时间复杂度`O(n)`：需要遍历树中全部元素，超过100%

空间复杂度`O(n)`：只需要系统栈来额外保存状态，不需要其他的空间。

