package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}
type Contact struct {
	Phone string
	email string
}

type Address struct {
	City  string
	State string
	Area  string
}

type Employee struct {
	Person_Details  Person
	Contact_Details Contact
	Address_Details Address
}

func main() {
	var sakar Person
	fmt.Println(sakar)
	sakar.FirstName = "Sakar"
	sakar.LastName = "Pandey"
	sakar.Age = 21
	fmt.Println("The detials fo the Person is :", sakar)

	person1 := Person{
		FirstName: "Pandey",
		LastName:  "Sakar",
		Age:       21,
	}
	fmt.Println("The details of the person1 is :", person1)

	var employee1 = Employee{
		Person_Details: Person{
			FirstName: "ramu",
			LastName:  "grg",
			Age:       11,
		},
		Contact_Details: Contact{
			Phone: "9841234567",
			email: "aalu@gmai.com",
		},
		Address_Details: Address{
			City:  "Kathmandu",
			State: "Bagmati",
			Area:  "Boudha",
		},
	}

	fmt.Println("The details of the employee1 is :", employee1)
	fmt.Println("The email of the employee is :", employee1.Contact_Details.email)
}
