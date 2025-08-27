import java.util.*;

/*
 * @ClassName: Exam_1
 * @Description:
 * @Author: Caozs
 * @Created: 2021/4/9 6:51 下午
 * @Version: 1.0
 */
public class Exam_1 {
    static class Node {
        int arrive;
        int leave;

        public Node(int a, int l ) { arrive = a; leave = l; }
    }

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);

        int t = in.nextInt();
        for (int i = 0; i < t; i++) {
            int n = in.nextInt();
            int arrive, leave;
            int time = 1;
            int[] res = new int[n];
            LinkedList<Node> queue = new LinkedList<>();

            for (int j = 0; j < n; j++) {
                arrive = in.nextInt();
                leave = in.nextInt();
                queue.add(new Node(arrive, leave));
            }
            Collections.sort(queue, (o1, o2) -> o1.arrive - o2.arrive);

            for (int j = 0; j < n; j++) {
                Node p = queue.remove();
                if (time < p.arrive) time = p.arrive;
                if (time > p.leave)  res[j] = 0;
                else {
                    res[j] = time;
                    time++;
                }
            }
            for (int j = 0; j < n; j++) {
                if (j > 0)
                    System.out.print(" ");
                System.out.print(res[j]);
            }
            System.out.println();
        }
    }
}
