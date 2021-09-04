# LeetCode 647. 回文子串

最朴素的想法当然是挨个遍历每一个起始位置，然后再往里收缩可能的回文字符串，如 `s = abab`，对于起始为第1个字符`b`的串，从`bab`开始搜，搜索就需要$O(n^2)$，而判断回文呢？也需要$O(n)$​，那这就需要立方级别的复杂度。

太高了~怎么改进呢？是否回文判断可以变成$O(1)$呢，当然可以对于一个字符串$S_{i_1}S_{i_2}...S_{i_m}$​，只要知道$S_{i_2}...S_{i_m-1}$​，那么很快就能判断是否回文，因此用DP是最容易想到的办法。

## 解法一 DP + $O(n^2)$​遍历

因此只需要按照长度和起始位置来搜索即可。

$$dp[i][j] = \begin{cases}
s[i] == s[j] \&\& dp[i+1][j-1] & len > 2\\
s[i] == s[j] & len = 2\\
\end{cases}$$​​​​

其中 $len = j - i + 1$，其中$dp[i][j]$表示的是字符串$S_{i}S_{i+1}...S_{j}$​是否是回文字符串

```java
class Solution {
    public int countSubstrings(String s) {
        if (s == null || s.length() == 0)
            return 0;
        int n = s.length();
        boolean[][] isPara = new boolean[n][n];
        for (int i = 0; i < n; i++) {
            isPara[i][i] = true;
        }
        int cnt = n;
        for (int len = 2; len <= n; len++) {
            for (int i = 0; i <= n - len; i++) {
                int j = i + len - 1;
                if (len == 2) 
                    isPara[i][j] = s.charAt(i) == s.charAt(j);
                else 
                    isPara[i][j] = s.charAt(i) == s.charAt(j) && isPara[i+1][j-1];
                cnt += (isPara[i][j] ? 1 : 0);
            }
        }
        return cnt;
    }
}
```

![image-20210724104751887](../../../Library/Application Support/typora-user-images/image-20210724104751887.png)

时间复杂度$O(n^2)$: 长度遍历$O(n)$， 挨个搜索$O(n)$

空间复杂度$O(n^2)$​​



## 解法二 中心扩展搜索 + Manacher 算法





