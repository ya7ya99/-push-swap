package funcs

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
			fmt.Println("Error converting argument to integer push swap")
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
