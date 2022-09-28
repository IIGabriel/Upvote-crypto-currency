package services

import (
	"encoding/json"
	"fmt"
	"github.com/IIGabriel/Upvote-crypto-currency.git/models"
	"io/ioutil"
	"net/http"
	"time"
)

func GetPrice(coin *models.Currency) {

	url := fmt.Sprintf("https://coingecko.p.rapidapi.com/coins/%s/market_chart/range?from=%d&vs_currency=BRL&to=%d",
		coin.Name, time.Now().AddDate(-1, 0, 0).Unix(), time.Now().Unix())

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "ee2a9221b5msh3c607db06792088p1ef1b4jsnb991cd368659")
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
