package sort

type List interface {
	Less(i, j int) bool
	Swap(i, j int)
	Len() int
}

func SelectSort(list List) {
	if nil == list || list.Len() < 2 {
		return
	}
	for i := 0; i < list.Len(); i++ {
		for j := i + 1; j < list.Len(); j++ {
			if !list.Less(i, j) {
				list.Swap(i, j)
			}
		}
	}
}
