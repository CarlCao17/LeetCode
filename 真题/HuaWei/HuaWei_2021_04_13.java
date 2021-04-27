package HuaWei;

import java.util.*;

public class Main {
	public staic void main(String[] args) {
		Scanner in = new Scanner(System.in);

		int n = in.nextInt();
		String[] url = new String[n];
		for (int i = 0; i < n; i++) {
			url[i] = in.next();
		}

		StringBuilder s = new StringBuilder(url[0]);
		for (int i = 1; i < n; i++) {
			if (s.charAt(s.length() - 1) == '/')
				s.deleteCharAt(s.length() - 1);
			s.append(url[i]);
		}
		System.out.println(s.toString());
	}
}