package loops

import (
	"fmt"
)

func Loops()  {
	fmt.Println("\n ..:: loops ::..")
	
	// Simple for loop
	sum := 0
	for i := 1; i<5; i++ {
		sum += i
	}
	fmt.Println(sum)

	// Simple while loop
	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Println(n)

	// Infinte loop
	// sum:=0
	// for {
	// 	sum++
	// 	// To stoped break or return
	// 	// break
	// 	// return
	// }

}