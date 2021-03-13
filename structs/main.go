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

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 123435,
		},
	}

	//&variable --> memory address of the value of this variable; turn value into memory address
	//*pointer --> value this memory address is pointing at; turn memory address into value

	//jimPointer := &jim
	//jimPointer.updateName("Jimmy")

	jim.updateName("Jimmy")
	jim.print()
}

//*person --> type of the variable that precedes it
//*pointerToPerson --> the operator, the value of type person

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
