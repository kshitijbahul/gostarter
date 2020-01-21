package main

import (
	"fmt"
)

type contact struct {
	email   string
	zipcode int
}
type person struct {
	firstName string
	lastName  string
	//contactinfo contact
	contact
}

func main() {
	//alex := person{"Alex", "Andeerson"}
	//alex := person{firstName: "Alex", lastName: "Andeerson"}
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Andeerson"
	alex.contact.email = "alex.email"
	alex.contact.zipcode = 171011
	jim := person{
		firstName: "Jim",
		lastName:  "Cary",
		contact: contact{
			email:   "abc@def.com",
			zipcode: 171011,
		},
	}
	fmt.Printf("%+v \n", alex)
	//fmt.Println(jim)
	jim.print()
	//jim.updateFirstName("Jimmy")
	//Above doesn't update cause Go passes by value
	jim.print()
	//pointerToJim := &jim
	//pointerToJim.updateFirstName("Jimmy")
	jim.updateFirstName("Jimmy") //same as above since go converts value implicitly based on reciever
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
