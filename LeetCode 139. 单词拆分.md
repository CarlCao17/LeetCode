# LeetCode 139. 单词拆分

```java
class Solution {
    private Set<String> words = new HashSet<>();
    private TreeSet<Integer> lens = new TreeSet<>((i1, i2) -> i2.compareTo(i1));
    public boolean wordBreak(String s, List<String> wordDict) {
 
        for (String word : wordDict) {
            words.add(word);
            lens.add(word.length());
        }
        return wordBreak(s);
    }

    public boolean wordBreak(String s) {
        int n;
        if (s == null || (n = s.length()) == 0)
            return true;

        for (int len : lens) {
            if (len > n)
                continue;
            
            if (words.contains(s.substring(0, len))) {
                if (wordBreak(s.substring(len)))
                    return true;
            }
        }
        return false;
    }
}
```

TLE

