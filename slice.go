package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var s []int = make([]int, 3)
	var input string
	fmt.Println("Enter a number(X to exit):")
	for {
		fmt.Scanln(&input)
		if input == "C" {
			break
		}

		ap, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		s = append(s, ap)
		sort.Ints(s[:])
		fmt.Println(s)
		fmt.Println("Again enter a number(X to exit):")
	}
}
