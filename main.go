package main

import "dayone/fizzbuzz"

func main() {
	// restapi.Init()
	// fizzbuzz.FizzBuzz(100)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fizzbuzz.FizzBuzzSlice(numbers)
}
