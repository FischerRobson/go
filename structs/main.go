package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newName string) {
	(*pointerToPerson).firstName = newName
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Doe",
		contactInfo: contactInfo{
			email:   "jimdoe@mail.com",
			zipCode: 12345,
		},
	}

	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")
	jim.updateName("jimmy")
	jim.print()
}
