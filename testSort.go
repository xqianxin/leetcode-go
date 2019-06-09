package main

import (
	"fmt"
	"leetcode-go/sort"
	"math/rand"
)

type intList []int

func (l intList) Len() int {
	return len(l)
}

func (l intList) Less(i, j int) bool {
	return l[i] < l[j]
}

func (l intList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func testSelectSort() {
	testList := make(intList, 0)
	for i := 0; i < 10; i++ {
		tmp := rand.Int() % 100
		testList = append(testList, tmp)
	}
	sort.SelectSort(testList)
	fmt.Println(testList)
}

func main() {
	testSelectSort()
}
