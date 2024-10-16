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

func LastThree(a []int) []int {
	first := a[0]
	second := a[1]
	third := a[2]

	// Case 1: 2 1 3
	if first > second && first < third {
		a = Sa(a)
		// Case 2: 3 1 2
	} else if first > second && second < third && first > third {
		a = Ra(a)
		// Case 3: 3 2 1
	} else if first > second && second > third {
		a = Sa(a)
		a = Rra(a)
		// Case 4: 1 3 2
	} else if first < second && second > third && first < third {
		a = Sa(a)
		a = Ra(a)
		// Case 5: 2 3 1
	} else if first < second && second > third && first > third {
		a = Rra(a)
	}
	return a
}

func IsSortable(a []int) bool {
	for {
		if len(a) == 0 {
			break
		}
		i, _ := SearchSmallNum(a)
		if i == 0 {
			a = a[1:]
		} else {
			return false
		}
	}
	return true
}