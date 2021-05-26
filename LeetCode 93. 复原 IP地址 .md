# LeetCode 93. 复原 IP地址 DFS+剪枝 打败99.94%（时间复杂度为O(1))

### 解题思路

首先这个题目很容易想想到DFS搜索所有情况，然后保存有效的情况并且回溯。
但是注意因为最多有4个0-255的整数且不含前导0组成有效IP地址。
因此在每次搜索前先判断剩下的长度是不是过长或者过短，如果是的话说明当前的路径对应的划分是不合理的，则不必继续搜索下去了（对应代码中计算`remind`的部分)。为此，我们保存了一个变量`k`表示当前已经划分好了`k`个部分
最后注意回溯需要还原状态`sb`，以及保存时由于结尾多了一个'.'即可。

### 代码

```java
class Solution {
        private List<String> res;
        private String s;

        public List<String> restoreIpAddresses(String s) {
            res = new ArrayList<>();
            this.s = s;
            int n = s.length();

            StringBuilder sb = new StringBuilder();
            dfs(0, sb, n, 0);
            return res;
        }

        private void dfs(int i, StringBuilder sb, int n, int k) {
            if (k == 4 || i == n) {
                if (k == 4 && i == n) {
                    res.add(sb.substring(0, sb.length()-1).toString());
                }
                return;
            }

            for (int j = 0; j < 3; j++) {
                if (checkValid(i, i+j+1, n)) {
                    int remind = n - (i + j + 1);
                    if (remind > 3 * (3 - k) || remind < (3 - k))
                        continue;
                    sb.append(s.substring(i, i+j+1)).append('.');
                    dfs(i+j+1, sb, n, k+1);
                    int len = sb.length();
                    sb.delete(len-j-2, len);
                }
            }
        }

        private boolean checkValid(int i, int j, int n) {
            if (j > n || (s.charAt(i) == '0' && j > i + 1))
                return false;
            return Integer.parseInt(s.substring(i, j)) <= 255;
        }
    }
```

**时间复杂度** O(1)*O*(1): 因为最多递归到第4层，而每一层都最多只有3种情况，递归与字符串长度`n`无关
**空间复杂度** O(n)*O*(*n*): 虽然递归深度与`n`无关，但是我们额外需要一个StringBuilder保存划分的结果