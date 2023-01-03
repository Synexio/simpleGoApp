package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Coin struct {
	Prices        [][]float32 `json:"prices"`
	Market_caps   [][]float32 `json:"market_caps"`
	Total_volumes [][]float32 `json:"total_volumes"`
}

func main() {

	req, err := http.NewRequest("GET", "https://api.coingecko.com/api/v3/coins/bitcoin/market_chart?vs_currency=eur&days=0", nil)
	CheckError(err)

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	CheckError(err)

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	CheckError(err)

	var data Coin
	if err := json.Unmarshal(b, &data); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data.Prices[0][1])
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
