package numbers

func reverse(x int) int {
	res := 0
	flag := 1
	intMax := int(^uint32(0) >> 1)
	print(intMax)
	intMin := int(^intMax)
	print(intMin)
	if x < 0 {
		flag = -1
	}

	x = flag * x
	for x != 0 {
		if res > intMax/10 || res < intMin/10 {
			return 0
		}
		res = res*10 + x%10
		x = x / 10
	}
	return flag * res
}
