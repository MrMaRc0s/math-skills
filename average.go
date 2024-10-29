package main

import "math"

func calculateAverage(numbers []int) int {
	total := 0.0
	for i := 0; i < len(numbers); i++ {
		total += float64(numbers[i])
	}

	// Calculate the average first
	average := total / float64(len(numbers))

	// Then round the average
	return int(math.Round(average))
}
