package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/IIGabriel/Upvote-crypto-currency.git/models"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"time"
)

func GetPrice(coin *models.Currency) error {

	url := fmt.Sprintf("https://coingecko.p.rapidapi.com/coins/%s/market_chart/range?from=%d&vs_currency=BRL&to=%d",
		coin.CoinId, time.Now().AddDate(-1, 0, 0).Unix(), time.Now().Unix())

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("X-RapidAPI-Key", "ee2a9221b5msh3c607db06792088p1ef1b4jsnb991cd368659")
	req.Header.Add("X-RapidAPI-Host", "coingecko.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	if res.StatusCode != 200 {
		fmt.Println("Erro")
		return errors.New("Error while getting currency price")
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	ReceivePrice := struct {
		Prices [][]float64 `json:"prices"`
	}{}

	if err = json.Unmarshal(body, &ReceivePrice); err != nil {
		fmt.Println("Erro")
		return err
	}
	var CoinPrice models.Price
	for _, item := range ReceivePrice.Prices {
		CoinPrice.Price = item[1]
		CoinPrice.Date = time.Unix(int64(item[0]/1000), 0).Format("02/01/2006")
		coin.Prices = append(coin.Prices, CoinPrice)
	}

	return nil
}

func GetAllCoins(db *gorm.DB) error {
	var allCoins []models.Currency
	url := "https://coingecko.p.rapidapi.com/coins/list"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Erro")
		return err
	}

	req.Header.Add("X-RapidAPI-Key", "ee2a9221b5msh3c607db06792088p1ef1b4jsnb991cd368659")
	req.Header.Add("X-RapidAPI-Host", "coingecko.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Erro")
		return err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Println("Erro")
		return errors.New("Erro")
	}
	body, _ := ioutil.ReadAll(res.Body)
	if err = json.Unmarshal(body, &allCoins); err != nil {
		fmt.Println("Erro")
		return err
	}

	for _, item := range allCoins {
		if err = item.CreateIfNotExist(db); err != nil {
			fmt.Println("Erro")
			return err
		}
	}
	return nil
}
