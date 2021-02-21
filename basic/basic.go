package basic

import "fmt"

func Basic()  {

	// Variables can be declarate with var
	var a string

	// Variables can also declarate with shorted way :=
	b := "hello"

	// Once a variable is declared, we use regular
	b = "world" 

	fmt.Println(b)

	// If does not need round brackets like JAVA
	if a == "world" {
		fmt.Println("is a world")
	} else {
		fmt.Println("is not a world")
	}

	if a == "" {
		fmt.Println("Primetive zero values are always not nil");
	}

}