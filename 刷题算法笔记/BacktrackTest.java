

/*
	result = []
	def backtrack(路径, 选择列表):
		if 满足结束条件:
			result.add(路径)
			return

		for 选择 in 选择列表:
			做选择
			backtrack(路径, 选择列表)
			撤销选择
*/


// 全排列问题
public class Solution {
	private static List<List<Integer>> res = new ArrayList<>();

	List<List<Integer>> permute(int[] nums) {
		int len =  nums.length;
		
		if (len == 0)
			return res;
		List<Integer> trace = new ArrayList<>();
		boolean[] used = new boolean[len];
		backtrack(nums, trace, used, 0, len);
		return res;
	}

	static void backtrack(int[] nums, List<Integer> trace, boolean[] used, int depth, int len) {
		if (depth == len) {
			res.add(new ArrayList(trace));
			return ;
		}

		for (int i = 0; i < len; i++) {
			if (!used[i]) {
				trace.add(nums[i]);
				used[i] = true;
				backtrack(nums, trace, used, depth+1, len);
				used[i] = false;
				trace.remove(trace.size() - 1);
			}
		}
	}
}