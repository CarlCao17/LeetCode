# 剑指Offer 17. 打印从1到最大的n位数

这道题很简单，只要循环就可以解决，但是它最主要的考察点是大数，当然返回值是`int[]`，也就默认了不会超出`int`可表示的范围。

### 不考虑大数版本

```java
class Solution {
    public int[] printNumbers(int n) {
        int num = (int) Math.pow(10, n);
        int[] res = new int[num-1];
        for (int i = 1; i < num; i++)
            res[i-1] = i;
        return res;
    }
}
```



### 考虑大数

如果是大数的话，那就只能用字符串`String`来表示。`int`型可以不断+1，但`String`如何加一呢？可以这样来看待列举的问题，也就是将每一个数都看成`n`位，比如`1 = [0 0 0 0 0... 0 1]`，那么这就是一个全排列问题了，我们可以用DFS来做，而且DFS正好符合我们列举的过程，想一想我们列举的时候是不是画一棵树，然后先固定最高位，列举第二位，再固定第二位，列举第三位，...， 直到最后一位，才可以确定一个数。

**但是，还存在几个问题**：

1. 题目要求的是`int[] `，我们得到的是一个个字符串，如果是大数问题，应该只会要我们返回一个大的`String`，因为`int[]`表示不了，也装不下，那么我们只要使用`StringBuilder `一直保存结果即可，可以每一个元素都加一个`','`，然后将最后一个删掉。但是这个题目我们可以利用`Integer.parseInt()`将字符串转换成`int`；

2. 1有前缀`0000……1`怎么去除呢？可以像我下面写的一样，当要保存的时候就判断是否有前缀0，然后不保存那些前缀0即可，利用`String.valueOf(char[], int offset, int length)`，当然这个方法还不是很好，每次都要循环n次；

   **可以有更简单的办法**

   我们注意，`[1, 9]`的时候前缀有n-1个零，`[10-99]`的时候有n-2个零，`[100-999]`的时候有n-3个零，也就是说每当到`99`这种全是9的数的时候前缀就少一个零。那我们如何判断这种数呢？可以保存两个变量，一个`start`为`char[] num`中第一个不为零的位，另一个`nine`表示当前9的个数， 如果`n-start == nine`即非零的位全是9，也就是我们要的全是9的数，那么我们就把start--。

### 大数版本代码 一

```java
class Solution {
    private int n;
    private char[] num, loop = {'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'};
    private int[] res;
    private int index; // 指示结果已经保存到第index-1个，下一个保存在res[index]处

    public int[] printNumbers(int n) {
        // 大数打印
        this.n = n;
        num = new char[n];
        res = new int[((int) Math.pow(10, n)) - 1];
        dfs(0);
        return res;
    }

    private void dfs(int x) {
        if (x == n) {
            int i = 0;
            while (i < n && num[i] == '0')
                i++;
            if (i == n)
                return;
            res[index++] = Integer.parseInt(String.valueOf(num, i, n-i));
            return ;
        }
        for (char c : loop) {
            num[x] = c;
            dfs(x+1);
        }
    }
}
```



### 大数版本代码二 采用上面更简单的方法二 解决，同时输出String（只针对大数问题）

```java
class Solution {
    private int n;
    private char[] num, loop = {'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'};
    private int start, nine;
    private StringBuilder res;

    public String printNumbers(int n) {
        this.n = n;
        num = new char[n];
        start = n-1;
        res = new StringBuilder();
        dfs(0);
        res.deleteCharAt(res.length() - 1);
        return res.toString();
    }

    private void dfs(int x) {
        if (x == n) {
            String s = String.valueOf(num).substring(start);
            if (!"0".equals(s)) res.append(s + ",");
            if (n - start == nine) start--;
        }
        for (char c : loop) {
            if (c == '9') nine++;
            num[x] = c;
            dfs(x+1);
        }
        nine--; // 回溯的时候要减少9的个数
    }
}
```

