# 剑指Offer 34. 二叉树中和为某一值的全部路径



[0,null, 0, null, 0]

[0,0,0,0]

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
    private List<List<Integer>> res;
    public List<List<Integer>> pathSum(TreeNode root, int target) {
        res = new ArrayList<List<Integer>>();
        if (root == null)
            return res;
        dfs(new ArrayList<Integer>(), root, 0, target);
        return res;
    }

    public void dfs(ArrayList<Integer> currPath, TreeNode currNode, int pathSum, int target) {
        if (currNode == null) {
            if (pathSum == target) {
                res.add(new ArrayList<>(currPath));
            }
            return ;
        }
        
        currPath.add(currNode.val);
        if (currNode.left != null || (currNode.left == currNode.right)) 
            dfs(currPath, currNode.left, pathSum + currNode.val, target);

        if (currNode.right != null)
            dfs(currPath, currNode.right, pathSum + currNode.val, target);
        currPath.remove(currPath.size()-1);
    }
}
```

