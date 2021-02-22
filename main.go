package main

import (
  "github.com/romelgomez/golang-hello-mod/flags"
  "github.com/romelgomez/golang-hello-mod/basic"
  "github.com/romelgomez/golang-hello-mod/loops"
  "github.com/romelgomez/golang-hello-mod/arrays"
  "github.com/romelgomez/golang-hello-mod/maps"
  "fmt"
)

var Version = "development"

func main() {
  fmt.Println("Version:\t", Version)
  fmt.Println("flags.Time:\t", flags.Time)
  fmt.Println("flags.User:\t", flags.User)
  fmt.Println("flags.Date:\t", flags.Date)

  fmt.Println("\n..:: Main ::..")

  basic.Basic()
  loops.Loops()
  arrays.Arrays()
  maps.Maps()
}