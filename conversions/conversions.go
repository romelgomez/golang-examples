package conversions

import (
	"fmt"
	"strings"
)

func Conversions() {

	example1("romel")

	example1(07)

	example2(true)

	example2("hello world 2")

}

func example1(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("I store string", strings.ToUpper(i.(string))) // need to cast
	case int:
		fmt.Println("i stored int", i)
	default:
		fmt.Println("i stored something else", i)
	}
}

func example2(i interface{}) {
	switch i := i.(type) {
	case string:
		fmt.Println("I stored string", strings.ToUpper(i))
	case int:
		fmt.Println("i stored int", i)
	default:
		fmt.Println("i stored something else", i)
	}
}
