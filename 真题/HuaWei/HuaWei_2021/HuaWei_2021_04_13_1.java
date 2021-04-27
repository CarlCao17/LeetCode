package HuaWei_2021;

import java.util.*;

/*
 * @ClassName: HuaWei_2021_04_13_1
 * @Description:  一个 0-1000 的整数，拆解为一个（本身）或多个连续自然数的和，按照自然数的个数从少到多输出各个方案 input = solution，方案内的自然数按照从小到大排列
 * @Author: Caozs
 * @Created: 2021/4/14 5:14 下午
 * @Version: 1.0
 */
public class HuaWei_2021_04_13_1 {

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        List<List<Integer>> res = new ArrayList<>();
        int left = 0, right = 0;
        int sum = 0;
        while (right <= n) {
            sum += right;
            right++;

            while (sum >= n) {
                if (sum == n) {
                    int size = right - left;
                    List<Integer> list = new ArrayList<>(size);
                    for (int i = 0; i < size; i++)
                        list.add(left + i);
                    res.add(list);
                }
                sum -= left;
                left++;
            }
        }
        Collections.sort(res, (l1, l2) -> l1.size() - l2.size());
        System.out.println(res);
    }
}
