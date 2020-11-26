package calc

import (
	"errors"
)

// Add - to calculate sum of two integers
func Add(numbers ...int) (int, error) {
	sum := 0
	if len(numbers) < 2 {
		return sum, errors.New("Minimum two arguments required for add function")
	} else {
		for _, num := range numbers {
			sum = sum + num
		}
	}
	return sum, nil
}
