# 剑指Offer 64. 求1+2+3+...+n

> 求 `1+2+...+n` ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。

目前能使用的：赋值、加减、移位和逻辑运算。

由于不能使用循环和条件判断，那么像循环累加 不可以；递归呢？

```java
public int recur(int num) {
    if (num == 0)
      	return 0;
    return recur(num-1) + num;
}
```

可以看到，递归需要有终止条件，就需要if，但是有没有办法可以绕开if呢？

有的，逻辑运算的短路性质就可以做到（其实在底层和if一样都是利用了`cmp`和条件跳转`jmp`指令来做到的，在C及x86体系中）。例如 `A || B `，只要A可以判断为`true`，那么B就不用再进行计算（逻辑与 `&&`类似），所以可以将递归终止放在A，将递归计算放在B中（还好Java也有短路效果）。

因此第一种解法就有了：

## 递归解法

```java
class Solution {
    public int sumNums(int n) {
        return recur(n);
    }
  
    public int recur(int num) {
        boolean flag = (num == 0) || (num += recur(num-1) > 0);
        return num;
    }
}
```

## 非递归解法

如果不使用递归呢？那么就应当将累加转换为$1+2+3+...+n = n*(n+1)/2$

那么如何表达乘法呢？容易想到快速乘法。

```java
public static int multi(int a, int b) {
    int res = 0;
    for (; b > 0; b >>>= 1, a <<= 1) { 
        res += a;
    }
    return res;
}
```

但是还有一个问题，需要用循环来不断判断乘数`b`最终是否为零，怎么替换循环呢？注意**`1 <= n <= 10000`**, 所以`n`和`n+1`最多只用到14位二进制，其他高位全是零，那我就手动加14次模拟循环足以，还有如何判断 `b>0`和前面一样用逻辑运算的短路效果来做，最终结果如下：

```java
class Solution {
    public int sumNums(int n) {
        int a = n+1;
        int b = n;
        int res = 0;
        boolean flag = (b & 1) > 0 && (res += a) > 0;
        a <<= 1;
        b >>= 1;
        flag = (b & 1) > 0 && (res += a) > 0;
        a <<= 1;
        b >>= 1;
        flag = (b & 1) > 0 && (res += a) > 0;
        a <<= 1;
        b >>= 1;
        flag = (b & 1) > 0 && (res += a) > 0;
        a <<= 1;
        b >>= 1;
        flag = (b & 1) > 0 && (res += a) > 0;
        a <<= 1;
        b >>= 1;
        flag = (b & 1) > 0 && (res += a) > 0;
        a <<= 1;
        b >>= 1;
        flag = (b & 1) > 0 && (res += a) > 0;
        a <<= 1;
        b >>= 1;
        flag = (b & 1) > 0 && (res += a) > 0;
        a <<= 1;
        b >>= 1;
        flag = (b & 1) > 0 && (res += a) > 0;
        a <<= 1;
        b >>= 1;
        flag = (b & 1) > 0 && (res += a) > 0;
        a <<= 1;
        b >>= 1;
        flag = (b & 1) > 0 && (res += a) > 0;
        a <<= 1;
        b >>= 1;
        flag = (b & 1) > 0 && (res += a) > 0;
        a <<= 1;
        b >>= 1;
        flag = (b & 1) > 0 && (res += a) > 0;
        a <<= 1;
        b >>= 1;
        flag = (b & 1) > 0 && (res += a) > 0;
        a <<= 1;
        b >>= 1;
        return res >>> 1;
    }
}
```



### 拓展

那么顺带复习一下快速幂乘

```java
public static long qPow(int a, int n) {
    // 计算a^n
    long res = 1;
    while (n > 0) {
        if (n & 1)
          res *= a;
        a *= a;
        n >>>= 1;
    }
    return res;
}
```

