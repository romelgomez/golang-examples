package main

import (
	"fmt"
)

func DeepSum(lst []interface{}) int {

	var acc = 0

	for i, v := range lst {

		var current interface{} = v

		intValue, isInt := current.(int)
		mapValue, isMap := current.(map[string]interface{})
		sliceValue, isSlice := current.([]interface{})

		if isInt {
			acc += intValue
		}

		if isMap {
			// converts it to a slice

			e := make([]interface{}, 0)
			e = append(e, acc)

			for _, value := range mapValue {
				e = append(e, value)
			}

			// rest of the data
			for _, value := range lst[i+1:] {
				e = append(e, value)
			}

			return DeepSum(e)
		}

		if isSlice {

			e := make([]interface{}, 0)
			e = append(e, acc)

			for _, value := range sliceValue {
				e = append(e, value)
			}

			// rest of the data
			for _, value := range lst[i+1:] {
				e = append(e, value)
			}

			return DeepSum(e)
		}

	}

	return acc
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	/*
	 * Code written here or in other files will be ignored by the automated tests.
	 * You can run this code when running the file with `go run deep-sum.go`
	 */

}
