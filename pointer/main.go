package main

import "fmt"

func main() {
	number := 10

	pointer := &number

	fmt.Println("Pointer is ", *pointer)
	fmt.Println("Value of the number is ", number)
}
