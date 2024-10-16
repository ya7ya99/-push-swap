package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    a := CheckArgs()
	if a == nil {
		return
	}
	b := []int{}
	scanner := bufio.NewScanner(os.Stdin)
    data := ""
	commandes := []string{"pa","pb","sa","sb","ss","ra", "rb","rr", "rra", "rrb", "rrr"}
	for scanner.Scan() {
		f := strings.TrimSpace(scanner.Text())
		var check = false
		for _, com := range commandes {
			if f == com {
				check = true
				break
			}
		}
		if !check {
			fmt.Println("Invalid commande")
			os.Exit(1)
		}
		data += f+"\n"
	}

    sli := strings.Split(data, "\n")
    for _, d := range sli {
        if d!="" {
            switch d {
            case "pa":
                if(len(b) == 0) {
                    continue
                }
                a, b = Pa(a, b)
            case "pb":
                if(len(a) == 0) {
                    continue
                }
                a, b = Pb(a, b)
            case "sa":
                if(len(a) == 0) {
                    continue
                }
                a = Sa(a)
            case "sb":
                if(len(b) == 0) {
                    continue
                }
                b = Sb(b)
            case "ss":
                if(len(a) == 0) {
                    continue
                }
                if(len(b) == 0) {
                    continue
                }
                a, b = Pa(a, b)
            case "ra":
                if(len(a) == 0) {
                    continue
                }
                a = Ra(a)
            case "rb":
                if(len(b) == 0) {
                    continue
                }
                b = Rb(b)
            case "rr":
                if(len(a) == 0) {
                    continue
                }
                if(len(b) == 0) {
                    continue
                }
                a, b = Rr(a, b)
            case "rra":
                if(len(a) == 0) {
                    continue
                }
                a = Rra(a)
            case "rrb":
                if(len(b) == 0) {
                    continue
                }
                b = Rrb(b)
            case "rrr":
                if(len(a) == 0) {
                    continue
                }
                if(len(b) == 0) {
                    continue
                }
                a, b = Rrr(a, b)
            default:
                fmt.Println("Invalid commande")
            }
        }
    }
	fmt.Println(a, b)
	if (len(a) == 0 && len(b) != 0) {
		a = b
		b = []int{}
	}
    if isSortable(a) && len(b) == 0 {
        fmt.Println("OK")
        return
    }else {
        fmt.Println("KO")
        return
    }
}

func CheckArgs() []int {
	args := os.Args[1:]
	if len(args) < 1 || len(args) > 1 {
		return nil
	}
	numbers := make(map[int]bool)
	a := []int{}
	sli := strings.Split(args[0], " ")
	for _, arg := range sli {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("Error converting argument to integer chekcer")
			os.Exit(1)
		}
		if numbers[num] {
			fmt.Println("Number already exists")
			os.Exit(1)
		}
		numbers[num] = true
		a = append(a, num)
	}
	return a
}

func Pb(a, b []int) ([]int, []int) {
	b = append(b, a[0])
	a = a[1:]
	return a, b
}

func Pa(a, b []int) ([]int, []int) {
	res := []int{}
	res = append(res, b[len(b)-1])
	res = append(res, a...)
	b = b[:len(b)-1]
	a = res
	return a, b
}

func Sa(a []int) []int {
	a[0], a[1] = a[1], a[0]
	return a
}

func Sb(b []int) []int {
	b[0], b[1] = b[1], b[0]
	return b
}

func Ss(a, b []int) ([]int, []int) {
	a = Sa(a)
	b = Sb(b)
	return a, b
}

func Ra(a []int) []int {
	temp1 := a[1:]
	temp := a[0]
	var res []int
	res = append(res, temp1...)
	res = append(res, temp)
	a = res
	return res
}

func Rb(b []int) []int {
	temp1 := b[1:]
	temp := b[0]
	var res []int
	res = append(res, temp1...)
	res = append(res, temp)
	b = res
	return res
}

func Rr(a, b []int) ([]int, []int) {
	a = Ra(a)
	b = Rb(b)
	return a, b
}

func Rra(a []int) []int {
	temp1 := a[:len(a)-1]
	temp := a[len(a)-1]
	var res []int
	res = append(res, temp)
	res = append(res, temp1...)

	return res
}

func Rrb(b []int) []int {
	temp1 := b[:len(b)-1]
	temp := b[len(b)-1]
	var res []int
	res = append(res, temp)
	res = append(res, temp1...)

	return res
}

func Rrr(a, b []int) ([]int, []int) {
	a = Rra(a)
	b = Rrb(b)

	return a, b
}

func isSortable(a []int) bool {
	for {
		if len(a) == 0 {
			break
		}
		i := SearchSmallNum(a)
		if i == 0 {
			a = a[1:]
		} else {
			return false
		}
	}
	return true
}

func SearchSmallNum(stack []int) (int) {
	min := stack[0]
	var index int
	for i, num := range stack {
		if num < min {
			min = num
			index = i
		}
	}
	return index
}
