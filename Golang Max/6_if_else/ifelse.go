package main

import "fmt"

func main() {
	// age := 17

	// if age >= 18 {
	// 	fmt.Println("Person is Adult")
	// } else {
	// 	fmt.Println("Person is not adult")
	// }

	// age := 11

	// if age >= 18 {
	// 	fmt.Println("Person is an adult")
	// } else if age >= 12 {
	// 	fmt.Println("Person is an teenager")
	// } else {
	// 	fmt.Println("Person is an kid")
	// }

	// var role = "admin"
	// var hasPermissions = false

	// if role == "admin" || hasPermissions {
	// 	fmt.Println("YES!")
	// }

	// var role2 = "admin"
	// var hasPermissions1 = false

	// if role2 == "admin" && hasPermissions1 {
	// 	fmt.Println("YES!")
	// } else {
	// 	fmt.Println("No")
	// }

	if age := 45; age >= 29 {
		fmt.Println("Person is an adult:", age)
	} else if age >= 11 {
		fmt.Println("Person is an kid:", age)
	} else {
		fmt.Println("Person is unknown:", age)
	}

}
