package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	// goes in order
	// not a great way to do things
	firstName string
	string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		string:    "Jimerson",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// technically, more correct
	// go will turn the struct into a pointer
	// to the struct, on it's own
	// jimPointer := &jim
	jim.updateName("Something")
	jim.print()

	mySlice := []string{"Hi", "There", "How", "Are", "You"}

	updateSlice(mySlice)

	fmt.Println(mySlice)

}

func updateSlice(s []string) {
	s[0] = "Bye"
}

func (p person) print() {
	fmt.Printf("%+v", p)

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
