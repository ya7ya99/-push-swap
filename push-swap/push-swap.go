package main

import (
	f "push-swap/funcs"
)

func main() {
	a := f.CheckArgs()
	if a == nil || f.IsSortable(a) {
		return
	}
	b := []int{}
	for {
		if len(a) == 3 {
			break
		}
		a, b = f.SortNums(a, b)
	}
	a = f.LastThree(a)
	for len(b) > 0 {
		a, b = f.Pa(a, b)
	}
}




