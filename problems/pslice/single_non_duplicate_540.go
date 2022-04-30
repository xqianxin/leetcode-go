package pslice

// You are given a sorted array consisting of only integers where every element appears exactly twice, except for one element which appears exactly once.

// Return the single element that appears only once.

// Your solution must run in O(log n) time and O(1) space.
// 112233455
// 012345678
// before x , 偶数坐标n = n+1, 奇数坐标n = n-1
// after x 偶数坐标 n = n-1, 技术坐标 n=n+1

func singleNonDuplicate(nums []int) int {
	low, high, mid := 0, len(nums)-1, -1
	for low < high {
		mid = (low + high) / 2
		// 偶数坐标
		if mid%2 == 0 {
			if nums[mid] == nums[mid+1] {
				low = mid + 2
			} else {
				high = mid
			}
		} else {
			if nums[mid] == nums[mid-1] {
				low = mid + 1
			} else {
				high = mid
			}
		}
	}
	return nums[low]
}
