package binarySearch;


public class BinarySearchTest 
{
	/* 寻找一个数的二分查找
	 * 该算法存在局限性，比如[1, 2, 2, 2, 3] 要查找2时，返回索引2；
	 * 但是如果我想得到最左边的target或者最右边的target，则只能再向左或者向右线性搜索，不好！
	 */
	public static int binarySearch(int[] nums, int target) {
		int left = 0;
		int right = nums.length - 1;

		while (left <= right) {
			int mid = left + (right - left) / 2;
			if (nums[mid] == target)
				return mid; 				// 直接返回
			else if (nums[mid] < target) 
				left = mid + 1;
			else if (nums[mid] > target)
				right = mid - 1;
		}
		return -1;							// 直接返回
	}


	/* 寻找左侧边界的二分查找
	 * 该算法是在左闭右开区间[left, mid)中进行搜索，返回的left(可能的取值范围是[0, nums.length]表示nums中小于target的元素的个数。
	 * 如果找到则为当前的索引位置，否则为应当放置的位置
	 */
	public static int leftBound(int[] nums, int target) {
		if (nums.length == 0) return -1;
		int left = 0;
		int right = nums.length;	

		while (left < right)   {	
			int mid = left + (right - left) / 2;
			if (nums[mid] == target) 
				right = mid;		
			else if (nums[mid] < target)
				left = mid + 1;
			else if (nums[mid] > target)
				right = mid;		
		}
		return left;
		/* 可以如下修改返回-1
		 if（left == nums.length) return -1;
		 return nums[left] == target ? left : -1; 
	}


	/* 闭搜索区间，寻找左边界值的二分查找 */
	public static int leftBound(int[] nums, int target) {
		int left = 0;
		int right = nums.length - 1;

		while (left <= right) {
			int mid = left + (right - left)/2;
			if (nums[mid] == target) 
				right = mid - 1;				// 别返回，收紧右边界，最后锁定左边界
			else if (nums[mid] < target)
				left = mid + 1;
			else if（nums[mid] > target）
				right = mid - 1;
		}

		if (left >= nums.length || nums[left] != target) // 最后检查left越界情况
			return -1;
		return left;
	}


	/* 左闭右开搜索区间， 寻找右边界的二分查找 
	 * 返回的left - 1可能是满足条件的元素，因为循环结束left必然不是；返回的left
	 */
	public static int rightBound(int[] nums, int target) {
		if (nums.length == 0) return -1;
		int left = 0;
		int right = nums.length;

		while (left < right) {
			int mid = left + (right - left) / 2;
			if (nums[mid] == target)
				left = mid + 1;
			else if (nums[mid] < target)
				left = mid + 1;
			else // nums[mid] > target
				right = mid;
		}

		return left - 1;
		/* 可以如下修改，返回-1；
		 if (left == 0)
		 	return -1;
		 return nums[left-1] == target ? left - 1;
	}


	/* 闭搜索区间，寻找右边界的二分查找 */
	public static int rightBound(int[] nums, int target) {
		int left = 0;
		int right = nums.length - 1;

		while (left <= right) {
			int mid = left + (right - mid)/2;
			if (nums[mid] == target)
				left = mid + 1;					// 别返回，收紧左边界，锁定右边界
			else if (nums[mid] < target)
				left = mid + 1;
			else if (nums[mid] > target)
				right = mid - 1;
		}

		// left - 1 == right
		/* if (left <= 0|| nums[left-1] != target)
			return -1;
		return left - 1;
		可改写为: 
		*/
		if (right < 0 || nums[right] != target) // 最后检查right越界的情况
			return -1;
		return right;
	}


}