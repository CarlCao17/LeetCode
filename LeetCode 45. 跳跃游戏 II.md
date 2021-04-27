# LeetCode 45. 跳跃游戏 II



## 暴力搜索——DFS

```java
class Solution {
    private int times = Integer.MAX_VALUE;

    public int jump(int[] nums) {
        int n = nums.length;
        dfs(nums, 0, 0);
        return times;
    }

    public void dfs(int[] nums, int curr, int path) {
        int n = nums.length;
        if (curr >= n-1) {
            times = Math.min(times, path);
            return;
        }

        for (int i = 1;  i <= nums[curr]; i++) {
            dfs(nums, curr + i, path+1);
        }
    }
}
```

时间复杂度太高了，TLE

可以做一个简单的优化，记忆从起始点到当前位置的最短路径长度

```java
class Solution {
    private int[] steps;
    public int jump(int[] nums) {
        int n = nums.length;
      	steps = new int[n];
      	Arrays.fill(steps, Integer.MAX_VALUE);
      	steps[0] = 0;
        dfs(nums, 0);
        return steps[n-1];
    }

    public void dfs(int[] nums, int curr) {
        int n = nums.length;
				int end = curr + nums[curr];
        for (int next = curr+1; next < n && next <= end; next++) {
            if (steps[next] <= steps[curr] + 1) continue; // 提前剪枝，如果下一步有比从当前调到下一步更短的路径，则无须更新也无须递归
          	steps[next] = steps[curr] + 1;
            dfs(nums, next);
        }
    }
}
```

用时1ms，打败了54.21%

## 优化一： DP

从上面已经看到了一点DP的影子，那我们为什么还需要递归呢？

如果我不用递归，而是每次从当前往下跳的时候选择的都是当前可以选择的最优的（即路径最短的），那我就不需要递归啦~

但是怎么保证这个最优呢？我们可以从最后一步开始反向遍历，类似于走台阶，第`n-1`步走到终点只需要0步，第`n-2`步要看它能否到达终点（即`nums[n-2] > 0`)，如果能就是1，否则就是不能到达。那第`n-3`步呢？

需要看第`n-3`步能否直接到达终点，如果能则它只需要1步，如果不能，则在它所有能跳到的路径中选择一条最短的。

这就是DP的思路：

- base case: 第`n-1`步只需要0步就可以到达终点

- 状态: 第`k`步到达终点需要的最短步数

- 转移: 也就是从第`k`步走到下一跳

- 状态转移方程: 

  $$dp[i] = \begin{cases}1 & i + nums[i] >= n-1,\\ min_{k=1,2,...,nums[i]}{dp[k]} +1 & else \end{cases}$$

```java

```



## 优化二： 贪心 I 反向



## 优化三： 贪心 II 正向

```java
class Solution {
    public int jump(int[] nums) {
        int n = nums.length;
        int end = 0;
        int maxPosition = 0;
        int steps = 0;
        for (int i = 0; i < n-1; i++) {
            maxPosition = Math.max(maxPosition, i + nums[i]);
            if (i == end) {
                end = maxPosition;
                steps++;
            }
        }
        return steps;
    }
}
```



