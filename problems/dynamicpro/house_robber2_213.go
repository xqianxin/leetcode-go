package dynamicpro

// You are a professional robber planning to rob houses along a street.
// Each house has a certain amount of money stashed. All houses at this place are arranged in a circle.
// That means the first house is the neighbor of the last one. Meanwhile, adjacent houses have a security system connected,
// and it will automatically contact the police if two adjacent houses were broken into on the same night.

// Given an integer array nums representing the amount of money of each house,
// return the maximum amount of money you can rob tonight without alerting the police.

// Example 1:

// Input: nums = [2,3,2]
// Output: 3
// Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2), because they are adjacent houses.
// Example 2:

// Input: nums = [1,2,3,1]
// Output: 4
// Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
// Total amount you can rob = 1 + 3 = 4.
// Example 3:

// Input: nums = [1,2,3]
// Output: 3

func rob2(nums []int) int {
	total := len(nums)
	if total == 0 {
		return 0
	}
	if total == 1 {
		return nums[0]
	}
	if total == 2 {
		return maxInt(nums[0], nums[1])
	}
	dp0 := make([]int, total)
	dp1 := make([]int, total)
	dp0[0] = nums[0]
	dp0[1] = maxInt(nums[0], nums[1])
	dp1[1] = nums[1]
	dp1[2] = maxInt(nums[1], nums[2])
	for i := 2; i < total-1; i++ {
		dp0[i] = maxInt(dp0[i-2]+nums[i], dp0[i-1])
	}
	for i := 3; i < total; i++ {
		dp1[i] = maxInt(dp1[i-2]+nums[i], dp1[i-1])
	}
	return maxInt(dp0[total-2], dp1[total-1])
}
