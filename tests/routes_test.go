package tests

import (
	"encoding/json"
	"fmt"
	"github.com/IIGabriel/Upvote-crypto-currency.git/models"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestRoteCreateCurrency(t *testing.T) {
	url := fmt.Sprintf("http://localhost:7777/currency")

	bodyReq := strings.NewReader(`{` +
		`"id": "testeId",` +
		`"name": "testeName",` +
		`"symbol": "TNM"` +
		`}`)

	req, err := http.NewRequest("POST", url, bodyReq)
	require.NoError(t, err)

	req.Header.Add("Permission_token", "ee2a9221b5msh3c607db06792088p1ef1b4jsnb991cd368659")

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	require.Equal(t, 201, res.StatusCode)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var response string
	err = json.Unmarshal(body, &response)
	require.NoError(t, err)
	require.Equal(t, "Currency added", response)
}

func TestRoteUpvote(t *testing.T) {
	url := fmt.Sprintf("http://localhost:7777/upvote/testeName")

	req, err := http.NewRequest("POST", url, nil)
	require.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	require.Equal(t, 201, res.StatusCode)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var response string
	err = json.Unmarshal(body, &response)
	require.NoError(t, err)
	require.Equal(t, "Upvote registered", response)
}

func TestRoteDownvote(t *testing.T) {
	url := fmt.Sprintf("http://localhost:7777/downvote/testeName")

	req, err := http.NewRequest("POST", url, nil)
	require.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	require.Equal(t, 201, res.StatusCode)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var response string
	err = json.Unmarshal(body, &response)
	require.NoError(t, err)
	require.Equal(t, "Downvote registered", response)
}

func TestRoteGetCurrency(t *testing.T) {
	url := fmt.Sprintf("http://localhost:7777/currency/testeName")

	req, err := http.NewRequest("GET", url, nil)
	require.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	require.Equal(t, 200, res.StatusCode)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var coin models.Currency
	err = json.Unmarshal(body, &coin)
	require.NoError(t, err)
	require.NotEqual(t, "", coin.Name)
}

func TestRoteUpdateCurrency(t *testing.T) {
	url := fmt.Sprintf("http://localhost:7777/currency/testeName")

	bodyReq := strings.NewReader(`{` +
		`"id": "testeIdUpdated",` +
		`"name": "testeNameUpdated",` +
		`"symbol": "TNMUpdated"` +
		`}`)

	req, err := http.NewRequest("PUT", url, bodyReq)
	require.NoError(t, err)

	req.Header.Add("Permission_token", "ee2a9221b5msh3c607db06792088p1ef1b4jsnb991cd368659")

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	require.Equal(t, 200, res.StatusCode)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var response string
	err = json.Unmarshal(body, &response)
	require.NoError(t, err)
	require.Equal(t, "Currency updated", response)
}

func TestRoteDeleteCurrency(t *testing.T) {
	url := fmt.Sprintf("http://localhost:7777/currency/testeNameUpdated")

	req, err := http.NewRequest("DELETE", url, nil)
	require.NoError(t, err)

	req.Header.Add("Permission_token", "ee2a9221b5msh3c607db06792088p1ef1b4jsnb991cd368659")

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	require.Equal(t, 200, res.StatusCode)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var response string
	err = json.Unmarshal(body, &response)
	require.NoError(t, err)
	require.Equal(t, "Currency deleted", response)
}
