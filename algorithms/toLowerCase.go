package algorithms

func toLowerCase(str string) string {
	res := make([]rune, 0)
	for _, c := range str {
		chr := c
		if c >= 'A' && c <= 'Z' {
			chr = c - 'A' + 'a'
		}
		res = append(res, chr)
	}

	return string(res)
}
