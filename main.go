package main

import (
	"fmt"

	"github.com/helvellyn-io/weather/src/api"
)

func main() {

	weather := new(api.Weather)
	api.GetJson("https://api.openweathermap.org/data/2.5/weather?q=Boulder&appid=050224087da57dabdc13099b40e747e0&units=imperial", weather)
	fmt.Printf("Location: %v Temp: %v F \n", weather.Name, weather.Main.Temp)

}
