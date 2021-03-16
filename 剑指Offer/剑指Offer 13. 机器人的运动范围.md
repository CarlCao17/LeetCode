# 剑指Offer 13. 机器人的运动范围

### 解题思路

这道题很容易想到在方格内进行搜索，可以利用DFS，和12.矩阵中的路径类似，只是具体处理有些许区别。

**首先**，DFS或者回溯的函数需要记录当前选择，以及历史路径或者说已经访问过的格子，类似于上一题的处理，可以使用一个矩阵`visited[m][n]`，初始全为0，如果要访问时先计算它的坐标数位之和，如果已经访问则标记为-1，利用几个不同的值来标记不同的状态。

**然后**，数位之和怎么计算呢？最直观的想法：

```
public int calSum(int num) {
    int res = 0;
    while (sum > 0) {
        res += (sum % 10);
        sum /= 10;
    }
    return res;
}
```

但是你想，矩阵的坐标是对称的，比如`[1, 2]`和`[2, 1]`难道我们要算两遍嘛？而且`calSum`的复杂度很高，为`O(k)`，`k`为`num`的位数。
因此我们可以借鉴DP或者备忘录的思想，因为矩阵的坐标`1 <= n, m <= 100`，所以我们可以提前计算并保存起来；
另外，如何解决`calSum`复杂度很高的问题呢？可以找找规律，因为范围不大，很容易发现`0-9`的数位和分别是0-9, `10-19`的数位和分别是1-10， `20-29`数位之和为2-11, ... 直到`90-99`的数位之和为9-18，100单算一下即可（其实不需要100也可以，因为下标最多到99)

**最后**，说明一下搜索过程：从起点`[0, 0]`出发，首先判断该节点是否能访问, 加入剪枝条件：坐标是否越界，该格子是否已经访问过，如果未访问是否不满足条件（行、列坐标数位之和 大于k）；然后标记该格子，使得之后不必重复计算；再去访问该格子的上下左右格子。

**等一下**,怎么没有终止条件，那不是会导致死循环呢？答案是不会，因为从格子`[0, 0]`开始，一直递归，然后不断深度优先搜索，直到所有搜索完返回到`[0,0]`那一层，最后所有递归栈全部退出，返回`movingCount`函数。

那还有一个问题，这样是否可以搜索到所有的格子呢，是否会漏？不会，可以简单理解，因为每一个格子都与其上一行和左一列相邻，利用递推易知，从`[0,0]`开始DFS，必然能搜索到所有的格子，只是我们利用剪枝省去搜索那些到达不了或者不满足条件的格子，但**本质上DFS或者回溯都是穷举**。
简单画个示意图：
![JianZhiOffer13_1.png](https://pic.leetcode-cn.com/1615866088-tUEthD-JianZhiOffer13_1.png)

### 代码

```
class Solution {
    private int[] sum;
    private int cnt;
    public int movingCount(int m, int n, int k) {
        int[][] visited = new int[m][n];

        sum = new int[101];
        getSum(sum);
        cnt = 0;

        backtrack(m, n, k, 0, 0, visited);
        return cnt;
    }

    public void backtrack(int m, int n, int k, int row, int col, int[][] visited) {
        // 剪枝，如果超出坐标范围，或者已经访问过
        if ((row < 0 || row >= m || col < 0 || col >= n) || visited[row][col] < 0)
            return;
        if (visited[row][col] == 0) // 之前没有计算过该格子的坐标数位之和
            visited[row][col] = sum[row] + sum[col]; 
        if (visited[row][col] > k) // 或者坐标数位之和大于k
            return;

        cnt++;
        visited[row][col] = -1; // 标记已经访问过

        backtrack(m, n, k, row, col-1, visited);
        backtrack(m, n, k, row, col+1, visited);
        backtrack(m, n, k, row-1, col, visited);
        backtrack(m, n, k, row+1, col, visited); 
    }

    public void getSum(int[] sum) {
        for (int i = 0; i < 10; i++)
            sum[i] = i;
        for (int i = 1; i < 10; i++) {
            for (int j = 0; j < 10; j++) {
                int num = i * 10 + j;
                sum[num] = sum[j] + i;
            }
        }
        sum[100] = 1;
    }
}
```

### 复杂度分析

时间复杂度 `O(mn)` 最坏情况下每个格子都要重复访问四遍
空间复杂度 `O(mn)`，构造的`visit`数组需要`O(mn)`，`sum`数组需要`O(1)`

## 另外一个剪枝版本

主要思路是在进入递归前先做判断，以减少不必要的递归深度

```
class Solution {
    private int[] sum;
    private int cnt;
    private int M, N, K;
    private int[][] visited;

    public int movingCount(int m, int n, int k) {
        visited = new int[m][n];
        sum = new int[101];
        
        cnt = 0;
        M = m;
        N = n;
        K = k;
        getSum(sum);

        backtrack(0, 0);
        return cnt;
    }

    public void backtrack(int row, int col) {
        
        cnt++;
        visited[row][col] = -1; // 标记已经访问过

        // 剪枝，如果超出坐标范围，或者已经访问过,或者坐标数位之和不满足条件
        if (couldVisit(row, col-1))
            backtrack(row, col-1);
        if (couldVisit(row, col+1))
            backtrack(row, col+1);
        if (couldVisit(row-1, col))
            backtrack(row-1, col);
        if (couldVisit(row+1, col))
            backtrack(row+1, col); 
    }

    /* 是否可以访问 */
    public boolean couldVisit(int row, int col) {
        if ((row < 0 || row >= M || col < 0 || col >= N) || visited[row][col] < 0)
            return false;
        if (visited[row][col] == 0) // 之前没有计算过该格子的坐标数位之和
            visited[row][col] = sum[row] + sum[col]; 
        if (visited[row][col] > K) // 或者坐标数位之和大于k
            return false;
        return true;
    }

    public void getSum(int[] sum) {
        for (int i = 0; i < 10; i++)
            sum[i] = i;
        for (int i = 1; i < 10; i++) {
            for (int j = 0; j < 10; j++) {
                int num = i * 10 + j;
                sum[num] = sum[j] + i;
            }
        }
        sum[100] = 1;
    }
}
```

#### 反思

但是好像差别不大，都是1ms，用时85.24%，反而内存消耗更大了，时间消耗大概是因为每次都要调用一个`couldVisit`函数做判断，所以其实与之前差别不大，但是空间消耗更大，可能是因为我把很多变量提升为域变量，而不是局部变量，我希望能减少栈的消耗，但可能反过来却加大了内存的消耗，以后需要注意