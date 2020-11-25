package calc

import (
	"errors"
)

// Add - to calculate sum of two integers
func Add(numbers ...int)(error, int)  {
	sum := 0;
	if len(numbers) < 2){
		return errors.New("Minimum two arguments required for add function"), sum
	} else {
		for _, num := range numbers {
			sum = sum + num
		}
	}	
	return nil, sum
}