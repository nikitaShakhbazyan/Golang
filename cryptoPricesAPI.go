package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func fetchPrice(currency string) string {
	var data map[string]float64

	url := fmt.Sprintf("https://min-api.cryptocompare.com/data/price?fsym=%s&tsyms=USD", currency)
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

	return fmt.Sprintf("The price of %s is %.2f USD", currency, data["USD"])
}

func main() {
	startNow := time.Now()
	currencies := []string{"BTC", "XRP"}

	for _, currency := range currencies {
		data := fetchPrice(currency)
		fmt.Println(data)
	}
	fmt.Println("The operation took:", time.Since(startNow))
}
