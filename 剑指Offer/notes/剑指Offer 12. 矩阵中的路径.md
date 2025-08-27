# 剑指Offer 12. 矩阵中的路径

请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。

[["a","b","c","e"],
["s","f","c","s"],
["a","d","e","e"]]

但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。





## 分析

1. 用什么方法？

   这里我第一时间想到了回溯，因为搜到某个格子如果不符合条件则必然要回到上一个格子。

   那么这里的选择是什么呢？就是下一个格子的选择——上下左右

   路径列表就是之前已经遍历过的格子，如何保存路径列表使得能给定一个选择`<i, j>`就能判断这个格子是否遍历过？

   我直接使用了一个和board相同的boolean数组`visit`来保存之前是否访问过。

   还有没有其他方法呢？用Set或者Map，可以很容易某个元素确定在不在集合中。但是这里需要把`<i, j>`作为一个整体，如果用Set则需要新建一个类来表示整体，如果用Map,则i做键，j做值也可以。

   

   **最优想法**，直接用board保存，如果遍历到这个格子，就将这个格子中的值修改为一个不可能取到的值比如`'\0'`，如果需要回溯则改回来。

2. 如何做？

   首先用一个双重for循环遍历矩阵以确定初始选择，然后回溯或者也可以称为DFS
   
   注意这里可以做一些剪枝，比如选择下标越界，当前字符不符或是已经访问过。
   
   然后先深度搜索，搜索所有选择，之后需要**回复状态**。

## DFS + 剪枝

### 代码——额外使用boolean数组

```java
class Solution {
    public boolean exist(char[][] board, String word) {
        if (word == null || word.length() == 0)
            return false;
        
        char[] words = word.toCharArray();
        int m = board.length;
        int n = board[0].length;
        
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (backtrack(board, words, 0, i, j, visited))
                    return true;
            }
        }
        return false;
    }

    public boolean backtrack(char[][] board, char[] words, int currChar, int i, int j, boolean[][] visited) {
        if (i < 0 || i >= board.length || j < 0 || j >= board[0].length || board[i][j] != words[currChar] || visited[i][j])
            return false;
        if (currChar == words.length - 1)
            return true;
        
        visited[i][j] = true;
        boolean res = backtrack(board, words, currChar+1, i, j-1, visited) || backtrack(board, words, currChar+1, i-1, j, visited) ||   
                      backtrack(board, words, currChar+1, i, j+1, visited) || backtrack(board, words, currChar+1, i+1, j, visited);
        visited[i][j] = false;
        return res;
    }
}
```

### DFS 不适用额外空间保存路径

注意此时剪枝条件中的当前字符不符或者已经访问过可以合并，判断

```java
class Solution {
    public boolean exist(char[][] board, String word) {
        if (word == null || word.length() == 0)
            return false;
        
        char[] words = word.toCharArray();
        int m = board.length;
        int n = board[0].length;
        
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (backtrack(board, words, 0, i, j))
                    return true;
            }
        }
        return false;
    }

    public boolean backtrack(char[][] board, char[] words, int currChar, int i, int j) {
        if (i < 0 || i >= board.length || j < 0 || j >= board[0].length || board[i][j] != words[currChar])
            return false;
        if (currChar == words.length - 1)
            return true;
        
        board[i][j] = '\0';
        boolean res = backtrack(board, words, currChar+1, i, j-1) || backtrack(board, words, currChar+1, i-1, j) ||   
                      backtrack(board, words, currChar+1, i, j+1) || backtrack(board, words, currChar+1, i+1, j);
        board[i][j] = words[currChar];
        return res;
    }
}
```



### 复杂度分析

假设矩阵为`m * n`，字符串长`k`

两种解法时间复杂度没什么区别，最坏情况下，需要遍历矩阵中长度为`k`的字符串的所有方案，时间复杂度为`O(3^k)`；矩阵中共有`mn`个起点，时间复杂度为`O(mn)`

- 遍历矩阵中长度为`k`的字符串的所有方案，在搜索过程中，该字符有上下左右四个方向，除去回头（上个字符）的方向，剩下3种选择，因此寻找方案的复杂度为`O(3^k)`

空间复杂度，第一个方案为`O(mn+k)`，搜索过程中递归深度不会超过`k`，因此系统因函数调用累计使用的栈空间占用`O(k)`，额外的boolean数组占用`O(mn)`空间。第二个方案为`O(k)`



### 思考

回溯法最重要的就是确定每一次的选择和记录之前的状态——路径，以及结束条件。和动规不同，回溯不存在重叠子问题，本质上就是穷举，因此复杂度必然很高，只是有时可以做一些剪枝工作。

BFS和DFS也是穷举，只是做选择时的策略不同，它们也可以看成是回溯。



