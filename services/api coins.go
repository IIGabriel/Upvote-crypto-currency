package services

import (
	"encoding/json"
	"fmt"
	"github.com/IIGabriel/Upvote-crypto-currency.git/models"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"time"
)

func GetPrice(coin *models.Currency) {

	url := fmt.Sprintf("https://coingecko.p.rapidapi.com/coins/%s/market_chart/range?from=%d&vs_currency=BRL&to=%d",
		coin.CoinId, time.Now().AddDate(-1, 0, 0).Unix(), time.Now().Unix())

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "")
	req.Header.Add("X-RapidAPI-Host", "coingecko.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	if res.StatusCode != 200 {
		fmt.Println("Erro")
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	if err := json.Unmarshal(body, coin); err != nil {
		fmt.Println("Erro")
	}

}

func GetAllCoins(db *gorm.DB) {
	var allCoins []models.Currency
	url := "https://coingecko.p.rapidapi.com/coins/list"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "")
	req.Header.Add("X-RapidAPI-Host", "coingecko.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Println("Erro")
	}
	body, _ := ioutil.ReadAll(res.Body)
	if err := json.Unmarshal(body, &allCoins); err != nil {
		fmt.Println("Erro")
	}
	for _, item := range allCoins {
		if err := item.Create(db); err != nil {
			fmt.Println("Erro")
		}
	}

}
