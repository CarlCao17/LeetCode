# LeetCode 15. 三数之和

求数组中所有三树之和为零的三元组，要求三元组不重复



### 暴力解法

三重循环暴力求解，但如何保证不重复呢？首先给数组排个序，然后令所有的循环都是有顺序的，比如`i = c, j = i+1...n,k = j+1...n `，并且保证`i`、`j`和`k`每次循环不会重复找相同数，这样才能保证三元组是唯一的。

栗子：`11223`，`i = 1`时，`j`可以取第1个2，但不能取第二个2，同理，`i`不可以取第二个1，但是注意`i = 1`时，`j`要可以取第二个1，因为这不矛盾。

由上，我们知道要加上这三个判断,框架如下：

```java
for (int i = 0; i < n; i++) {
  	if (i > 0 && nums[i] == nums[i-1])
			continue;
    
		for (int j = i + 1; j < n; j++) {
      	if (j > i+1 && nums[j] == nums[j-1])
          	continue;
      	for (int k = j + 1; k < n; k++)
          if (k > j + 1 && nums[k] == nums[k - 1])
              continue;
          ...
    }
}

```

或者用`HashSet`过滤一遍所有的重复元素。

```java
class Solution {
    public List<List<Integer>> threeSum(int[] nums) {
        int n;
        List<List<Integer>> res = new ArrayList<>();
        Set<List<Integer>> set = new HashSet<>();
        if (nums == null || (n = nums.length) == 0)
            return res;
        Arrays.sort(nums);
        for (int i = 0; i < n; i++) {
            for (int j = i+1; j < n; j++) {
                int sum = nums[i] + nums[j];
                for (int k = j+1; k < n; k++)
                    if (nums[k] == -sum) {
                        List<Integer> list = Arrays.asList(nums[i], nums[j], nums[k]);
                        if (set.contains(list))
                            break;
                        res.add(list);
                        set.add(list);
                    }
            }
        }
        return res;
    }
}
```

##### TLE

时间复杂度$O(n^3)$：最坏情况下还是要三重遍历，尽管做了一些优化



### 双指针+ 排序

碰撞双指针必定是要有序的，才能使用