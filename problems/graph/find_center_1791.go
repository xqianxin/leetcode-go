package graph

func findCenter(edges [][]int) int {
	fmap := make(map[int]bool)
	fmap[edges[0][0]] = true
	fmap[edges[0][1]] = true
	if _, ok := fmap[edges[1][0]]; ok {
		return edges[1][0]
	}
	return edges[1][1]
}
