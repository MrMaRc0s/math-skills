package main

import (
	"fmt"
	"strconv"
	"strings"
)

func strToInt(file []string) []int {
	var numbers []int
	for _, line := range file {
		// Trim any whitespace
		line = strings.TrimSpace(line)
		if line == "" {
			continue // skip empty lines
		}

		// Convert the line to an integer
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Error converting line to integer:", err)
			continue // or handle the error as needed
		}

		// Append the integer to the slice
		numbers = append(numbers, num)
	}
	return numbers
}
