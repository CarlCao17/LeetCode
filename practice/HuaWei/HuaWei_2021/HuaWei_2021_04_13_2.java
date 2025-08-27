package HuaWei_2021;

import java.util.*;

/*
 * @Author caozs
 * @Description: 给定 URL 的前后缀，诸如 /abc/,/def （以英文逗号分隔），输出合并结果如 /abc/def （结果进行中间分隔符的去重）
 * @Date 5:04 下午 2021/4/14
 * @Param
 * @Return
 **/

public class HuaWei_2021_04_13_2 {
	public static void main(String[] args) {
		Scanner in = new Scanner(System.in);

		String str = in.next();
		String[] url = str.split(",");
		int n = url.length;
		StringBuilder s = new StringBuilder(url[0]);
		for (int i = 1; i < n; i++) {
			if (s.charAt(s.length() - 1) == '/')
				s.deleteCharAt(s.length() - 1);
			s.append(url[i]);
		}
		System.out.println(s.toString());
	}
}