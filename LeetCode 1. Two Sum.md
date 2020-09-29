# LeetCode 1. Two Sum

## 题目描述

给定一个整数数组`nums`和一个整数`target`，请返回两个和为`target`的数的下标。

你可以假设每个输入只有一个解，并且你不可以使用同一个元素两次。



## 解法一 暴力搜索

```java
class Solution {
  public int[] twoSum(int[] nums,  int target) {
    for (int i = 0; i < nums.length; i++) {
      for (int j = i + 1; j < nums.length; j++) {
        if (nums[i] + nums[j] == target) 
          return new int[] {i, j};
      }
    }
    return new int[2];
  }
}
```



## 解法二 使用HashMap Two Pass

```java
class Solution {
  public int[] twoSum(int[] nums, int target) {
    Map<Integer, Integer> map = new HashMap<>();
    for (int i = 0; i < nums.length; i++)
      map.put(nums[i], i);
    
    for (int i = 0; i < nums.length; i++) {
      if (map.containsKey(target - nums[i]) && map.get(target-nums[i]) != i)
        return new int[] {i, map.get(target-nums[i])};
    }
    
    return new int[2];
  }
}
```

有一个疑问：HashMap的键(Key)是惟一的，如果样例如下：

Example:

Input: [3, 3]

​		  6

Output: [0, 1]

说明：如果有重复元素，则第2个重复元素会将第1个重复元素覆盖，即 map = {3 : 0} -> map = {3 : 1}

而我们搜索是从前往后搜，因此若map包含这个键，且map.get() != i是成立的，对于重复键的情况

## 使用HashMap One Pass

```java
class Solution {
  public int[] two Sum(int[] nums, int target) {
    Map<Integer, Integer> map = new HashMap<>();
    
    for (int i = 0; i < nums.length; i++) {
      if (map.containsKey(target - nums[i]))
        return new int[] {map.get(target - nums[i]), i};
      map.put(nums[i], i);
    }
    
    return new int[2];
  }
}
```

说明：

从前往后搜索，如果不满足条件，则将键值对放到HashMap中，注意：若存在解，则必然一前一后，所以搜索到第一个值则必然不满足条件，一定将它放进去，搜索到其他值都不会满足条件，一直到搜索到第二个值，满足条件取出值并返回即可。【这一定是建立在唯一解的假设下】

## 如果假设不唯一呢？或者解不一定存在呢？

1. 解<= 1的情况是一样的。

2. 解>=2的情况：要求输出任意一对，则也一样

3. 解>=2，且要求输出所有满足的情况：

   则遍历整个数组之后再返回，并且以`Map<Integer, ArrayList<Integer>>`保存键为值，值为序号列表（翻转，若重复，则保存在列表中），将所有可能都组合起来

   ```java
   class Solution {
     public List<int[]> twoSum(int[] nums, int target) {
           Map<Integer, ArrayList<Integer>> map = new HashMap<>();
           List<int[]> res = new ArrayList<>();
   
           for (int i = 0; i < nums.length; i++) {
               if (map.containsKey(target - nums[i]))
                   for (Integer index : map.get(target - nums[i]))
                       res.add(new int[]{index, i});
               if (map.containsKey(nums[i]))
                   map.get(nums[i]).add(i);
               else {
                   ArrayList<Integer> tmp = new ArrayList<>();
                   tmp.add(i);
                   map.put(nums[i], tmp);
               }
           }
           return res;
       }
   }
   ```

   



## Python 题解

```python
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        num_dict = {}
        for i, num in enumerate(nums):
            other_num = target - num
            if other_num not in num_dict:
                num_dict[num] = num_dict.get(num, []) ## 这一步没有必要，因为解唯一
                num_dict[num].append(i) ## num_dict[num] = i
            else:
                idx = num_dict[other_num]
                return [idx[0], i]
```

