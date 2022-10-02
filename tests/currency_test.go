package tests

import (
	"github.com/IIGabriel/Upvote-crypto-currency.git/config"
	"github.com/IIGabriel/Upvote-crypto-currency.git/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMethodCurrencyCreate(t *testing.T) {
	config.Migrations()
	db := config.OpenConnection()
	defer config.CloseConnection(db)
	var coin models.Currency

	err := coin.CreateIfNotExist(db)

	require.Equal(t, 0, int(coin.Id))

	coin.Name = "ValidCoin"

	err = coin.CreateIfNotExist(db)
	require.NoError(t, err)
	require.NotEqual(t, 0, int(coin.Id))
}

func TestMethodCurrencyFindBy(t *testing.T) {
	db := config.OpenConnection()
	defer config.CloseConnection(db)

	coin := models.Currency{Name: "ValidCoin"}
	err := coin.FindBy(db)

	require.NoError(t, err)
	require.NotEqual(t, 0, int(coin.Id))
}

func TestMethodCurrencyUpdate(t *testing.T) {
	db := config.OpenConnection()
	defer config.CloseConnection(db)

	coin := models.Currency{Name: "ValidCoin"}
	err := coin.FindBy(db)

	require.NoError(t, err)
	require.NotEqual(t, 0, int(coin.Id))

	coin.Name = "SecondName"
	coin.CoinId = "TestId"

	err = coin.Update(db)

	require.NoError(t, err)

}

func TestMethodCurrencyDelete(t *testing.T) {
	db := config.OpenConnection()
	defer config.CloseConnection(db)

	coin := models.Currency{Name: "SecondName"}
	err := coin.FindBy(db)

	require.NoError(t, err)
	require.NotEqual(t, 0, int(coin.Id))

	err = coin.CreateUpVote(db)
	require.NoError(t, err)

	err = coin.CreateDownVote(db)
	require.NoError(t, err)

	err = coin.Delete(db)

	require.NoError(t, err)
}
