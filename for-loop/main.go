package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		println(i)
	}
	counter := 0
	for {
		fmt.Println("Infinite loop")
		counter++
		if counter == 5 {
			break
		}
	}

	numbers := []int{3, 4, 5, 6, 7}
	for index, number := range numbers {
		fmt.Println("Index:", index, "Number:", number)
	}

	name := "Hello world"
	for index, letter := range name {
		fmt.Println("index of the Data is %d, and value is %c\n", index, letter)
	}
}
