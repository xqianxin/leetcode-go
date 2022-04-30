package pslice

func findLUSlength(a string, b string) int {
	dupChar := map[rune]int{}
	for _, e := range a {
		dupChar[e] = 1
	}
	for _, e := range b {
		if _, ok := dupChar[e]; ok {
			dupChar[e] = 2
		}
	}
	la, lb := 0, 0
	for _, e := range a {
		if dupChar[e] != 2 {
			la += 1
		}
	}
	for _, e := range b {
		if dupChar[e] != 2 {
			lb += 1
		}
	}
	if la > lb {
		return la
	}
	return lb
}
