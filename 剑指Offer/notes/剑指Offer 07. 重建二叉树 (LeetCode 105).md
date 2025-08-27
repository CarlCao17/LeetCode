# 剑指offer 07. 重建二叉树 (LeetCode 105)

给定某二叉树的前序和中序遍历结果，要求重构二叉树

### 思路一 递归

优化的点：用一个hashmap保存反向 中序值-中序索引的映射，使得查找快速

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
    public TreeNode buildTree(int[] preorder, int[] inorder) {
        if (preorder == null || preorder.length == 0)
            return null;
        Map<Integer, Integer> inorderIdx = new HashMap<>();
        for (int i = 0; i < inorder.length; i++)
            inorderIdx.put(inorder[i], i);
        TreeNode root = buildTree(preorder, 0, preorder.length-1, inorder, 0, preorder.length-1, inorderIdx);
        return root;
    }

    public TreeNode buildTree(int[] preorder, int preStart, int preEnd, int[] inorder, int inStart, int inEnd, Map<Integer, Integer> inorderIdx) {
        if (preStart > preEnd)
            return null;
        int rootVal = preorder[preStart];
        int rootInorderIdx = inorderIdx.get(rootVal);
        TreeNode root = new TreeNode(rootVal);
        if (preStart < preEnd) {
            int leftNodes = rootInorderIdx - inStart, rightNodes = inEnd - rootInorderIdx;
            TreeNode leftSubtree = buildTree(preorder, preStart+1, preStart+leftNodes, inorder, inStart, rootInorderIdx-1, inorderIdx);
            TreeNode rightSubtree = buildTree(preorder, preStart+leftNodes+1, preEnd, inorder, rootInorderIdx+1, inEnd, inorderIdx);
            root.left = leftSubtree;
            root.right = rightSubtree;
        }
        return root;
    }
}
```

### 思路二 迭代

因为前序遍历的结果是从根节点到最左节点，然后到右节点；而中序遍历的结果是从最左节点依次往上，遇到有右节点则转向。

主要思路是：沿着前序遍历往下，一直到最左边节点（利用一个指向当前中序遍历的指针和当前for循环遍历的上一个节点值不相等），直到最左之后，遍历再倒序返回（此时和中序的结果是一样的，让中序的指针跟着走），一直到找到根节点或者找到有右节点为止，此时新建右节点，放入栈中（放入栈中即继续访问这颗子树）。

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
    public TreeNode buildTree(int[] preorder, int[] inorder) {
        if (preorder == null || preorder.length == 0)
            return null;
        
        int len = preorder.length;
        LinkedList<TreeNode> stack = new LinkedList<>();
        TreeNode root = new TreeNode(preorder[0]);
        stack.push(root);
        
        int inNodeIdx = 0;
        for (int i = 1; i < len; i++) {
            int preNodeVal = preorder[i];
            TreeNode node = stack.peek();
            if (node.val != inorder[inNodeIdx]) {
                node.left = new TreeNode(preNodeVal);
                stack.push(node.left);
            }
            else {
                while (!stack.isEmpty() && stack.peek().val == inorder[inNodeIdx]) {
                    node = stack.pop();
                    inNodeIdx++;
                }
                node.right = new TreeNode(preNodeVal);
                stack.push(node.right);
            }
        }
        return root;
    }
}
```

