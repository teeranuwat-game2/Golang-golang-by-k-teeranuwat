package fizzbuzz

import "strconv"

func FizzBuzz(n int) []string {
	result := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			result[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			result[i-1] = "Fizz"
		} else if i%5 == 0 {
			result[i-1] = "Buzz"
		} else {
			s := strconv.Itoa(i)
			result[i-1] = string(s)
		}
	}
	return result
}

func FizzBuzzSlice(numbers []int) {
	for _, i := range numbers {
		if i%3 == 0 && i%5 == 0 {
			println("FizzBuzz")
		} else if i%3 == 0 {
			println("Fizz")
		} else if i%5 == 0 {
			println("Buzz")
		} else {
			println(i)
		}
	}
}
