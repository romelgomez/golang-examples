package main

import "fmt"

func FinalValue(movs []int) int {

	// max steps of the safe konb
	var stepsListLen int = 100

	konbStepsList := make([]int, stepsListLen)

	for i := range konbStepsList {
		konbStepsList[i] = i
	}

	var konbStartAt int = 0

	for i := range movs {
		var currentMove = movs[i]

		var stepTo int = konbStartAt + (currentMove % stepsListLen)

		var step = 0

		if stepTo >= 0 {
			step = stepTo % stepsListLen
		} else {
			step = (stepTo % stepsListLen) + stepsListLen
		}

		konbStartAt = konbStepsList[step]

	}

	return konbStartAt

}

func main() {
	/*
	 * Code written here or in other files will be ignored by the automated tests
	 * You can run this code when running the file with `go run safe.go`
	 */

	var test_0 = 100 % 100
	var test_1 = -1 % 100
	var test_2 = 0 % 100
	var test_3 = (-1 % 100) + 100

	fmt.Println("test_0", test_0)
	fmt.Println("test_1", test_1)
	fmt.Println("test_2", test_2)
	fmt.Println("test_3", test_3)
}
