package main

import (
	"math"
)

// Function to calculate the standard deviation of a list of integers and return it as an integer
func calculateStandardDeviation(numbers []int) int {
	// Step 1: Calculate the mean
	var sum int
	for _, num := range numbers {
		sum += num
	}
	mean := float64(sum) / float64(len(numbers))

	// Step 2: Calculate the variance
	var squaredDiffsSum float64
	for _, num := range numbers {
		diff := float64(num) - mean
		squaredDiffsSum += diff * diff
	}
	variance := squaredDiffsSum / float64(len(numbers))

	// Step 3: Calculate the standard deviation and return as int
	stdDev := math.Sqrt(variance)
	return int(math.Round(stdDev)) // Round to the nearest integer
}
