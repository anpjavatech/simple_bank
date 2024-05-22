package db

import (
	"context"
	"testing"

	"github.com/anpjavatech/simple_bank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateTransfer(t *testing.T) {

	from_account := createRandomAccount(t)
	to_account := createRandomAccount(t)
	arg := CreateTransferParams{
		FromAccountID: from_account.ID,
		ToAccountID:   to_account.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
}
