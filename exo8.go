package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Coin struct {
	Current_price         float32                `json:"current_price"`
	Market_cap            float32                `json:"market_cap"`
	Market_cap_change_24h float32                `json:"market_cap_change_24h"`
	High_24h              float32                `json:"high_24h"`
	Low_24h               float32                `json:"low_24h"`
	Price_change_24h      float32                `json:"price_change_24h"`
	X                     map[string]interface{} `json:"-"`
}

func main() {

	CoinGeckoAPI := "https://api.coingecko.com/api/v3/"
	req, err := http.NewRequest("GET", CoinGeckoAPI+"coins/markets?vs_currency=usd&ids=bitcoin%2Cethereum%2Cbinancecoin%2Ccardano%2Ctezos&order=market_cap_desc&price_change_percentage=24h", nil)
	CheckError(err)

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	CheckError(err)

	defer resp.Body.Close()

	s, err := io.ReadAll(resp.Body)
	CheckError(err)
	fmt.Println(string(s))

	var f []Coin
	if err := json.Unmarshal(s, &f); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(s, &f[0].X); err != nil {
		panic(err)
	}
	delete(f[0].X, "current_price")
	delete(f[0].X, "market_cap")
	delete(f[0].X, "market_cap_change_24h")
	delete(f[0].X, "high_24h")
	delete(f[0].X, "low_24h")
	delete(f[0].X, "price_change_24h")

	fmt.Printf("%+v", f)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
