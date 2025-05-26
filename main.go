package main

import "fmt"

type User struct {
	Name  string
	Age   int
	Email string
}

func main() {
	fmt.Println(getUsername())
	user, err := getUser()
	if err != nil {
		fmt.Println("Error retrieving user:", err)
		return
	}
	fmt.Println("User Name:", user)
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
