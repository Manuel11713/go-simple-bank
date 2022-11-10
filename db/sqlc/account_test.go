package db

import (
	"context"
	"testing"

	"github.com/Manuel11713/simple-bank/utils"
	"github.com/stretchr/testify/require"
)

var testAccount Account

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := queries.CreateAccount(context.Background(), arg)

	require.ErrorIs(t, err, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	testAccount = account

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {

	account, err := queries.GetAccount(context.Background(), testAccount.ID)

	require.ErrorIs(t, err, err)
	require.NotEmpty(t, account)

	require.Equal(t, testAccount.Owner, account.Owner)
	require.Equal(t, testAccount.Balance, account.Balance)
	require.Equal(t, testAccount.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestUpdateAccount(t *testing.T) {
	newMoney := utils.RandomMoney()

	params := UpdateAccountParams{
		ID:      testAccount.ID,
		Balance: newMoney,
	}

	account, err := queries.UpdateAccount(context.Background(), params)

	require.ErrorIs(t, err, err)
	require.NotEmpty(t, account)

	require.Equal(t, testAccount.Owner, account.Owner)
	require.Equal(t, newMoney, account.Balance)
	require.Equal(t, testAccount.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

}

func TestDeleteAccount(t *testing.T) {
	err := queries.DeleteAccount(context.Background(), testAccount.ID)
	require.ErrorIs(t, err, err)

	// Code from TestGetAccount

	account, err := queries.GetAccount(context.Background(), testAccount.ID)

	require.ErrorIs(t, err, err)
	require.Empty(t, account)

}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	args := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := queries.ListAccounts(context.Background(), args)

	require.ErrorIs(t, err, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
