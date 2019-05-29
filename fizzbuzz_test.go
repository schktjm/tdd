package fizzbuzz

import (
	"fmt"
	"testing"
)

func TestFizz(t *testing.T) {
	table := []int{
		3, 6,
	}

	for _, n := range table {
		t.Run(fmt.Sprintf("%dの時はFizz", n), func(t *testing.T) {
			expect := "Fizz"
			actual, _ := FizzBuzz(n)

			if actual != expect {
				t.Fatal("Not Fizz")
			}
		})
	}
}

func TestBuzz(t *testing.T) {

	table := []int{
		5, 10,
	}

	for _, n := range table {
		t.Run(fmt.Sprintf("%dの時はBuzz", n), func(t *testing.T) {
			expect := "Buzz"
			actual, _ := FizzBuzz(n)

			if actual != expect {
				t.Fatal("Not Buzz")
			}
		})
	}
}

func TestFizzBuzz(t *testing.T) {
	table := []int{
		15, 30,
	}

	for _, v := range table {
		t.Run(fmt.Sprintf("%dの時はFizzBuzz", v), func(t *testing.T) {
			expect := "FizzBuzz"
			actual, _ := FizzBuzz(v)

			if actual != expect {
				t.Fatal("Not FizzBuzz")
			}
		})
	}

}

func TestNum(t *testing.T) {
	table := []int{
		1, 2, 4, 7, 8, 11,
	}
	for _, v := range table {
		t.Run(fmt.Sprintf("%dの時は1", v), func(t *testing.T) {
			expect := fmt.Sprint(v)
			actual, _ := FizzBuzz(v)

			if actual != expect {
				t.Fatalf("Not %d", v)
			}
		})
	}
}

func TestError(t *testing.T) {
	// range  0 < n < 100
	table := []int{
		-1,
		101,
	}

	for _, n := range table {
		t.Run(fmt.Sprintf("%dの時はError", n), func(t *testing.T) {
			expect := ErrorOutOfRange
			_, err := FizzBuzz(n)

			if err != expect {
				t.Fatal("Not out of range")
			}
		})
	}
}

func TestNil(t *testing.T) {
	for i := 0; i < 101; i++ {
		t.Run(fmt.Sprintf("%dの時はErrorがnil", i), func(t *testing.T) {
			_, err := FizzBuzz(i)

			if err != nil {
				t.Fatal("Not nil")
			}
		})
	}
}
