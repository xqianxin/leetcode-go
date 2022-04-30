package pslice

// You are given an integer array nums. The unique elements of an array are the elements that appear exactly once in the array.

// Return the sum of all the unique elements of nums.

func sumOfUnique(nums []int) int {
	umap := make(map[int]int)
	for _, num := range nums {
		umap[num] += 1
	}
	res := 0
	for num, c := range umap {
		if c == 1 {
			res += num
		}
	}
	return res
}
