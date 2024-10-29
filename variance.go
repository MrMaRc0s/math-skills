package main

import "math"

func calculateVariance(numbers []int) int {
	// Calculate the mean
	var sum int
	for _, num := range numbers {
		sum += num
	}
	mean := float64(sum) / float64(len(numbers))

	// Calculate the squared differences from the mean
	var squaredDiffsSum float64
	for _, num := range numbers {
		diff := float64(num) - mean
		squaredDiffsSum += diff * diff
	}

	// Calculate the variance as float64 and convert to int
	variance := squaredDiffsSum / float64(len(numbers))
	return int(math.Round(variance)) // Rounding to the nearest integer
}
