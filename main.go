package main

import (
	"fmt"
)

func main() {
	fmt.Println("Learn Go lang b y hellow world")
	fmt.Println("------------")

	var name string = "ramu"
	var version = "15.5"
	fmt.Println(name)
	fmt.Println(version)

	var dimension float32 = 22.22
	fmt.Println(dimension)

	var person = "ramu thapa"
	fmt.Println(person)

	var Public = "this is the public data  "
	// the main difference between the Public and private is the Public starts with the capital lettter so that it can be export on other file where as the private is unable to export
	var private = " this is the private data "

	fmt.Println(Public)
	fmt.Println(private)

	PublicFunction()
	privateFunction()

}

func PublicFunction() {
	fmt.Println("this is the public function ")

}
func privateFunction() {
	fmt.Println("this is the private function ")
}
