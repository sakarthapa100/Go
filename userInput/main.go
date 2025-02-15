package main

import "fmt"

func main() {
	fmt.Println("hey what's your name")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Hello What's up man", name)
}
