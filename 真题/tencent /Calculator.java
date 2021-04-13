import java.util.*;

public class Calclator {
	/*
	目前的计算器只支持处理不带括号的四则运算，可以处理空格，不支持大数
	不支持处理除了0-9, \space, +, -, *, /之外的字符
	*/
	public static double calculate (String s) {
		// if (s == null)
		LinkedList<Double> numStack = new LinkedList<>();
		LinkedList<Character> opStack = new LinkedList<>();

		char[] array = s.toCharArray();
		int n = array.length;
		int i = 0;
		char op;
		double num; // just store the value parsed in the String, like 1, 2 and 3 in "1+2+3"
		double num1, num2; // left operand and right operand
		while (i < n) {
			while (i < n && array[i] == ' ') i++;
			if (i >= n)
				break;
			if (array[i] >= '0' && array[i] <= '9') {
				num = 0;
				while (i < n && array[i] >= '0' && array[i] <= '9') {
					num = num * 10 + array[i] - '0';
					i++;
				}
				numStack.push(num);
			}
			

			else if ((op =array[i]) == '*' || op == '/') {
				num1 = numStack.pop();
				num2 = 0;
				
				while (i < n && array[++i] == ' ') ;
				while (i < n && array[i] >= '0' && array[i] <= '9') {
					num2 = num2 * 10 + array[i] - '0';
					i++;
				}	
				if ( i >= n)
					break;
				if (op == '*') num1 *= num2;
				else num1 /= num2;
				numStack.push(num1);
			}
			else {
			// else if (array[i] == '+' || array[i] == '-') {
				if (!opStack.isEmpty()) {
					op = opStack.pop();
					num2 = numStack.pop();
					num1 = numStack.pop();
					if (op == '+') num1 += num2;
					else num1 -= num2;
					numStack.push(num1);
					
				}
				opStack.push(array[i]);
				i++;
			}
			
		}

		while (!opStack.isEmpty()) {
			op = opStack.pop();
			num2 = numStack.pop();
			num1 = numStack.pop();
			if (op == '+') num1 += num2;
			else num1 -= num2;
			numStack.push(num1);
		}
		return numStack.pop();
	}

	public static void main(String[] args) {
		Scanner in = new Scanner(System.in);
		String input;
		System.out.println("Welcome to Calculator!");
		System.out.println("Please enter the expression you want to calculate: ");
		while ((input = in.nextLine()) != null) {
			if ("quit".equalsIgnoreCase(input) || "q".equalsIgnoreCase(input)) {
				System.out.println("Thanks! Exited.");
				break;
			}
			double res = calculate(input);
			System.out.print(input + " = ");
			System.out.printf("%.6f\n", res);
			System.out.println("Please enter the expression you want to calculate: ");
		}
	}
}