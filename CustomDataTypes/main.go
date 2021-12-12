package main

import (
	"errors"
	"fmt"
)

//type alias , reference to another Typ

type Number = int

//type Declaration

type Chars string

func main() {
	p := Person{firstName: "Mahmoud", lastName: "Belal"}
	fmt.Println(p.firstName, p.lastName, p.ID())

	p2 := NewPerson("Mahmoud", "Belal")
	fmt.Println(p2.firstName, p2.lastName, p2.ID())

	err := p2.SetFirstName("Ahmed")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(p2.firstName)

}

//Define Custom Data Type , that implement interface
type identifiable interface {
	ID() string
}
type Person struct {
	//lowerCase is Private, Private in Go Word means we can't access it outside the Package
	firstName string
	lastName  string
}

//Go doesn't have meaning of Constructor, so we must define Method that internalize our Struct

func NewPerson(firstName, lastName string) Person {
	return Person{
		firstName: "Mahmoud",
		lastName:  "Reda",
	}
}

//add Method to the Struct

func (p Person) ID() string {
	return "1453"
}

// here we pass point of Person Struct , So we can modify its properties

func (p *Person) SetFirstName(firstName string) error {
	fmt.Println(len(firstName))

	if len(firstName) == 0 {
		return errors.New("first Name must not be Empty")
	} else {
		p.firstName = firstName
	}

	return nil
}

// GoLang Prefer Composition not inheritance
