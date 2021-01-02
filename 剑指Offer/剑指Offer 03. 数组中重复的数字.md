# 剑指Offer 03. 数组中重复的数字

## 解法一 先排序再查找

时间复杂度： `O(nlogn)`

空间复杂度: `O(1)`

### API讲解

- `java.util.Arrays `类

  `public class Arrays extends Object`
  这个类包含许多操作数组的方法(比如排序，搜索)，并且还包含一个静态工厂方法，以允许将数组(arrays)视为列表(lists)。
  这个类中的方法都会抛出空指针异常，如果指定的数组引用为空，例外情况如Note所示。

  **重要方法** :

  - `public static <T> List<T> asList(T... a)`

    返回一个由特定数组支持的列表(list)（实际上是特定数组的列表视图）。对返回的列表做的修改，都会"写回"到数组中。这一方法作为基于数组和基于集合的API的桥梁，和`Collection.toArray()`一起。返回的数组时序列化的 (serializable) 并且实现了`RandomAccess`。这个方法同样提供了便捷的方法来创建一个固定大小且初始化的列表：

    `List<String> stooges = Arrays.asList("Larry", "Moe", "Curly");`

  - `public static int hashCode(long[] a)`

    返回基于特定数组内容的hash code值。对于任何两个`long` 数组`a`和`b`，若`Arrays.equals(a, b)`, 那么同样有`Arrays.hashCode(a) == Arrays.hashCode(b)`。该方法返回的值同样等于调用具有相同元素及相同顺序的Long实例序列的列表`List`的`hashCode`方法的返回值。如果`a`为空，则返回0。

    类似的还有支持：`boolean[] A`,`byte[] A]`, `char[] A`, `double[] A`, `float[] A`, `int[] A`, `long[] A`, `Object[] A`和`short[] A`。

  - `public static int binarySearch(byte[] a, byte key)`

    使用二分查找算法在特定`byte`数组中搜索特定值。该数组必须是在调用该函数之前被排序（如`sort(byte[]) `方法）。如果未排序，则结果未定义。如果数组包含特定值的多个元素，则无法保证获取哪一个。

    **返回值**：搜索的键的索引，如果其值在数组中；否则，返回`(-(insertion point) - 1)`。`insertion point`指的是这个`key`本应该被插入到数组中的地方：第一个大于该`key`值的索引，或者是`a.length`，如果数组中所有元素都小于指定的key。注意这一点保证了当且仅当key被找到时，才返回$>=0$

    该API还有如：`public static binarySearch(byte[] A, int start, int end, byte key)`形式

    并且还支持：`double[]`, `float[]`, `int[]`, `long[]`, `Object[]`, `short[]`, `T[]`

  - `public static boolean[] copyOf(boolean[] original, int newLength)`

    复制指定数组，截断或者用`false`填充，保证复制后的数组`copy`有指定长度。对于所有在原始数组和复制数组`copy`中都合法的索引，两个数组有相同值。对于任何在`copy`中合法但在原始数组中不合法，`copy`数组将包含`false`值。这些索引当且仅当指定的长度超过原始数组的长度。

    还支持以下类型：`byte[] `, `char[] `, `double[]`, `float[]`, `int[]`, `long[]`, `short[]`, `T[] `

    同时还有：`public static <T, U> T[] copyOf(U[] original, int newLength, Class<? Extends T[]> newType)`

    `copyOfRange(boolean[] A, int, int)`

  - `public static boolean equals(byte[] a, byte[] a2)`

    如果两个数组等于另外一个，则返回`true`。两个数组被认为是相等的，如果连个数组包含相同个数的元素，并且所有的元素对都相等。换句话说，如果他们以相同的顺序包含相同的元素，则两个数组相等。并且如果两个都为`null`，则两个数组引用被认为是相等的。

    同样还支持其他基本类型及Object数组

  - `public static void sort(byte[] a)`

    将特定数组排成升序数字顺序。

    `Implementation note`: 排序算法为双`pivot`的快排。这个算法提供了`O(nlogn)`的性能在许多可能使其他快排算法退化到二次性能的数据集上，并且比传统的一个`pivot`的快排实现要更快。

    还支持部分排序和其他基本类型及Object数组，及参数类型`T[]`

    `public static <T> void sort(T[] a,  int from, int to, Comparator<? super T> c)`

```java
class Solution {
    public int findRepeatNumber(int[] nums) {
        Arrays.sort(nums);
        int res = nums[0];
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] == nums[i-1]) {
                res = nums[i];break;
            }
        }
        return res;
    }
}
```

## 解法二 使用hashmap

```java
class Solution {
    public int findRepeatNumber(int[] nums) {
        Map<Integer, Integer> hashmap = new HashMap<>();
        int res = nums[0];
        for (int i : nums) {
            if (! hashmap.containsKey(i))
                hashmap.put(i, 0);
            else {
                res = i;
                break;
            }
        }
        return res;
    }
}
```

## 解法三 数组重排

思路：

只有一部分是重复的，其他都是不重复的，且在`[0, n-1]`之间。如果数组中没有重复的数字，那么在数组排序之后，数字`i`必然出现在下标`i`的位置。由于数组有重复的数，则有的位置会有多个数，有的位置没有数。

现在重排这个数组，从头到尾扫描一遍。假设扫描到下标`i`，不妨记该位置的数为`m`。则若`m == i`，说明该位置已经排好了，继续扫描下一个数字；否则，需要找到在其他地方的数`i`，将之放到下标`i`处。因此，如下处理：判断第`i`个数字(即`m`)和第`m`个数字，如果相等，说明找到重复的数了；否则，交换两数，保证第`m`个数字已经放到正确的位置。再试探交换过来的数是不是`i`，如果还不是，继续交换，这样交换下去，要么找到`i`，要么找到重复的数字。

如此遍历整个数组，只要保证数组中有重复数，必然能找到。

可以如下来说明，走到下标`i`处，则必然`0-(i-1)`都已经排好了，并且下标是`k， 0<=k<=i-1`，且`nums[k] = k`。则在对下标`i`排序的过程中，分几种情况：

- 如果`nums[i] < i`，那么必然重复了，找到了
- 如果`nums[i] == i`，该位置已经排好了
- 否则，需要排序（即去找数字`i`），在这过程中，又要交换，这样，必然归于前面两种情况或者把`i`找到。这个过程的出口就是`i`一直循环到`nums.length`，又由于题目已经保证必然有重复数，所以一定能找到。

其实，如果没有重复数，则这就是对数组进行排序。

**或者**简而言之，我在排序的过程中，一定将数字`i`放到下标`i`处，所以，一旦有重复，必然可以轻松检查到。如果没有重复，那就会把整个数组排好序。

时间复杂度 O(n)：每个数字都至多交换**2次**就能找到属于它自己的位置**（是至多2次，还是1次？）**

空间复杂度 O(1)

```java
class Solution {
    public int findRepeatNumber(int[] nums) {
        int i = 0;
        int res = nums[0];
        int tmp;
        while (i < nums.length) {
            if (nums[i] == i) {
                i++; continue;
            }
            else if (nums[i] == nums[nums[i]]) {
                res = nums[i]; break;
            }
            
            tmp = nums[i];
            nums[i] = nums[tmp];
            nums[tmp] = tmp;
            
        }
        return res;
    }
}
```

