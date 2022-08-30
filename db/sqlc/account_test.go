package db

import (
	"context"
	"testing"

	"github.com/jaeyoung0509/go-banking/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)

	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}
func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)

}

func TestGetAccount(t *testing.T) {
	//create account
	accountInDB := CreateRandomAccount(t)
	account, err := testQueries.GetAccount(context.Background(), accountInDB.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, accountInDB.Owner, account.Owner)
	require.Equal(t, accountInDB.Balance, account.Balance)
	require.Equal(t, accountInDB.Currency, account.Currency)
}

func TestUpdateAccount(t *testing.T) {
	accountInDB := CreateRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      accountInDB.ID,
		Balance: util.RandomMoney(),
	}
	account, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, account.ID, account.ID)
	require.Equal(t, account.Balance, arg.Balance)
}

func TestDeleteAccount(t *testing.T) {
	accountInDB := CreateRandomAccount(t)
	err := testQueries.DeleteAccounts(context.Background(), accountInDB.ID)
	require.NoError(t, err)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
