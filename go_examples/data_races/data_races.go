package data_races

// go run -race main.go

import (
	"fmt"
)

func DataRaces() {
	fmt.Println(getNumber())
}

func getNumber() int {
	var i int
	go func() {
		i = 5
	}()

	return i
}
