package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		return
	}

	numbers := make(map[int]bool)
	a := []int{}
	// stackB := []int{}

	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("Error converting argument to integer")
			return
		}

		if numbers[num] {
			fmt.Println("Number already exists")
			return
		}

		numbers[num] = true
		a = append(a, num)
	}

	fmt.Println(a)
	// fmt.Printf("Number of unique inputs: %d\n", len(numbers))
}
