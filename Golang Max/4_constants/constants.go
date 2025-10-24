package main

import "fmt"

const age2 = 22

func main() {
	const name string = "golang"
	const age = 23

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Age2:", age2)

	const (
		port = 5555
		host = "localhost"
	)
	fmt.Println(port, host)
}
