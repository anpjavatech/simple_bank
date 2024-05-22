package db

import (
	"context"
	"testing"

	"github.com/anpjavatech/simple_bank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAnEntry(t *testing.T) {

	account := createRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.NotEmpty(t, account)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

}
