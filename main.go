package main

import (
	"dayone/adapter"
	"dayone/user"
	"fmt"
)

func main() {

	getUser := adapter.NewAdapter()

	user := user.NewUser(getUser)
	fmt.Println(user.GetUsername("1"))
	fmt.Println(user.GetUsername("999"))
}
