package pslice

func maxArea(height []int) int {
	var xmin = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var xmax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	i, j := 0, len(height)-1
	ret := 0
	for i < j {
		area := xmin(height[i], height[j]) * (j - i)
		ret = xmax(area, ret)
		if height[i] <= height[j] {
			i++
		} else {
			j--
		}
	}
	return ret
}
