package main

import "fmt"

func main() {
	fmt.Println("Slice Learning")
	number := []int{3, 4, 5, 6, 7}
	number = append(number, 8, 9, 11, 12)
	fmt.Println(number)
	fmt.Printf("Number has the data type : %T\n", number)
}
