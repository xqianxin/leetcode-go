package pslice

import (
	"fmt"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	fmt.Println(nums)
	i, j := 0, k-1
	min := nums[len(nums)-1] - nums[0]
	for j < len(nums) {
		if nums[j]-nums[i] <= min {
			min = nums[j] - nums[i]
		}
		fmt.Println(min, i, j)
		i++
		j++
	}
	return min
}
