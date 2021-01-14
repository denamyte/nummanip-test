package calc

import "errors"

func Add(numbers ...int) (error, int) {

	if len(numbers) < 2 {
		return errors.New("provide more than 1 number"), 0
	}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return nil, sum
}
