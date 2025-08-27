# 剑指Offer 45. 把数组排成最小的数

## 快速排序



## 内建函数——自定义排序

```java
class Solution {
    public String minNumber(int[] nums) {
        // int n;
        // if (nums == null || (n = nums.length) == 0)
        //    return "";
        String[] s = new String[n];
        for (int i = 0; i < n; i++)
            s[i] = String.valueOf(nums[i]);
        Arrays.sort(s, (s1, s2) -> {
            return s1.concat(s2).compareTo(s2.concat(s1));
        });
        return String.join("", s);
    }
}
```

## 小顶堆 

```java
class Solution {
    public String minNumber(int[] nums) {
        // int n;
        // if (nums == null || (n = nums.length) == 0)
        //     return "";
  
        Queue<String> queue = new PriorityQueue<>(new Comparator<String>() {
            @Override
            public int compare(String s1, String s2) {
                return (s1 + s2).compareTo(s2 + s1);
            }
        });
        for (int num : nums)
            queue.add(String.valueOf(num));
        
        StringBuilder builder = new StringBuilder();
        while (!queue.isEmpty()) {
            builder.append(queue.poll());
        }
        return builder.toString();
    }
}
```

