# 剑指Offer 54. 二叉搜索树的第k大节点

### DFS正常访问+维护一个k大的队列

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
    public int kthLargest(TreeNode root, int k) {
        ConstantCapacityQueue queue = new ConstantCapacityQueue(k);
        LinkedList<TreeNode> stack = new LinkedList<TreeNode>();
        
        while (root != null || !stack.isEmpty()) {
            while (root != null) {
                stack.addFirst(root);
                root = root.left;
            }
            root = stack.removeFirst();
            queue.add(root.val);
            root = root.right;
        }
        return queue.getFirst();
    }

    class ConstantCapacityQueue {
        private LinkedList<Integer> queue;
        private int capacity;
        /*
        如果固定容量队列未满，则直接添加元素；如果满了，则移除最久的元素，然后添加
        */
        public ConstantCapacityQueue(int capacity) { 
            queue = new LinkedList<Integer>();
            this.capacity = capacity;
        }

        public void add(int value) { // 由于二叉搜索树中序遍历一定是从小到大的顺序，因此每次新添加的元素一定比在固定容量队列中的元素都大，直接添加到最后，
                                     // 保证该队列是有序的，且是当前访问的所有节点中最大的k个元素
            if (queue.size() >= capacity) { 
                queue.removeFirst();
            }
            queue.addLast(value);
        }

        public Integer getFirst() { // 只从头部取元素
            return queue.getFirst();
        }
    }
}
```



### DFS 倒序访问

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
    int res;
    int cnt;

    public int kthLargest(TreeNode root, int k) {
        cnt = 0;
        dfs(root, k);
        return res;
    }

    public void dfs(TreeNode root, int k) {
        if (root == null)
            return;
        dfs(root.right, k); // 访问它的右子树
        cnt++; // 访问这个元素
        if (cnt > k)
            return;
        if (cnt == k) { // 如果已经累计访问到二叉搜索树中序倒序的第k个元素，直接返回
            res = root.val;
            return;
        }
        dfs(root.left, k); // 访问它的左子树
    }
}
```

