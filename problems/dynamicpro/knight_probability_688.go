package dynamicpro

// 在一个 n x n 的国际象棋棋盘上，一个骑士从单元格 (row, column) 开始，并尝试进行 k 次移动。行和列是 从 0 开始 的，所以左上单元格是 (0,0) ，右下单元格是 (n - 1, n - 1) 。

// 象棋骑士有8种可能的走法，如下图所示。每次移动在基本方向上是两个单元格，然后在正交方向上是一个单元格。

// 每次骑士要移动时，它都会随机从8种可能的移动中选择一种(即使棋子会离开棋盘)，然后移动到那里。

// 骑士继续移动，直到它走了 k 步或离开了棋盘。

// 返回 骑士在棋盘停止移动后仍留在棋盘上的概率 。

//

// 示例 1：

// 00 01 02 03
// 10 11 12 13
// 20 21 22 23
// 30 31 32 33

// 输入: n = 3, k = 2, row = 0, column = 0
// 输出: 0.0625
// 解释: 有两步(到(1,2)，(2,1))可以让骑士留在棋盘上。
// 在每一个位置上，也有两种移动可以让骑士留在棋盘上。
// 骑士留在棋盘上的总概率是0.0625。

// 示例 2：

// 输入: n = 1, k = 0, row = 0, column = 0
// 输出: 1.00000
// dp[step][i][j] = 1/8 *
func knightProbability(n int, k int, row int, column int) float64 {
	// init
	dp := make([][][]float64, k)
	for step := 0; step < k+1; step++ {
		dp[step] = make([][]float64, n)
		for i := 0; i < n; i++ {
			dp[step][i] = make([]float64, n)
		}
	}
	moves := [][]int{{-2, -1}, {-2, 1}, {2, -1}, {2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}}
	for step := 0; step <= k; step++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if step == 0 {
					dp[step][i][j] = 1
					continue
				}
				prob := float64(0)
				for _, move := range moves {
					pi, pj := i-move[0], j-move[1]
					if pi >= 0 && pi < n && pj >= 0 && pj < n {
						prob += dp[step-1][pi][pj] * 0.125
					}
				}
				dp[step][i][j] = prob
			}
		}
	}
	return dp[k][row][column]
}
