package main

import "fmt"

func divide(x, y float64) (float64, string) {
	if y == 0 {
		return 0, "cannot divide by zero"
	}
	return x / y, "nil"
}

func main() {
	fmt.Println("Hello World")
	ans, _ := divide(10, 0)
	// if error != nil {
	// 	fmt.Println("Error occured", error)
	// 	return
	// }
	fmt.Println("The division of the two number is ", ans)

}
