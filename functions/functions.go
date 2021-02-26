package functions

import (
	"fmt"
)

type Employee struct {
	FirstName string
	LastName  string
}

func fullName(firstName string, lastName string) (fullName string) {
	fullName = firstName + " " + lastName
	return fullName
}

func Functions() {

	var e Employee = Employee{
		FirstName: "romel",
		LastName:  "gomez",
	}

	fmt.Println("Mi full name is ", fullName(e.FirstName, e.LastName))

	fmt.Println("yes, ", e.fullName())

	e.changeLastName("herrera")

	fmt.Println("yes, too ", e.fullName())

}

func (e Employee) fullName() string {
	return e.FirstName + " " + e.LastName
}

func (e *Employee) changeLastName(firstName string) {
	e.FirstName = firstName
}
