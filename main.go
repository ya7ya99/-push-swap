package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please enter numbers separated by newlines (Ctrl+D to finish):")

	scanner := bufio.NewScanner(os.Stdin)
	numbers := make(map[int]bool)
	var a []int

	
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue 
		}

		numStrings := strings.Split(line, "\n")
		for _, numStr := range numStrings {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Error converting input to integer:", err)
				continue
			}

			
			if numbers[num] {
				fmt.Println("Number already exists:", num)
				continue
			}

			numbers[num] = true
			a = append(a, num)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println("Unique numbers:", a)
	fmt.Printf("Number of unique inputs: %d\n", len(numbers))
}
