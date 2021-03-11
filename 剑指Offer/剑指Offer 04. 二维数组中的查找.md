# 剑指Offer 04. 二维数组中的查找

在一个 `n * m` 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

## 思路一 暴力求解

直接遍历整个`n * m`的矩阵，然后判断是否存在

### 代码

```java
class Solution {
    public boolean findNumberIn2DArray(int[][] matrix, int target) {
        if (matrix == null || matrix.length == 0 || matrix[0].length == 0)
            return false;
            
        int n = matrix.length, m = matrix[0].length;
        
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                if (matrix[i][j] == target) 
                    return true;
            }
        }
        return false;
    }
}
```

### 分析

时间复杂度 `O(mn)`

空间复杂度`O(1)`



## 思路二 二分查找

对每一行都做二分查找，快速判断每一行是否存在`target`

### 代码

```java
class Solution {
    public boolean findNumberIn2DArray(int[][] matrix, int target) {
        if (matrix == null || matrix.length == 0 || matrix[0].length == 0)
            return false;
            
        int n = matrix.length, m = matrix[0].length;

        for (int i = 0; i < n; i++) {
            int low = 0, high = m-1;
            while (low <= high) {
                int mid = low + (high - low) / 2;
                if (matrix[i][mid] == target)
                    return true;
                else if (matrix[i][mid] > target) 
                    high = mid - 1;
                else 
                    low = mid + 1;
            }
        }   
        
        return false;
    }
}
```

### 分析

时间复杂度 `O(nlogm)`, 遍历每一行`O(n)`，对每一行做二分`O(logm)`

空间复杂度`O(1)`



### 错误思路辨析

先在第一行利用二分搜索，找到第一个比它大的元素，然后在它前一列里再使用二分，如果存在则找到了，反之如果不存在则返回`false`。

**上述分析**✅？

当然错误，如果找到返回`true`没问题，但是如果在这一行这一列里没找到能说明没有嘛？

答案是不能，来看反例：target  = 5

 1    2    3   7  9

 4    5    6   8 10

11 12 13 14 15

**如果我先利用二分，找到的必然是3这一列，而事实上，5在2这一列，必然找不到，但是不能说明没有**



## 思路三 利用矩阵的性质，利用二分思想

要利用二分，就必须一边小，一边大，然后才能排除一边。这里矩阵很明显，它的同一列的下面的元素必然比该元素大，它的同一行的左边的元素必然比该元素小（反之亦可）

所以从右上顶点元素（也可以左下）开始，类似二分

### 代码

```java
class Solution {
    public boolean findNumberIn2DArray(int[][] matrix, int target) {
        if (matrix == null || matrix.length == 0 || matrix[0].length == 0)
            return false;
            
        int n = matrix.length, m = matrix[0].length;

        int i = 0, j = m - 1;
        while (i < n && j >= 0) {
            if (matrix[i][j] == target)
                return true;
            else if (matrix[i][j] > target)
                j--;
            else 
                i++;
        } 
        
        return false;
    }
}
```

### 分析

时间复杂度： `O(n+m)`

空间复杂度： `O(1)`

