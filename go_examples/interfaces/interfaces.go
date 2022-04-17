package interfaces

import "fmt"

type InterfaceName interface {
	Area() float64
	Perimeter() float64
}

type StructName struct {
	width  float64
	height float64
}

func (r StructName) Area() float64 {
	return r.height * r.width
}

func (r StructName) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

var s InterfaceName

func Interfaces() {

	s = StructName{width: 5.0, height: 4.0}
	r := StructName{5.0, 4.0}

	fmt.Printf("type of s is %T\n", s)
	fmt.Printf("value of s is %v\n", s)
	fmt.Println("area of StructNameagle", s.Area())
	fmt.Println("s === r ", s == r)

}

func printArea(s InterfaceName) {
	fmt.Printf("Area of shape is: %v", s.Area())
}
