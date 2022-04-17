package structs

import (
	"fmt"
)

//
// ref: https://stackoverflow.com/a/50320595/2513972
//

type person struct {
	// taggin can  be added for serialization
	FirstName string `json:"firstName" yaml:"firstName"`
	lastName  string
	age       int

	// Unexported struct fields are invisible to the JSON package.
	// Export a field by starting it with an uppercase letter.
	unexported string

	// {"Exported": ""}
	Exported string

	// {"custom_name": ""}
	CustomName string `json:"custom_name"`
}

func Structs() {

	p1 := person{
		FirstName: "rome",
	}

	// Assigning values to the fields in the person struct
	fmt.Println("He us the man: ", p1)

	type animal struct {
		name            string
		characteristics []string
	}

	animal1 := animal{
		name: "Lion",
		characteristics: []string{
			"Eats plants",
			"Will animal",
		},
	}

	fmt.Println("Animal name:", animal1.name)

	for _, v := range animal1.characteristics {
		fmt.Printf("\t %v\n", v)
	}

	type herbivore struct {
		animal
		eatPotatos bool
		// sync.Mutex
	}

	herb := herbivore{
		animal: animal{
			name: "romel",
			characteristics: []string{
				"good looking",
				"smart",
			},
		},
		eatPotatos: true,
	}

	fmt.Println(herb)
}

// package main

// import (
//     "encoding/json"
//     "fmt"
//     "io/ioutil"
//     "net/http"
// )

// type Time struct {
//     time
// }
// type time struct {
//     id                    string  `json:"$id"`
//     currentDateTime       string  `json:"currentDateTime,string"`
//     utcOffset             float64 `json:"utcOffset,string"`
//     isDayLightSavingsTime bool    `json:"isDayLightSavingsTime,string"`
//     dayOfTheWeek          string  `json:"dayOfTheWeek,string"`
//     timeZoneName          string  `json:"timeZoneName,string"`
//     currentFileTime       float64 `json:"currentFileTime,string"`
//     ordinalDate           string  `json:"ordinalDate,string"`
//     serviceResponse       string  `json:"serviceResponse,string"`
// }

// func (t *Time) GetTime() (Time, error) {
//     result := Time{}

//     return result, t.Timenow(result)
// }
// func (t *Time) Timenow(result interface{}) error {

//     res, err := http.Get("http://worldclockapi.com/api/json/utc/now")
//     if err != nil {
//         fmt.Println("Cannot get Json", err)
//     }

//     body, err := ioutil.ReadAll(res.Body)
//     if err != nil {
//         fmt.Println("Cannot create Body", err)
//     }

//     defer res.Body.Close()

//     var resultJson interface{}
//     return json.Unmarshal(body, &resultJson)

// }

// func main() {

//     var a Time
//     t, err := a.GetTime()
//     if err != nil {
//         fmt.Println("Error ", err)
//     }
//     fmt.Println("Time:", t)
// }
