# LeetCode 1190.反转每对括号间的子串（栈方法+预处理两种方法 都是100%)

### 解题思路一 栈方法依次反转

将每一层括号都看作一层，最外面是第0层，第一个括号包裹起来的字符串是第1层，等等。例如：`(u(love)i)`，第0层是左右两边的空串加上`(ulove)i)`，第1层是`u(love)i`，第2层是`love`，每深入一层，该层字符串就需要反转，如`love`就需要反转2次才能到第0层。

因此我们可以使用一个字符串来保存当前层的结果（可能处理完了，也可能后面还没处理完），当需要进入下一层时将字符串暂存到栈中，当返回上一层时从栈中弹出。算法如下：
\- 遇到左括号时，需要进入下一层，保存当前层的结果到栈中，当前层字符串清空
\- 遇到右括号时，当前层已处理完，需要反转当前层的结果，并且将反转后的结果与上一层字符串合并（拼接到末尾）
\- 否则，直接加入到当前层的字符串中

### 代码

```java
class Solution {
    public String reverseParentheses(String s) {
        Deque<String> stack = new LinkedList<>();
        StringBuilder str = new StringBuilder();
        int n = s.length();
        for (int i = 0; i < n; i++) {
            char ch = s.charAt(i);
            if (ch == '(') {
                stack.push(str.toString());
                str.setLength(0);
            } else if (ch == ')') {
                str.reverse();
                str.insert(0, stack.pop());
            } else {
                str.append(ch);
            }
        }
        return str.toString();
    }
}
```

**时间复杂度** O(n^2)*O*(*n*2)，由于反向处理需要O(n)*O*(*n*)
**空间复杂度** 需要一个额外的栈，最坏情况下需要O(n)*O*(*n*)复杂度

### 解题思路二 预处理+反序遍历

括号内的字符反转可以看做是反序遍历，我们来看两个例子：

1. s = "(abcd)"*s*="(*a**b**c**d*)"
   最开始处于第0层，无需反序（反转）。当前指向第0个元素'('，需要进入下一层（即反序遍历），跳转到与左括号匹配的右括号处即第5个元素')'然后反向遍历，依次访问'd', 'c', 'b', 'a', 然后又到了'('，此时应当从该层返回到上一层即第0层继续遍历，因此再跳转到与该左括号匹配的右括号处继续遍历（因为括号内的元素即下一层元素都已经访问过了），然后正向遍历到了末尾，结束。
2. 再看一个例子：s = "(u(love)i)"*s*="(*u*(*l**o**v**e*)*i*)"
   最开始处于第0层，正向遍历。当前指向第0个元素'('，进入下一层——第1层，跳转到第9个元素')'，反向遍历，依次遍历第8个元素'i'，然后又遇到第7个元素')'继续进入第2层，需要反向（即变为正向遍历了），跳转到与该括号匹配的第2个元素'('，然后反转方向变为正向，继续遍历第二层元素，依次遍历'l', 'o', 'v', 'e', 再次遇到括号即该层结束，返回上一层即回到上一层的遍历位置及方向——跳转到第2个元素'('，方向继续为反向，然后遍历'u'。再次遇到括号'('，第1层结束，返回第0层遍历位置及方向——跳转到第9个元素，方向为正向，然后结束。

通过这两个例子，我们发现有两种情况：进入下一层（即遇到左括号'('）以及返回上一层（即遇到右括号')'），都需要跳转到相应的位置并且翻转遍历方向，无需区分是进入还是返回。
算法如下：
\- 遇到括号，即跳转到与之匹配的括号处，并且翻转方向
\- 否则在当前位置，按照当前方向继续遍历
\- 直到整个字符串都遍历完
在这里，由于我们定义了最外层——第0层，它的方向一定是正向，因此不论如何最终一定会回到第0层并且正向将第0层访问完，到达字符串末尾。

当然为了方便从一个括号跳到另外一个括号，我们可以先利用栈来预处理，找到每一个括号的匹配位置，使得整个遍历过程只需要O(n)*O*(*n*)

```java
class Solution {
    public String reverseParentheses(String s) {
        int n = s.length();
        int[] pair = new int[n];
        Deque<Integer> stack = new LinkedList<>();
        for (int i = 0; i < n; i++) {
            char ch = s.charAt(i);
            if (ch == '(') 
                stack.push(i);
            else if (ch == ')') {
                int j = stack.pop();
                pair[i] = j;
                pair[j] = i;
            }
        }
        int index = 0, step = 1;
        StringBuilder str = new StringBuilder();
        while (index < n) {
            char ch = s.charAt(index);
            if (ch == '(' || ch == ')') {
                index = pair[index];
                step = -step;
            } else {
                str.append(ch);
            }
            index += step;
        }
        return str.toString();
    }
}
```

**时间复杂度** O(n)*O*(*n*)，预处理O(n)*O*(*n*), 遍历O(n)*O*(*n*)
**空间复杂度** O(n)*O*(*n*)，使用了栈和额外的数组pair*p**a**i**r*保存匹配信息