# LeetCode 50. Pow(x, n) 快速幂 + 特殊情况排除 超过98.87%

### 解题思路

利用快速幂，需要注意两点：

1. 如果指数为负数，可以先将指数转换为正数，也就是代码中的`if (n < 0)`
2. 对于特殊的 0^0 = 100=1 需要优先判断，而，然后对于底数是1或者0可以提前判断(0^k = 0, k != 00*k*=0,*k*!=0，或者1^k=11*k*=1)，这样可以提升case的运行

### 代码

```java
class Solution {
    public double myPow(double x, int n) {
        double res = 1;
        if (Double.compare(x, 1) == 0 || n == 0) // 为了加快计算速度，排除掉底数为1或者0，以及特殊的0^0=1
            return 1;
        if (Double.compare(x, 0) == 0)
            return 0;
        if (n < 0) {
            x = 1/x;
            n = -n;
        }
        while (n > 0) {
            if ((n & 1) > 0)
                res *= x;
            x *= x;
            n >>>= 1;
        }
        return res;
    }
}
```