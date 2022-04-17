package pointers

import "fmt"

// ..:: Struct example

type PersonStruct struct {
	firstName string
	lastName  string
}

func changeNameVal(p PersonStruct) {
	p.firstName = "unchange"
}

func changeName(p *PersonStruct) {
	p.firstName = "frisbee"
}

// Array example

func changeSecondItemVal(arr [7]int) {
	arr[1] = 1
}

func changeSecondItem(arr *[7]int) {
	arr[1] = 1
}

// slice example

func changeSecondItemSlice(arr []int) {
	arr[1] = 1
}

func Pointers() {

	// ..:: Struct

	PersonStruct := PersonStruct{
		firstName: "romel",
		lastName:  "gomez",
	}

	fmt.Println("PersonStruct,", PersonStruct)

	changeNameVal(PersonStruct)

	fmt.Println("PersonStruct did not change,", PersonStruct)

	changeName(&PersonStruct)

	fmt.Println("PersonStruct did change,", PersonStruct)

	// ..:: Arrays

	arr := [7]int{}

	fmt.Println("Array example", arr)

	changeSecondItemVal(arr)

	fmt.Println("Array example did not change", arr)

	changeSecondItem(&arr)

	fmt.Println("Array example did change", arr)

	// ..:: slices example

	var slice []int = []int{0, 0, 0, 0}

	fmt.Println("Slice example", slice)

	changeSecondItemSlice(slice)

	fmt.Println("Slice did change", slice)

	// slices, maps, channels - always sent by ref :)

}
