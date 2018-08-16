package main

import (
	"fmt"
)

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// person1 := person{"Alex", "Anderson"}
	// person1 := person{firstName: "Alex", lastName: "Anderson"}

	// var person1 person
	// //Assigns "Zero value to properties"
	// person1.firstName = "Karan"
	// fmt.Println(person1)
	// fmt.Printf("%+v", person1)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email: "odjoeo",
			zip:   12355,
		},
	}

	(&jim).updateName("Jimmy") // Go even allows jim.updateName("Jimmy") and works same
	jim.print()

}

func (ptr *person) updateName(name string) {
	(*ptr).firstName = name
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}

// Notes :

// package main

// import "fmt"

// func main() {
//  name := "bill"

//  namePointer := &name

//  fmt.Println(&namePointer)
//  printPointer(namePointer)
// }

// func printPointer(namePointer *string) {
//  fmt.Println(&namePointer)
// }

// This prints different values
