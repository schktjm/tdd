package fizzbuzz

import (
	"errors"
	"fmt"
)

var ErrorOutOfRange = errors.New("input out of range")

func FizzBuzz(n int) (string, error) {
	if n < 0 || n > 100 {
		return "", ErrorOutOfRange
	}
	if (n % 15) == 0 {
		return "FizzBuzz", nil
	} else if (n % 5) == 0 {
		return "Buzz", nil
	} else if (n % 3) == 0 {
		return "Fizz", nil
	} else {
		return fmt.Sprint(n), nil
	}
}
