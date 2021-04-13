import java.util.Scanner;

/*
 * @ClassName: Exam_3
 * @Description:
 * @Author: Caozs
 * @Created: 2021/4/7 3:59 下午
 * @Version: 1.0
 */
public class Exam_3 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        String all = in.next();
        String[] strs = all.split(";");

        StringBuilder s1 = new StringBuilder(strs[0]);
        StringBuilder s2 = new StringBuilder(strs[1]);

        if (strs[0].length() < strs[1].length()) {
            StringBuilder tmp = s1;
            s1 = s2;
            s2 = tmp;
        }


        StringBuilder res = new StringBuilder();
        for (int i = 0; i < s1.length(); i++)
            res.append('0');

        System.out.println("s1=" + s1.toString());
        System.out.println("s2=" + s2.toString());
        for (int i = s2.length()-1; i >= 0; i--) {
            if (s2.charAt(i) == '1') {
                res = add(s1, res);
                System.out.print(s2.charAt(i) " : " + res.toString());
                System.out.print("=" + res.toString() + "\n");


            }
            System.out.println("res=" + res.toString());
            s1.append('0');
        }
        System.out.println(res.toString());
    }

    public static StringBuilder add(StringBuilder s1, StringBuilder s2) {
        int n = Math.max(s1.length(), s2.length());
        while (s1.length() < n) {
            s1.insert(0, '0');
        }
        while (s2.length() < n) {
            s2.insert(0, '0');
        }
        int carry = 0;
        StringBuilder sum = new StringBuilder(n);
        for (int i = 0; i < n;  i++)
            sum.append('0');

        for (int i = n-1; i >= 0; i--) {
            if (s1.charAt(i) == '1' && s2.charAt(i) == '1') {
                if (carry == 1) {
                    sum.setCharAt(i, '1');
                }
                else {
                    carry = 1;
                }
            }
            else if (s1.charAt(i) == '1' || s2.charAt(i) == '1') {
                if (carry == 1) {
                    carry = 1;
                }
                else {
                    sum.setCharAt(i, '1');
                }
            }
            else {
                if (carry == 1) {
                    sum.setCharAt(i, '1');
                    carry = 0;
                }
            }
        }
        if (carry == 1) {
            sum.insert(0, '1');
        }
        return sum;
    }
}
