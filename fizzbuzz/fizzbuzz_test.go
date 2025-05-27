package fizzbuzz

import "testing"

func TestFizzbuzz(t *testing.T) {

	t.Run("Test FizzBuzz with 15", func(t *testing.T) {
		expected := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
		result := FizzBuzz(15)
		for i, v := range result {
			if v != expected[i] {
				t.Errorf("Expected %s but got %s at index %d", expected[i], v, i)
			}
		}
	})
}
