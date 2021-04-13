/*使用任意语言实现一种算法来对pai值进行估算*/
public class EvaluatePI {
	public static void main(String[] args) {
      int n = 1000000;
      Random rand = new Random(1234);
      double x, y, distance;
      int cnt = 0;
      double p;
      for (int i = 0; i < n;i++) {
        x = rand.nextDouble(2) - 1;
        y = rand.nextDouble(2) - 1;
        distance = calDistance(x, y);
        if (x >= 0 && y >= 0 && distance <= 1) 
          cnt++;
      }
      p = cnt / n;
      System.out.println(16 * p);
    }
          
  	public double calDistance(double x, double y) {
          return Math.sqrt(Math.pow(x-0, 2) + Math.pow(y-0, 2));
    }
      
}


/*
问题：一个有序链标，里面可能有相同的元素，问题
1. 把重复的元素去掉，只剩下一个，例如1->2->2->3->3->4 去掉相同元素后为 1->2->3->4
2. 如果有重复的元素，则都去掉，例如1->1->2->2->3->3->4->5 去掉相同元素后为 4->5
1->2->3->3->4->5
请用你熟悉的语言实现上述函数
*/
public class Solution {
  class Node {
    int val;
    Node next;
    
    public Node() {}
    
    public Node(int val) { this.val = val; }
  }
  
  public Node solution1(Node head) {
    	Node dum = new Node();
    	dum.next = head;
    	Node p = dum.next, prev = dum;
    	while (p != null) {
          if (p.next != null && p.val == p.next.val) {
            prev.next = p.next;
            p = p.next;
          } else {
            p = p.next;
          	prev = prev.next;
          }
          
        }
    	return dum.next;
  }
  
  public Node solution2(Node head) {
    	Map<Integer, Integer> map = new HashMap<>();
    	Node p = head;
    	while (p != null) {
          map.put(p.val, map.getOrDefault(p.val, 0) + 1);
        }
    	
    	
    	Node dumn = new Node(0);
    	dumn.next = head;
    	Node prev = dumn;
    	p = head;
    	while (p != null) {
          int k = map.get(p.val);
          for (int i = k; i > 1; i++) {
              p = p.next;
          }
          if (k > 1) {
          	prev.next = p.next;
          } else {
            prev = p;
          }
          p = p.next;
        }
    	return dumn.next;
  }
}