import java.util.*;

/*
 * @ClassName: Exam_2
 * @Description: n, k, 翻转区间
 * @Author: Caozs
 * @Created: 2021/4/9 6:51 下午
 * @Version: 1.0
 */
public class Exam_2 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt(), k = in.nextInt();
        int start, end;
        int[] a = new int[n];
        for (int i = 0; i < n; i++)
            a[i] = i+1;
        for (int i = 0; i < k; i++) {
            start = in.nextInt();
            end = in.nextInt();
            rotate(a, start-1, end-1);
        }
        for (int i = 0; i < n; i++) {
            if (i > 0)
                System.out.print(" ");
            System.out.print(a[i]);
        }
    }

    public static void rotate(int[] a, int start, int end) {
        for (int i = start, j = end; i < j; i++, j--) {
            int tmp = a[i];
            a[i] = a[j];
            a[j] = tmp;
        }
    }
}
