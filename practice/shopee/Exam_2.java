//import java.util.Scanner;
//
///*
// * @ClassName: Exam_2
// * @Description:
// * @Author: Caozs
// * @Created: 2021/4/7 3:29 下午
// * @Version: 1.0
// */
//public class Exam_2 {
//
//    public static void main(String[] args) {
//        Scanner in = new Scanner(System.in);
//        int n = in.nextInt();
//        long res = 1 << (n-1);
//
//
//        System.out.println(res % 1000000369);
//    }
//    public int nonValidCodeNum(int n) {
//
//    }
//}
合法的编码数量
        详细描述
        为每个数字定义一个只有”Y“,"N"的编码方式，每一位表示某个数字是i的倍数，i从左往右，从1开始，比如 "YNYYNNY" 表示这个数是1，3，4，7的倍数，但不是2，5，6的倍数。

        现给定正整数 L, 表示某数字的编码后长度，请求出合法的编码的个数总和。

        提示：不合法表示符合编码的数字不存在，比如 "YNNY" 是不合法的，因为不存在是4的倍数但是不是2的倍数的数字

        其他
        时间限制: 1000ms

        内存限制: 256.0MB

        输入输出描述
        输入描述
        输入整数L (1 ≤ L ≤ 10^6),表示数字编码后的长度

        输出描述
        编码合法的编码个数，因为答案可能比较大，所以结果需要对 1000000369 取模

        备注
        例如 n = 4:

        合法的编码有:

        YNNN YNYN YYNY YYYY YYYN YYNN

        所以输出6

        输入输出示例
        示例1
        输入
        复制
        4
        输出
        复制
        6
