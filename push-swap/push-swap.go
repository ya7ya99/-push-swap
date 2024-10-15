package main

import (
	"fmt"

	f "push-swap/funcs"
)

func main() {
	a := f.CheckArgs()
	if a == nil || isSortable(a) {
		return
	}
	b := []int{}
	for {
		if len(a) == 1 {
			break
		}
		a, b = f.SortNums(a, b)
	}
	a, b = f.Pb(a, b)
	a = b
	b = []int{}
	fmt.Println(a)
}

func isSortable(a []int) bool {
	for {
		if len(a) == 0 {
			break
		}
		i, _ := f.SearchSmallNum(a)
		if i == 0 {
			a = a[1:]
		} else {
			return false
		}
	}
	return true
}
