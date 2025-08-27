# 剑指Offer 11. 旋转数组的最小数字

## 暴力解法

由于递增数组旋转的特性，使得当前的数组必然呈现先上升后下降再上升，如图：

![](https://assets.leetcode-cn.com/solution-static/jianzhi_11/1.png)

则遍历数组，找到第一个比前一个元素小的元素必然是原始数组的第一个元素，即为最小数字，当然如果该数组未经过旋转（即为递增数组）则必然找不到，那么返回头一个元素即可。

```java
class Solution {
    public int minArray(int[] numbers) {
        if (numbers == null || numbers.length == 0)
            return -1; // Illegal Param

        int n = numbers.length;
        for (int i = 1; i < n; i++) {
            if (numbers[i] < numbers[i-1])
                return numbers[i];
        }
        return numbers[0];  // the array `numbers` is sorted in asccend order (including the case like `2 2 2`).
    }
}
```

### 复杂度分析

时间复杂度 `O(n)`

空间复杂度 `O(1)`

超过23%

## 二分查找

不妨假设旋转数组最后一个数字为`x`，而当前最小的数如图中所示为`min`，则在`min`左侧的元素都大于`x`，在`min`右侧的元素都小于`x`，利用这个性质来做二分。



给定初始区间为整个数组`low = 0, high = numbers.length - 1`，必然满足性质：最小值`min`在区间中；

区间中点`mid`的值

- 如果小于区间右端点`high`的值，那么区间中点必然在`min`的右侧，因此可以忽略当前区间右半部分；

- 如果大于区间右端点`high`的值，那么区间中点必然在`min`的左侧，因此可以忽略当前区间左半部分；
- 如果相等，那么从区间中点`mid`到区间右端点`high`都相等，当然此时最小值`min`可能在区间右半部分`[mid, high]`的左边，也可能在右边，极端情况如下图。则此时无法预测，但是我们知道`high`可以用它前面一个元素代替，缩小区间为`[low, high-1]` （如果是下图中`pivot`在右边，那么`[mid, high]`之间都可以，但是如果是`pivot`在左边，则只有`high`前面一部分元素才可，保守起见，只往前试探一个元素）

  ![](https://assets.leetcode-cn.com/solution-static/jianzhi_11/4.png)

### 代码

```java
class Solution {
    public int minArray(int[] numbers) {
        if (numbers == null || numbers.length == 0)
            return -1; // Illegal Param

        int n = numbers.length;
        int low = 0, high = n-1;
        while (low < high) {
            int mid = low + (high - low) / 2;
            if (numbers[mid] < numbers[high])
                high = mid;
            else if (numbers[mid] > numbers[high])
                low = mid + 1;
            else 
                high = high - 1;
        }

        return numbers[low];
    }
}
```

### 分析

循环终止条件`low >= high`, 但是由于更新条件，则必然区间缩小到`[low==high]`（因为`mid`可能位于区间中点——奇数，或者偏左边一个元素——偶数，则`high = mid`保证`high`不会越过`low`，而`low = mid + 1`最多使得`low`到达`high`)，**这也是与一般的二分不一样的地方**

当区间缩小到一个元素时，且区间一直保持性质：区间内包含最小值，则必然`min = numbers[low]`。

至于为什么区间一直保持该性质呢？

- 初始情况，`[low, high]`为整个区间，必然成立
- 假设之前的区间满足性质，则当缩小后的区间必然满足条件，可以参见上面的算法设计过程

又区间必然从初始区间，缩小到1个元素（不论上一次是2个，还是3个），因此成立。



时间复杂度分析：最坏情况下如`[2, 2, 2, 2]`，区间一个元素一个元素的缩小，`O(n)`；一般来说，利用二分可以达到`O(logn)`

空间复杂度分析：`O(1)`





### 补充

能否使用最左边的位置进行判断呢？

![image-20210315111026671](/Users/caozhengcheng/developEnv/GitHubRepo/LeetCode/pic/jianzhi-11_1.png)