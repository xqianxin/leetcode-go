package search

import (
	"fmt"
	"sort"
)

// 2009. Minimum Number of Operations to Make Array Continuous
// Add to List

// Share
// You are given an integer array nums. In one operation, you can replace any element in nums with any integer.

// nums is considered continuous if both of the following conditions are fulfilled:

// All elements in nums are unique.
// The difference between the maximum element and the minimum element in nums equals nums.length - 1.
// For example, nums = [4, 2, 5, 3] is continuous, but nums = [1, 2, 3, 5, 6] is not continuous.

// Return the minimum number of operations to make nums continuous.

// Example 1:

// Input: nums = [4,2,5,3]
// Output: 0
// Explanation: nums is already continuous.
// Example 2:

// Input: nums = [1,2,3,5,6]
// Output: 1
// Explanation: One possible solution is to change the last element to 4.
// The resulting array is [1,2,3,5,4], which is continuous.
// Example 3:

// Input: nums = [1,10,100,1000]
// Output: 3
// Explanation: One possible solution is to:
// - Change the second element to 2.
// - Change the third element to 3.
// - Change the fourth element to 4.
// The resulting array is [1,2,3,4], which is continuous.

// Constraints:

// 1 <= nums.length <= 105
// 1 <= nums[i] <= 109

func minOperations(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	numMap := make(map[int]int)
	for idx, num := range nums {
		numMap[num] = idx
	}
	fmt.Println(numMap)
	res := len(nums)
	for num, idx := range numMap {
		gap := num - idx
		count := len(nums)
		for j := 0; j < len(nums); j++ {
			if _, ok := numMap[j+gap]; ok {
				count--
			}
		}
		fmt.Println(num, idx, ": ", count)
		if count < res {
			res = count
		}
	}
	return res
}
