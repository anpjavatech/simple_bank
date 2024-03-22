package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/anpjavatech/simple_bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Owner, account.Owner)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	fetchedAccount, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, fetchedAccount)
	require.Equal(t, account1.ID, fetchedAccount.ID)
	require.Equal(t, account1.Balance, fetchedAccount.Balance)
	require.Equal(t, account1.Currency, fetchedAccount.Currency)
	require.WithinDuration(t, account1.CreatedAt, fetchedAccount.CreatedAt, 1*time.Second)
}

func TestUpdateAccount(t *testing.T) {
	previousAccount := createRandomAccount(t)
	arg := UpdateAccountBalanceAndReturnParams{
		ID:      previousAccount.ID,
		Balance: util.RandomMoney(),
	}

	updatedAccount, err := testQueries.UpdateAccountBalanceAndReturn(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)
	require.Equal(t, arg.Balance, updatedAccount.Balance)
	require.Equal(t, previousAccount.ID, updatedAccount.ID)
	require.Equal(t, previousAccount.Currency, updatedAccount.Currency)
	require.WithinDuration(t, previousAccount.CreatedAt, updatedAccount.CreatedAt, 1*time.Second)

}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	accountDeleted, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, accountDeleted)

}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
