# 剑指Offer40. 最小的k个数

直接利用$O(nlongn)$的排序我们就不考虑了，当`k`较小时完全没有必要

## 方法一 堆排序

1. 直接自己构造小顶堆，然后交换`k`次，即可取出最小的`k`个元素
2. 利用`Java`中的`PriorityQueue`（小顶堆），构造容量为k的大顶堆，维护最小的`k`个元素

### 代码

#### 自己维护小顶堆

```java
class Solution {
    public int[] getLeastNumbers(int[] arr, int k) {
        int n;
        if (arr == null || (n = arr.length) == 0 || k == 0)
            return new int[0];

        int[] res = new int[k];
        int[] sortedArray = Arrays.copyOf(arr, n);
        for (int i = (n - 1) / 2; i >= 0; i--) {
            adjustHeap(sortedArray, i, n);
        }
        System.out.println(Arrays.toString(sortedArray));
        for (int i = n - 1, cnt = 0; cnt < k ; i--, cnt++) {
            swap(sortedArray, 0, i);
            adjustHeap(sortedArray, 0, i);
            res[cnt] = sortedArray[i];
        }
        return res;
    }

    public void adjustHeap(int[] a, int parent, int len) {
        int temp = a[parent];
        int lchild = 2 * parent + 1;

        while (lchild < len) {
            int rchild = lchild + 1;
            if (rchild < len && a[rchild] < a[lchild])
                lchild++;
            if (temp < a[lchild])
                break;

            a[parent] = a[lchild];
            parent = lchild;
            lchild = 2 * parent + 1;
        }
        a[parent] = temp;
    }

    public void swap(int[] a, int i, int j) {
        int temp = a[i];
        a[i] = a[j];
        a[j] = temp;
    }
}
```

#### 利用`PriorityQueue`

```java
class Solution {
    public int[] getLeastNumbers(int[] arr, int k) {
        int n;
        if (arr == null || (n = arr.length) == 0 || k == 0)
            return new int[0];
        int[] res = new int[k];
        PriorityQueue<Integer> pq = new PriorityQueue<>((o1, o2) -> o2.compareTo(o1));

        for (int num : arr) {
            if (pq.size() < k) {
                pq.offer(num);
            } else if (pq.peek() > num) {
                pq.poll();
                pq.offer(num);
            }
        }
        for (int i = k-1; i >= 0; i--) {
            res[i] = pq.poll();
        }
        return res;
    }
}
```

花费9ms，超过50%

## 利用快排思想

```java
class Solution {
    public int[] getLeastNumbers(int[] arr, int k) {
        int n;
        if (arr == null || (n = arr.length) == 0 || k == 0)
            return new int[0];
        return quickSearch(arr, 0, arr.length-1, k-1);
    }

    /* 从数组a中，范围是[low, high]的区间内用快排的思想找 排序后第k个元素（从零开始计数），并返回前k大的元素即[0, 1, ... k]*/
    public int[] quickSearch(int[] a, int low, int high, int k) {
        int index = partition(a, low, high);
        if (index == k)
            return Arrays.copyOf(a, k+1);
        
        return index > k ? quickSearch(a, low, index-1, k) : quickSearch(a, index+1, high, k);
    }

    public int partition(int[] a, int low, int high) {
        int privot = a[low];
        int i = low, j = high;
        while (i < j) {
            while (i < j && a[j] > privot)
                j--;
            if (i < j)
                a[i++] = a[j];
            while (i < j && a[i] < privot)
                i++;
            if (i < j)
                a[j--] = a[i];
        }
        a[i] = privot;
        return i;
    }
}
```

花费2ms，打败100%

## 利用二叉搜索树

类似于最大堆的做法，只要维护一个K个元素的二叉搜索树，就可以快速找到最小的K个元素

```java
class Solution {
    public int[] getLeastNumbers(int[] arr, int k) {
        int n;
        if (arr == null || (n = arr.length) == 0 || k == 0)
            return new int[0];
        int[] res = new int[k];
        TreeMap<Integer, Integer> bst = new TreeMap<>();

        int cnt = 0;
        for (int num : arr) {
            if (cnt < k) {
                bst.put(num, bst.getOrDefault(num, 0) + 1);
                cnt++;
            } else {
                Map.Entry<Integer, Integer> entry = bst.lastEntry();
                if (num < entry.getKey()) {
                    bst.put(num, bst.getOrDefault(num, 0) + 1);

                    if (entry.getValue() == 1)
                        bst.pollLastEntry();
                    else
                        bst.put(entry.getKey(), entry.getValue() - 1);
                }

            }
        }

        int i = 0;
        for (Map.Entry<Integer, Integer> entry : bst.entrySet()) {
            int key = entry.getKey();
            int freq = entry.getValue();
            while (freq-- > 0) {
                res[i++] = key;
            }
        }
        return res;
    }
}
```

花费32ms，打败 11.96%

## 利用计数排序

