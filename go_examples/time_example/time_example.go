package time_example

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Time struct {
	time
}
type time struct {
	Id                    string  `json:"$id"`
	CurrentDateTime       string  `json:"currentDateTime,string"`
	UtcOffset             float64 `json:"utcOffset,string"`
	IsDayLightSavingsTime bool    `json:"isDayLightSavingsTime,string"`
	DayOfTheWeek          string  `json:"dayOfTheWeek,string"`
	TimeZoneName          string  `json:"timeZoneName,string"`
	CurrentFileTime       float64 `json:"currentFileTime,string"`
	OrdinalDate           string  `json:"ordinalDate,string"`
	ServiceResponse       string  `json:"serviceResponse,string"`
}

func (t *Time) GetTime() (Time, error) {
	result := Time{}

	return result, t.Timenow(result)
}
func (t *Time) Timenow(result interface{}) error {

	res, err := http.Get("http://worldclockapi.com/api/json/utc/now")
	if err != nil {
		fmt.Println("Cannot get Json", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Cannot create Body", err)
	}

	defer res.Body.Close()

	var resultJson interface{}
	return json.Unmarshal(body, &resultJson)

}

func TimeExample() {

	var a Time

	t, err := a.GetTime()

	if err != nil {
		fmt.Println("Error ", err)
	}

	fmt.Println("Time:", t)

}
