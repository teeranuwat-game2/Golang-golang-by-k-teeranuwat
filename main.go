package main

import (
	"fmt"
)

type User struct {
	Name  string
	Age   int
	Email string
}

func main() {
	fmt.Println(getUsername())
	_, err := getUser()
	if err != nil {
		fmt.Println("Error retrieving user:", err)
		return
	}
	// fmt.Println("User Name:", user)

	// loop
	numbers := []int{1, 2, 3, 4, 5}
	for i, num := range numbers {
		fmt.Printf("Number %d: %d\n", i, num)
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Number %d: %d\n", i, numbers[i])
	}

}

func getUsername() string {
	return "Alice"
}

func getUser() (User, error) {
	// Simulating a user retrieval
	return User{
		Name:  "Bob",
		Age:   30,
		Email: "",
	}, nil
}
