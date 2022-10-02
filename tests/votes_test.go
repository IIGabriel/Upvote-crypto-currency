package tests

import (
	"github.com/IIGabriel/Upvote-crypto-currency.git/config"
	"github.com/IIGabriel/Upvote-crypto-currency.git/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateUpVote(t *testing.T) {
	config.Migrations()
	db := config.OpenConnection()
	defer config.CloseConnection(db)

	var coin = models.Currency{Name: "VoteTest"}
	err := coin.CreateIfNotExist(db)
	require.NoError(t, err)

	if coin.Id == 0 {
		err = coin.FindBy(db)
		require.NoError(t, err)
	}

	for i := 0; i < 10; i++ {
		err = coin.CreateUpVote(db)
		require.NoError(t, err)
	}

}
func TestCreateDownVote(t *testing.T) {
	db := config.OpenConnection()
	defer config.CloseConnection(db)

	var coin = models.Currency{Name: "VoteTest"}
	err := coin.CreateIfNotExist(db)
	require.NoError(t, err)

	if coin.Id == 0 {
		err = coin.FindBy(db)
		require.NoError(t, err)
	}

	for i := 0; i < 10; i++ {
		err = coin.CreateDownVote(db)
		require.NoError(t, err)
	}

}

func TestGetAllVotes(t *testing.T) {
	db := config.OpenConnection()
	defer config.CloseConnection(db)

	var coin = models.Currency{Name: "VoteTest"}
	err := coin.CreateIfNotExist(db)
	require.NoError(t, err)

	if coin.Id == 0 {
		err = coin.FindBy(db)
		require.NoError(t, err)
	}

	err = coin.FindVotes(db)
	require.NoError(t, err)
	require.NotEqual(t, 0, coin.Votes.Up)
	require.NotEqual(t, 0, coin.Votes.Down)

	err = coin.Delete(db)

	require.NoError(t, err)
}
