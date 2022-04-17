package maps

import (
  "fmt"
)

var sampleMap map[string]int

func Maps() {

  fmt.Printf("\n ..:: maps ::..")

  sampleMap = map[string]int{
    "monday": 0,
    "tuesday": 1, 
  }

  currency := map[string]string{
    "AUD":"Austria Dollar",
    "GBP":"Great Britain Pound",
    "JPY":"Japan Yen",
    "CHF":"Switzerland Franc",
  }

  // Addind to the map:
  currency["USD"] = "USA Dollar"

  fmt.Println("Currency with USD added:", currency)

  // Remove from the map
  delete(currency, "GBP")
  fmt.Println("Currency with key GBP was deleted", currency)

  // updating the value
  currency["AUD"] = "Austria Dollar 2"
  fmt.Println("Currency with key AUD was udpated", currency)

  for k, v := range currency {
    fmt.Printf("%v might be equal to %v\n", k, v)
  }

}