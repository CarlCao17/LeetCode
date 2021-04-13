# LeetCode 30. 串联所有单词的子串

详细分析可以看https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-w-6/

## 暴力搜索解法

本质是让`left`指针在区间`[0, s.length() - allLen+1]`进行搜索，每次搜索一个固定长度`allLen`的子串，如果这个子串符合就记录下来否则就让`left`从下一个位置开始重新搜索，其中`allLen`是words所有单词拼接出来的字符串长度。

```java 
class Solution {
    public List<Integer> findSubstring(String s, String[] words) {
        List<Integer> res = new ArrayList<>();
        if (s == null || s.length() == 0)
            return res;
        if (words == null || words.length == 0)
            return res;

        int match; // 当前搜索的子串中，不重复的单词与words匹配（即个数相同）的数量
        int left, right;

        int step = words[0].length(); // 单词长度，即每次跳跃的长度
        int num = words.length;
        int n = s.length();
        int end = n - step * num; // 最后搜索到的位置

        Map<String,Integer> needs = new HashMap<>();
        Map<String, Integer> curr = new HashMap<>(num);

        for (String w : words) needs.put(w, needs.getOrDefault(w, 0) + 1);

        for (left = 0; left <= end; left++) {
            curr.clear();
            match = 0;
            right = left;
            for (int i = 0; i < num; i++) { // 搜索[left, left + num * step]，每次查一个单词
                String word = s.substring(right, right+step);
                if (!needs.containsKey(word))
                    break;
                curr.put(word, curr.getOrDefault(word,0) + 1);
                if (curr.get(word).equals(needs.get(word)))
                    match++;
                right += step;
            }
            if (match == needs.size()) {
                res.add(left);
            }
        }
        return res;
    }
}
```



## 优化：滑动窗口

我的理解，滑动窗口或者双指针本质上就是利用已有的信息减少不必要的搜索。

在这里只需要分别做`wordLen`次滑动窗口即可， `wordLen`为一个单词的长度。

具体的三种情况，子串已经和所有单词完全匹配、遇到不属于这里面的单词或者某个单词出现次数多了，怎么利用前面的信息（看前面的题解）。







