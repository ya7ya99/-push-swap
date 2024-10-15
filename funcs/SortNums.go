package funcs

func SortNums(a, b []int) ([]int, []int) {
	i, num := SearchSmallNum(a)
	lenA := len(a)
	if i <= lenA/2 {
		if i == 1 {
			a = Sa(a)
			return a, b
		}
		for j := 0; j < lenA; j++ {
			lenA = len(a)
			if a[0] != num {
				a = Ra(a)
			} else {
				a, b = Pb(a, b)
				break
			}
		}

	} else {
		for j := 0; j < lenA; j++ {
			lenA = len(a)
			if a[lenA-1] != num {
				a = Rra(a)
			} else {
				a = Rra(a)
				a, b = Pb(a, b)
				break
			}
		}
	}
	return a, b
}

func SearchSmallNum(stack []int) (int, int) {
	min := stack[0]
	var index int
	for i, num := range stack {
		if num < min {
			min = num
			index = i
		}
	}
	return index, min
}
