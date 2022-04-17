package main

import (
	"fmt"

	"github.com/romelgomez/golang-hello-mod/go_examples/flags"
	todos "github.com/romelgomez/golang-hello-mod/go_examples/todos"
)

var Version = "0.0.0"

func main() {
	fmt.Println("Version:\t", Version)
	fmt.Println("Environment:\t", flags.Environment)
	fmt.Println("Time:\t", flags.Time)
	fmt.Println("User:\t", flags.User)
	fmt.Println("Date:\t", flags.Date)

	fmt.Println("\n..:: starting the server ...")
	todos.Server()
}
