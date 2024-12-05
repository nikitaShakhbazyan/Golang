package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func fetchWeather(city string) string {
	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=10cdd5e70108c087661ce63737780047", city)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return ""
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		return ""
	}

	return fmt.Sprintf("The temperature in %s is %.2f°C", city, data.Main.Temp)
}

func main() {
	startNow := time.Now()
	cities := []string{"Tbilisi","London","Paris","Tokyo"}

	for _,city := range cities {
		data := fetchWeather(city)
		fmt.Println("This is data",data)
	}
	fmt.Println("The operation took:",time.Since(startNow))
}
