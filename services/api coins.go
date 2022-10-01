package services

import (
	"encoding/json"
	"fmt"
	"github.com/IIGabriel/Upvote-crypto-currency.git/config"
	"github.com/IIGabriel/Upvote-crypto-currency.git/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"time"
)

func GetPrice(coin *models.Currency) error {

	url := fmt.Sprintf("https://coingecko.p.rapidapi.com/coins/%s/market_chart/range?from=%d&vs_currency=%s&to=%d",
		coin.CoinId, time.Now().AddDate(-1, 0, 0).Unix(), config.GetEnv("main_currency"), time.Now().Unix())

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("X-RapidAPI-Key", config.GetEnv("coingecko_token"))
	req.Header.Add("X-RapidAPI-Host", "coingecko.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		zap.L().Warn("Error Currency - GetPrice(): HTTP status != 200")
		return nil
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
		zap.L().Info("Error GetAllCoins():", zap.Error(err))
		return err
	}

	req.Header.Add("X-RapidAPI-Key", config.GetEnv("coingecko_token"))
	req.Header.Add("X-RapidAPI-Host", "coingecko.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		zap.L().Info("Error GetAllCoins():", zap.Error(err))
		return err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		zap.L().Info("Error GetAllCoins(): HTTP status != 200")
		return nil
	}
	body, _ := ioutil.ReadAll(res.Body)
	if err = json.Unmarshal(body, &allCoins); err != nil {
		zap.L().Info("Error GetAllCoins():", zap.Error(err))
		return err
	}

	for _, item := range allCoins {
		if err = item.CreateIfNotExist(db); err != nil {
			zap.L().Info("Error creating currencies in db GetAllCoins():", zap.Error(err))
			return err
		}
	}
	return nil
}
