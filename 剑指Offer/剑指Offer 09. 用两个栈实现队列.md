### 解题思路一
栈是先进后出，所有的元素都依次压到栈中，但只在栈顶操作，因此每次取出的都是最后一个“进来”的元素；而队列是在队尾插入，队头删除，因此每次取出的都是第一个“进来”的元素。
要使用两个栈来模拟队列，最关键的是要想到将所有元素压入栈中之后，取出的必定是反序，因此再将取出的元素压入另一个栈中，再次取出，那么这样的顺序必定是和插入顺序相同的。

不妨，一个栈作为插入栈，第二个栈作为“临时”存放的栈。

**所以**，`CQueue`可以这样来设计：用一个栈保存队列的所有元素，另一个栈只做“倒腾”用。即插入时只插入这个队列，然后删除时，因为需要取出队列中头个元素，因此将插入栈的元素导入另一个栈中，然后再pop栈顶的元素（即为队列头），再把元素又倒回插入栈。

### 代码
```
class CQueue {
    private LinkedList<Integer> stack1;
    private LinkedList<Integer> stack2;

    public CQueue() {
        stack1 = new LinkedList<>();
        stack2 = new LinkedList<>();
    }
    
    public void appendTail(int value) {
        stack1.push(value);
    }
    
    public int deleteHead() {
        if (stack1 == null || stack1.isEmpty())
            return -1;
        while (!stack1.isEmpty())
            stack2.push(stack1.pop());
        int val = stack2.pop();

        while (!stack2.isEmpty())
            stack1.push(stack2.pop());
        return val;
    }
}

/**
 * Your CQueue object will be instantiated and called as such:
 * CQueue obj = new CQueue();
 * obj.appendTail(value);
 * int param_2 = obj.deleteHead();
 */
```
最后运行时间218ms， 超过12%

**分析**：主要的耗时操作在于反复的将元素倒腾来，倒腾去。那么是否可以减少倒腾的次数呢？注意，第一遍倒过去是必须的，为了得到正确排序的队列从而取出队头，但是我们之后就将排好序的队列（即第二个栈中的元素）又倒回插入栈中，这样实际没有保存我们之前已经排好序的结果，而且之后还要再反复倒腾，浪费！
这里有一个考虑在于，队列的顺序 = 第一部分队列的顺序 + 第二部分队列的顺序 + ...
也就是说，我们将插入栈中的元素倒入第二个栈时，这一部分队列已经得到了（排好了），即使之后还有元素也是排在这一部分队列元素之后的，因此这一部分元素不必再倒回，由此得到我们的优化思路。
### 解题思路二： 优化
同样适用两个栈，第一个栈作为插入栈，第二个栈作为保存队列头一部分元素的栈（具体有多少我们不知道，但是知道这些一定是当前队列的头部元素）。
插入仍然是插入到插入栈中，在删除时对第二个栈进行考虑。如果第二栈不为空，则说明队头的元素还有，直接pop即可；否则，将当前插入栈中的元素倒入第二个栈中作为新的队列首部元素。
当然，加一个特判，判断是否队列为空，注意这里的判断条件和解题思路一的不同，这里队列元素分散在两个栈中，必须同时为空才说明队列为空（而第一种解题，队列的元素都在插入栈中，只需检测插入栈是否为空即可)

### 代码

```java
class CQueue {
    private LinkedList<Integer> stack1;
    private LinkedList<Integer> stack2;

    public CQueue() {
        stack1 = new LinkedList<>();
        stack2 = new LinkedList<>();
    }
    
    public void appendTail(int value) {
        stack1.push(value);
    }
    
    public int deleteHead() {
        if (stack1.isEmpty() && stack2.isEmpty())
            return -1;
        if (stack2.isEmpty())
            while (!stack1.isEmpty())
                stack2.push(stack1.pop());
        int val = stack2.pop();

        return val;
    }
}

/**
 * Your CQueue object will be instantiated and called as such:
 * CQueue obj = new CQueue();
 * obj.appendTail(value);
 * int param_2 = obj.deleteHead();
 */
```