package numbers

import (
	"fmt"
)

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53}

func simplifiedFractions(n int) []string {

	resMap := make(map[string]bool, 0)
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			c := j
			m := i
			for _, p := range primes {
				if p > c {
					break
				}
				for c%p == 0 && m%p == 0 {
					c, m = c/p, m/p
				}
			}
			resMap[fmt.Sprintf("%d/%d", c, m)] = true
		}
	}
	res := make([]string, 0)
	for i := range resMap {
		res = append(res, i)
	}
	return res
}
