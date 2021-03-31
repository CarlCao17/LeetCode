# 剑指Offer 26.树的子结构

### 解题思路

如果B是A的子结构，则必然B的根节点是A中的某个子节点。
所以算法有两步（分别对应`isSubStructure`和`equals`）：

1. 遍历树A，在A中找到与B的根节点相同的值，如果没有找到则必然B不是A的子结构
2. 在A中找到与B的根节点相同的节点后，再依次遍历树B判断是否相同
   遍历两颗树都可以使用递归加非递归

- `isSubStructure(TreeNode A, TreeNode B)`是判断树A中是否存在和树B相同的结构：
  - A或者B为空，则返回false，因为空子树不是任何树的子结构，任何树都不是空树的子结构
  - 首先访问A的根节点（调用`equals`判断以A根节点为根的子树是否存在和B相同的结构和值），然后再递归访问A的左子树和右子树
- `equals(TreeNode A, TreeNode B)`是判断A子树和B子树是否相同：
  - 如果B为空，则说明比较已经访问了B这棵树的叶子节点，返回true
  - 如果A为空，又知道B不为空，则必然当前以节点A为根节点的子树与以节点B为根节点的子树不相同，返回false
  - 否则，说明A子树的根节点和B子树的根节点相同，继续递归遍历A子树的左子树和B子树的左子树

### 递归版本

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
    public boolean isSubStructure(TreeNode A, TreeNode B) {
        return (A != null && B != null) && (equals(A, B) || isSubStructure(A.left, B) || isSubStructure(A.right, B));
    }

    public boolean equals(TreeNode A, TreeNode B) {
        if (B == null)
            return true;
        if (A == null || A.val != B.val)
            return false;
        if (B.left == null && B.right == null) // 增加一个判断，以减少一层不必要的递归深度
            return true;
        return equals(A.left, B.left) && equals(A.right, B.right); 
    }
}
```

### 分析

时间复杂度 `O(MN)`, 其中M是树A的节点个数，N是树B的节点个数。因为`isSubStructure`需要遍历树A，而最坏情况每访问一个A的节点都需要遍历B中的每一个节点，
空间复杂度 `O(M)`, 当树A和树B都退化为链表时递归深度最大，若M <= N，最多递归到树A的叶子节点，因此遍历树A以及递归判断的总深度为`O(M)`；若M > N，递归树A及递归判断的总深度必然为`O(M)`

**时间消耗 超过100%**

### 非递归版本

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
    public boolean isSubStructure(TreeNode A, TreeNode B) {
        if (B == null || A == null)
            return false;
        LinkedList<TreeNode> stack = new LinkedList<>();

        while (A != null || !stack.isEmpty()) {
            while (A != null) {
                if (equals(A, B))
                    return true;
                stack.push(A);
                A = A.left;
            }
            A = stack.pop().right;
        }
        return false;
    }

    public boolean equals(TreeNode A, TreeNode B) {
        if (B == null)
            return true;
        if (A == null || A.val != B.val)
            return false;
        if (B.left == null && B.right == null)
            return true;
        return equals(A.left, B.left) && equals(A.right, B.right); 
    }
}
```

时间消耗超过4%