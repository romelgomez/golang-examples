package main

import (
	"github.com/romelgomez/golang-hello-mod/basic"
	"github.com/romelgomez/golang-hello-mod/flags"
	"fmt"
)

var Version = "development"

func main() {
    fmt.Println("Version:\t", Version)
    fmt.Println("flags.Time:\t", flags.Time)
    fmt.Println("flags.User:\t", flags.User)
    fmt.Println("flags.Date:\t", flags.Date)

	fmt.Println("Hello, 世界")

	basic.Basic()
}