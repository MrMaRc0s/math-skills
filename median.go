package main

import "math"

func calculateMedian(numbers []int) int {
	n := len(numbers)
	if n%2 == 1 { // If odd
		return numbers[n/2] // Return the middle element as int
	} else { // If even
		mid1 := numbers[n/2-1]
		mid2 := numbers[n/2]
		// Return the average of the two middle elements, rounded to the nearest int
		return int(math.Round(float64(mid1+mid2) / 2.0))
	}
}
