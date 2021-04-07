# LeetCode 55.跳跃游戏

## 采用穷举——DFS

会超时的，或者栈溢出。

```java
class Solution {
    public boolean canJump(int[] nums) {
        return dfs(nums, 0);
    }

    public boolean dfs(int[] nums, int i) {
        if (i >= nums.length - 1)  // 如果某一次跳过了终点，则必然存在某个更小的步数恰好跳到终点，因为最多步数取1，此时必然上一次在终点前
            return true;
        
        boolean res = false;
        for (int j = nums[j]; j > 0; j--) { // 直接从最大步数开始跳，可能可以加快搜索速度
            res = res || dfs(nums, i+j);
            if (res)
                break;
        }
        return res;
    }
}
```

时间复杂度$O(n^m)$: $n$是数组的长度，$m$是数组元素的最大值，虽然已经做了足够的剪枝以及优化（如果已经找到立即返回，并且从最大开始搜，这样直觉上更容易找过去——当然并不一定，有可能步伐太大反而跨不过去）

## DP

