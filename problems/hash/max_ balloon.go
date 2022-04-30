package hash

func maxNumberOfBalloons(text string) int {
	cmap := map[rune]int{
		'b': 0,
		'a': 0,
		'n': 0,
		'l': 0,
		'o': 0,
	}
	for _, c := range text {
		if _, ok := cmap[c]; ok {
			cmap[c] += 1
		}
	}
	cmap['l'] = cmap['l'] / 2
	cmap['o'] = cmap['o'] / 2
	min := 90000
	for _, c := range []rune{'b', 'a', 'n', 'l', 'o'} {
		if cmap[c] <= min {
			min = cmap[c]
		}
	}
	return min
}
