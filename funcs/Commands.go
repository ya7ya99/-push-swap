package funcs

import "fmt"

func Pb(a, b []int) ([]int, []int) {
	b = append(b, a[0])
	a = a[1:]
	fmt.Println("pb")
	return a, b
}

func Pa(a, b []int) ([]int, []int) {
	res := []int{}
	res = append(res, b[len(b)-1])
	res = append(res, a...)
	b = b[:len(b)-1]
	a = res
	fmt.Println("pa")
	return a, b
}

func Sa(a []int) []int {
	a[0], a[1] = a[1], a[0]
	fmt.Println("sa")
	return a
}

func Sb(b []int) []int {
	b[0], b[1] = b[1], b[0]
	fmt.Println("sb")
	return b
}

func Ss(a, b []int) ([]int, []int) {
	a = Sa(a)
	b = Sb(b)
	fmt.Println("ss")
	return a, b
}

func Ra(a []int) []int {
	temp1 := a[1:]
	temp := a[0]
	var res []int
	res = append(res, temp1...)
	res = append(res, temp)
	a = res
	fmt.Println("ra")
	return res
}

func Rb(b []int) []int {
	temp1 := b[1:]
	temp := b[0]
	var res []int
	res = append(res, temp1...)
	res = append(res, temp)
	b = res
	fmt.Println("rb")
	return res
}

func Rr(a, b []int) ([]int, []int) {
	a = Ra(a)
	b = Rb(b)
	fmt.Println("rr")
	return a, b
}

func Rra(a []int) []int {
	temp1 := a[:len(a)-1]
	temp := a[len(a)-1]
	var res []int
	res = append(res, temp)
	res = append(res, temp1...)
	fmt.Println("rra")
	return res
}

func Rrb(b []int) []int {
	temp1 := b[:len(b)-1]
	temp := b[len(b)-1]
	var res []int
	res = append(res, temp)
	res = append(res, temp1...)
	fmt.Println("rrb")
	return res
}

func Rrr(a, b []int) ([]int, []int) {
	a = Rra(a)
	b = Rrb(b)
	fmt.Println("rrr")
	return a, b
}
