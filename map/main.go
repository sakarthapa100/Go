package main

import "fmt"

func main() {
	studentGrades := make(map[string]int)

	studentGrades["John"] = 92
	studentGrades["Alice"] = 78
	studentGrades["Bob"] = 68

	fmt.Println("Marks of the Alice is ", studentGrades["Alice"])
}
