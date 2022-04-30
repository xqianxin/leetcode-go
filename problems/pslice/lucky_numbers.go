package pslice

func luckyNumbers(matrix [][]int) []int {
	minMap := make(map[int]bool, 0)
	for i := 0; i < len(matrix); i++ {
		minNode := 9999999
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] < minNode {
				minNode = matrix[i][j]
			}
		}
		minMap[minNode] = true
	}
	maxMap := make(map[int]bool)
	for j := 0; j < len(matrix[0]); j++ {
		maxNode := -1
		for i := 0; i < len(matrix); i++ {
			if matrix[i][j] > maxNode {
				maxNode = matrix[i][j]
			}
		}
		maxMap[maxNode] = true
	}
	res := make([]int, 0)
	for min := range minMap {
		if _, ok := maxMap[min]; ok {
			res = append(res, min)
		}
	}
	return res
}
