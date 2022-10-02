package services

import (
	"encoding/json"
	"fmt"
	"github.com/IIGabriel/Upvote-crypto-currency.git/config"
	"github.com/IIGabriel/Upvote-crypto-currency.git/models"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"time"
)

func GetPrice(coin *models.Currency, c *fiber.Ctx) error {

	// Getting url params
	mainCurrency := c.Query("compare_currency")
	if len(mainCurrency) == 0 {
		mainCurrency = config.GetEnv("main_currency")
	}
	finalDate, _ := time.Parse("2006-01-02", c.Query("final_date"))
	if finalDate.IsZero() {
		finalDate = time.Now()
	}
	initialDate, _ := time.Parse("2006-01-02", c.Query("initial_date"))
	if initialDate.IsZero() {
		initialDate = finalDate.AddDate(-1, 0, 0)
	}

	url := fmt.Sprintf("https://coingecko.p.rapidapi.com/coins/%s/market_chart/range?from=%d&vs_currency=%s&to=%d",
		coin.CoinId, initialDate.Unix(), mainCurrency, finalDate.Unix())

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
		zap.L().Info("Error Currency - GetPrice():", zap.Int("HTTP status", res.StatusCode))
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
		CoinPrice.Date = time.Unix(int64(item[0]/1000), 0).Format("2006/01/02 15:01:05")
		coin.Prices = append(coin.Prices, CoinPrice)
	}

	return nil
}

func GetAllCoins() {
	var allCoins []models.Currency
	url := "https://coingecko.p.rapidapi.com/coins/list"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		zap.L().Warn("Error GetAllCoins():", zap.Error(err))
	}

	req.Header.Add("X-RapidAPI-Key", config.GetEnv("coingecko_token"))
	req.Header.Add("X-RapidAPI-Host", "coingecko.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		zap.L().Warn("Error GetAllCoins():", zap.Error(err))
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		zap.L().Warn("Error GetAllCoins(): HTTP status != 200")
	}
	body, _ := ioutil.ReadAll(res.Body)
	if err = json.Unmarshal(body, &allCoins); err != nil {
		zap.L().Info("Error GetAllCoins():", zap.Error(err))
	}
	db := config.OpenConnection()
	defer config.CloseConnection(db)

	for _, item := range allCoins {
		if err = item.CreateIfNotExist(db); err != nil {
			zap.L().Warn("Error creating currencies in db GetAllCoins():", zap.Error(err))
		}
	}
}
