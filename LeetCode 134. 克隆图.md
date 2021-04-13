# LeetCode 134. 克隆图

主要考察图的遍历，BFS和DFS

## BFS 

BFS需要一个辅助队列，还需要一个记录是否已访问的数组`visited`，另外由于需要将原图中的元素和克隆之后的元素对应起来的`HashMap`，可以利用`HashMap`的键来存储`visited`。因为一旦访问过，就需要克隆该元素，所以添加到`visited`可以和`HashMap.put(node, copyNode)`看成一个步骤。

```java
/*
// Definition for a Node.
class Node {
    public int val;
    public List<Node> neighbors;
    public Node() {
        val = 0;
        neighbors = new ArrayList<Node>();
    }
    public Node(int _val) {
        val = _val;
        neighbors = new ArrayList<Node>();
    }
    public Node(int _val, ArrayList<Node> _neighbors) {
        val = _val;
        neighbors = _neighbors;
    }
}
*/

class Solution {
    public Node cloneGraph(Node node) {
        if (node == null)
            return node;
            
        Node cloneNode = new Node(node.val);
        Map<Node, Node> visited = new HashMap<>();
        visited.put(node, cloneNode);
        LinkedList<Node> queue = new LinkedList<>();
        queue.add(node);

        while (!queue.isEmpty()) {
            Node p = queue.remove();
            for (Node neighbor : p.neighbors) {
                if (!visited.containsKey(neighbor)) {
                    visited.put(neighbor, new Node(neighbor.val));
                    queue.add(neighbor);
                }
                visited.get(p).neighbors.add(visited.get(neighbor));
            }
        }
        return cloneNode;
    }
}
```



## DFS 

```java
/*
// Definition for a Node.
class Node {
    public int val;
    public List<Node> neighbors;
    public Node() {
        val = 0;
        neighbors = new ArrayList<Node>();
    }
    public Node(int _val) {
        val = _val;
        neighbors = new ArrayList<Node>();
    }
    public Node(int _val, ArrayList<Node> _neighbors) {
        val = _val;
        neighbors = _neighbors;
    }
}
*/

class Solution {
    private Map<Node, Node> visited = new HashMap<>();
    public Node cloneGraph(Node node) {
        if (node == null)
            return node;

        if (visited.containsKey(node))
            return visited.get(node);

        Node copyNode = new Node(node.val);
        visited.put(node, copyNode);

        for (Node neighbor : node.neighbors) {
            copyNode.neighbors.add(cloneGraph(neighbor));
        }
        return copyNode;
    }
}
```

