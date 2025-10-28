package main

import "fmt"

// for -> only construct in go for looping
func main() {
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}

	// classic for loop:
	fmt.Println("Table of 2:")
	for i = 1; i <= 10; i++ {
		if i == 6 {
			continue
		}
		fmt.Println(i * 3)
	}
	fmt.Println("----------------------------------")
	for e := 1; e <= 5; e++ {
		if e == 3 {
			continue
		}
		fmt.Println(e)
	}

	fmt.Println("-----------------Range------------------------")

	for x := range 10 {
		fmt.Println(x)
	}
}
