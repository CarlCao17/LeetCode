# 剑指Offer 57 II. 和为s的连续整数序列

 

## 暴力 + 穷举

```java
class Solution {
    public int[][] findContinuousSequence(int target) {
        int half = target / 2 ;
        List<int[]> res = new ArrayList<>();

        int j, sum;
        for (int i = 1; i <= half; i++) {
            j = i+1;
            sum = i;
            while (j < target && sum < target) {
                sum += j;
                j++;
            }
            if (sum == target) {
                int[] tmp = new int[j-i];
                for (int k = i; k < j; k++)
                    tmp[k-i] = k;
                res.add(tmp);
            }
        }
        int[][] res1 = res.toArray(new int[res.size()][]);
        return res1;
    }
}
```

时间复杂度 

空间复杂度

超过 79.83%





## 穷举+数学优化

对于连续数列`[x, y]`的和是否等于`target`，不必累加，可以直接利用等差数列的和`(x + y)(y - x + 1) / 2`是否等于`target`。所以对于以`i`开始的子序列中是否有某一个的和等于`target`，可以直接求解方程，如果存在某个正整数解，且大于`x`，则必定存在。

可写出如下代码：

```java
class Solution {
    public int[][] findContinuousSequence(int target) {
        int half = target / 2 ;
        double delta2;
        int delta, y;
        List<int[]> res = new ArrayList<>();

        for (int i = 1; i <= half; i++) {
            delta2 = 4 * i * i - 4 * i + 8 * target + 1;
            if (delta2 < 0)
                continue;
            delta = (int)Math.sqrt(delta2);
            if (Double.compare(delta2, delta * delta) != 0)
                continue;
            if ((delta & 1) == 0) 
                continue;
            y = (delta - 1) / 2;
            int[] tmp = new int[y-i+1];
            for (int k = i; k <= y; k++)
                tmp[k-i] = k;
            res.add(tmp);
        }
        int[][] res1 = res.toArray(new int[res.size()][]);
        return res1;
    }
}
```



### 小葵花课堂来了~

细心的同学有发现问题嘛？

运行之后发现不对劲：

<img src="../pic/image-20210408100611822.png" alt="image-20210408100611822" style="zoom:50%;" />

仔细想想，大数出了问题。。。

应该能想到吧，在计算`delta`时，$98160 ^ 2 > 2 * 10^9$ 溢出了，注意看因为我使用的是$4 * i * i$，而 $i$ 定义成 `int`，所以会出错，同样在判断`Double.compare(delta2, delta * delta) != 0)`这里 `delta`被定义成`int`，同样可能溢出。

这里有个问题，那需要改成`BigInteger`或者`BigDecimal`嘛？并不需要，

所以改为：

```java
class Solution {
    public int[][] findContinuousSequence(int target) {
        int half = target / 2 ;
        List<int[]> res = new ArrayList<>();

        for (int i = 1; i <= half; i++) {
            long delta = 4 * i * i - 4 * i + 8 * target + 1l; // 1 - 4 * (x - (long) x * x - 2 * target))
            if (delta < 0)
                continue;
            int delta_sqrt = (int)Math.sqrt(delta);
            if ((long) delta_sqrt * delta_sqrt == delta && (delta & 1) == 1) {
                int y = (delta_sqrt - 1) / 2;
                if (i < y) {
                    int[] tmp = new int[y-i+1];
                    for (int k = i; k <= y; k++)
                        tmp[k-i] = k;
                    res.add(tmp);
                }           
            }
        }
        int[][] res1 = res.toArray(new int[res.size()][]);
        return res1;
    }
}
```



### 双指针 

给定双指针 $[l, r]$，记区间$[l,r]$内所有值的和为$Sum[l,r]$，显然$Sum[l,r] = (l + r) (r - l +1) / 2$:

- 如果 $Sum[l,r] == target$，那么说明找到了一个序列，且知道以$l$开头的子序列必不存在第二个和等于$target$，因此可以继续搜索下一个以$l+1$开头的子序列。（此时，利用已知信息，$[l+1, r]$区间任意连续序列的和都小于$target$，因此不必将右指针挪回到左指针右边，可以保留在原地或者向右移一格；
- 如果 $Sum[l,r] > target$，那么右指针 $r$之后的序列也不必再搜索了，必定大于$target$，可以直接搜索下一个以 $l+1$开头的子序列，这时注意必定是$Sum[l, r-1] < target$ 且 $Sum[l, r] > target$，所以同样利用这个信息，$l$向右移，而$r$不需要回溯，因为$Sum[l+1, r-1] < Sum[l, r-1] < target$；
- 如果$Sum[l, r] < target$，则右指针 $r$右移一格

最后，何时终止呢？我们知道右指针必定不能跨过$half+1 = target / 2 + 1$ ，因为当它跨过时，其区间内至少有两个元素（$half + half+1 >= target$）——这个时候假设左指针还没走过来，所以当右指针走到这个位置时必定会停下来并且之后都不会再增加了，因为区间内元素已经比$target$大了，这个时候应该缩小区间长度，`l++​`，直到 $l$ 跨过 $r$ 或者到达 $r$ 所在的位置，迭代终止。

，则当左指针跨过$half$时，必定$Sum[l, r] > target$



### 代码

```java
class Solution {
    public int[][] findContinuousSequence(int target) {
        List<int[]> vec = new ArrayList<int[]>();
        for (int l = 1, r = 2; l < r;) {
            int sum = (l + r) * (r - l + 1) / 2;
            if (sum == target) {
                int[] res = new int[r - l + 1];
                for (int i = l; i <= r; ++i) {
                    res[i - l] = i;
                }
                vec.add(res);
                l++;
            } else if (sum < target) {
                r++;
            } else {
                l++;
            }
        }
        return vec.toArray(new int[vec.size()][]);
    }
}
```

