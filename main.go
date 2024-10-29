package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// check if the program have only 1 arg
	if len(os.Args) != 2 {
		fmt.Println("u need to pass an argument")
		os.Exit(1)
	}

	txt, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("something is wrong with the txt file")
		os.Exit(2)
	}
	// sort our list
	file := strings.Split(string(txt), "\n")
	numbers := strToInt(file)
	numbers = SortIntegerTable(numbers)

	fmt.Println("Average: ", calculateAverage(numbers))
	fmt.Println("Median: ", calculateMedian(numbers))
	fmt.Println("Variance: ", calculateVariance(numbers))
	fmt.Println("Standard Deviation: ", calculateStandardDeviation(numbers))
}
