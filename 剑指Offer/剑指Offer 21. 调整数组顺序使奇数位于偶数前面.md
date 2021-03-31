# 剑指Offer 21. 调整数组顺序使奇数位于偶数前面

### 原地置换： 置换偶数

```java
class Solution {
    public int[] exchange(int[] nums) {
        int n = nums.length;

        int i = 0, j = n - 1;
        while (i < j) {
            if ((nums[i] & 1) == 0) {
                int temp = nums[j];
                nums[j--] = nums[i];
                nums[i] = temp;
            }
            else 
                i++;
        }
        return nums;
    }
}
```

### 原地置换: 置换奇数(又叫快慢指针）

```java
class Solution {
    public int[] exchange(int[] nums) {
        int n = nums.length;

        int low = 0, fast = 0
        while (low <= fast) {
            if ((nums[fast] & 1) == 1) {
                int temp = nums[low];
                nums[low] = nums[fast];
                nums[fast] = temp;
                low++;
            }
            else 
                fast++;
        }
        return nums;
    }
}
```

## 