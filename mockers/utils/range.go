package utils

import "math/rand"

// IntRange generate a random integer in a range of min-max
func IntRange(min, max int) int {
	if min == max {
		return IntRange(min, min+1)
	}
	if min > max {
		return IntRange(max, min)
	}
	n := rand.Int() % (max - min)
	return n + min
}
