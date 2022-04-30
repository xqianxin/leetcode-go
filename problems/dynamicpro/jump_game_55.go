package dynamicpro

// You are given an integer array nums. You are initially positioned at the array's first index,
// and each element in the array represents your maximum jump length at that position.

// Return true if you can reach the last index, or false otherwise.

// Example 1:

// Input: nums = [2,3,1,1,4]
// Output: true
// Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
// Example 2:

// Input: nums = [3,2,1,0,4]
// Output: false
// Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.

// Constraints:

// 1 <= nums.length <= 104
// 0 <= nums[i] <= 105

func canJump(nums []int) bool {
	total := len(nums)
	dp := make([][]bool, 0)
	for i := 0; i < total; i++ {
		row := make([]bool, 0)
		for j := 0; j < total; j++ {
			row = append(row, false)
		}
		dp = append(dp, row)
	}
	dp[0][0] = true
	for i := 0; i < total; i++ {
		jlen := nums[i]
		for j := 1; j <= jlen && i+j < total; j++ {
			dp[i][j] = dp[i][i]
		}
	}
	return dp[total][total]
}

// class Solution {
// 	public:
// 		bool canJump(vector<int>& nums) {
// 			int n = nums.size();
// 			if(n==1)
// 				return true;
// 			int next = nums[n-1], nextIdx=n-1;

// 			for(int i=n-2; i>=0; i--) {
// 				if(nums[i] + i >= nextIdx) {
// 					next = nums[i];
// 					nextIdx = i;
// 				}
// 			}
// 			if(nextIdx == 0) {
// 				return true;
// 			} else {
// 				return false;
// 			}
// 		}
// 	};
