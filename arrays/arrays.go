package arrays

import (
	"fmt"
)

func Arrays()  {
	
	fmt.Println("\n ..:: Arrays ::..")
	
	// Initialize and declare
	a := [5]int{10, 20, 30, 40, 50}
	a[2] = 25
	fmt.Println("a", a) 

	// Declare, then initialize
	var b [5]int
	b[0] = 10
	b[1] = 20
	b[2] = 30
	b[3] = 40
	b[4] = 50
	fmt.Println("b:", b)

	fmt.Println("\n ..:: Slice - dinamic size array ::..")
	var c []int
	fmt.Println("c empty:", c)
	c = []int{10,20,30,40}
	fmt.Println("c not empty:", c)

	d := []int{10, 20, 30, 40, 50}
	d[2] = 25
	fmt.Println("d", d) 

	e := make([]int, 4)
	e = append(e, 20, 40)
	fmt.Println("e", e) 

	e[2] = e[len(e)-1]
	fmt.Println("e", e)
	fmt.Println("e(3-5)", e[2:5])

	for k, v := range e {
		fmt.Printf("%d is %d", k, v)
		fmt.Println("")
	}


}