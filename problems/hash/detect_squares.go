package hash

// DetectSquares ...
type DetectSquares struct {
	xMap map[int]map[int]bool
	yMap map[int]map[int]bool
}

func Constructor() DetectSquares {
	return DetectSquares{
		xMap: make(map[int]map[int]bool),
		yMap: make(map[int]map[int]bool),
	}
}

func (this *DetectSquares) Add(point []int) {

	if _, ok := this.xMap[point[0]]; !ok {
		this.xMap[point[0]] = make(map[int]bool)
	}
	this.xMap[point[0]][point[1]] = true
	if _, ok := this.yMap[point[1]]; !ok {
		this.yMap[point[1]] = make(map[int]bool)
	}
	this.yMap[point[1]][point[0]] = true
}

func (this *DetectSquares) Count(point []int) int {
	x, y := point[0], point[1]
	count := 0
	chuizhiYs := this.xMap[x]
	for chuizhiY := range chuizhiYs {
		chuizhiShuipingXs := this.yMap[chuizhiY]
		for chuizhishuipingX := range chuizhiShuipingXs {
			shuipingYs := this.xMap[chuizhishuipingX]
			for shuipingY := range shuipingYs {
				if shuipingY == y {
					count++
				}
			}
		}
	}
	return count
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */
