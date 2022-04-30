package pslice

// 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。

// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

func RemoveDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			j += 1
			continue
		}
		if nums[i] != nums[i-1] {
			j += 1
		}
	}
	return j
}
