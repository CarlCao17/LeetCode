package HuaWei;

import java.util.*;

public class Main {
	public static void main(String[] args) {
		Scanner in = new Scanner(System.in);
		int n = in.nextInt();
		int num, type;
		int[] num1 = new int[3];
		int[] num2 = new int[3];
		int[] idx1 = new int[3];
		int[] idx2 = new int[3];

		for (int i = 1; i <= n; i++) {
			num = in.nextInt();
			type = in.nextInt();

			if (type == 1) {
				replace(num1, idx1, num, i);
			} else {
				replace(num2, idx2, num, i);
			}
		}

		int res1 = num1[0] + num1[1] + num1[2];
		int res2 = num2[0] + num2[1] + num2[2];
		Arrays.sort(idx1);
		Arrays.sort(idx2);

		if (res1 == res2) {
			if (idx2[0] < idx1[0]) {
				res1 = res2;
				idx1 = idx2;
			}

		}

		if (res1 >= res2) {
			System.out.println(idx1[0] + " " + idx1[1] + " " + idx1[2]);
			System.out.println(1);
			System.out.println(res1);
		} else if (res1 < res2) {
			System.out.println(idx2[0] + " " + idx2[1] + " " + idx2[2]);
			System.out.println(2);
			System.out.println(res2);
		} 
	}

	public replace(int[] num, int[] idx, int num, int i) {
		if (num > num[2]) {
			down(num, 2);
			down(idx, 2);
			num[2] = num;
			idx[2] = i;
		} else if (num > num[1]) {
			down(num, 1);
			down(idx, 1);
			num[1] = num;
			idx[1] = i;
		}  else if (num > num[0]) {
			num[0] = num;
			idx[0] = i;
		}
	}

	public void down(int[] a, int end) {
		for (int i = 0; i < end; i++) {
			a[i] = a[i+1];
		}
	}
}