package main

import (
	"fmt"
)

func main() {
	// Simple Switch Case Statement

	// i := "Sunday"

	// switch i {
	// case "Sunday":
	// 	fmt.Println("This is Sunday")
	// case "Monday":
	// 	fmt.Println("This is Monday")
	// }

	// t := 103

	// switch t {
	// case 1:
	// 	fmt.Println("This is One.")
	// case 2:
	// 	fmt.Println("This is Two")
	// case 3:
	// 	fmt.Println("This is Three")
	// case 4:
	// 	fmt.Println("This is Four")
	// case 5:
	// 	fmt.Println("This is Five")
	// case 6:
	// 	fmt.Println("This is Six")
	// case 7:
	// 	fmt.Println("This is Seven")
	// case 8:
	// 	fmt.Println("This is Eight")
	// case 9:
	// 	fmt.Println("This is Nine")
	// case 10:
	// 	fmt.Println("This is Ten")
	// default:
	// 	fmt.Println("None of these")
	// }

	// Multiple Condition Switch

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("Is is Weenend")
	// default:
	// 	fmt.Println("It is workday")
	// }

	// Type Switch

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("It is an Integer")
		case string:
			fmt.Println("It is an String")
		case bool:
			fmt.Println("It is an Boolean")
		case float64:
			fmt.Println("It is an Float")
		default:
			fmt.Println("Other", t)
		}
	}
	whoAmI("Golang")
	whoAmI(534)
	whoAmI(false)
	whoAmI(54.4)

}
