package gotype

import "fmt"

type User struct {
	Name string
	Age  int
}

func GoAnyType() {
	var a int = 42
	var b float64 = 3.1415
	var c string = "Gopher"
	var d bool = true

	fmt.Println("int:", a)
	fmt.Println("float64:", b)
	fmt.Println("string:", c)
	fmt.Println("bool:", d)

	// Struct
	u := User{Name: "Alice", Age: 30}
	fmt.Println("struct:", u)
	fmt.Println("struct field (Name):", u.Name)

	// Slice
	nums := []int{1, 2, 3, 4}
	fmt.Println("slice:", nums)

	// Map
	grades := map[string]int{
		"Alice": 90,
		"Bob":   85,
	}
	fmt.Println("map:", grades)
	fmt.Println("map value (Bob):", grades["Bob"])

	// Interface (any type)
	var x interface{}
	x = "can be string"
	fmt.Println("interface{} (string):", x)
	x = 123
	fmt.Println("interface{} (int):", x)
	x = u
	fmt.Println("interface{} (struct):", x)

}

func getUser() User {
	return User{Name: "Bob", Age: 25}
}

func GetUser() User {
	return getUser()
}
