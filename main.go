package main

import "net/http"

type User struct {
	Name  string
	Age   int
	Email string
}

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		username := getUsername()
		user, err := getUser()
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello " + username + "! User: " + user.Name + ", Age: " + string(user.Age)))
	})

	http.ListenAndServe(":8080", nil)
	// curl -X GET http://localhost:8080/hello
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
