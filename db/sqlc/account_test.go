package db

import (
	"context"
	"github.com/ramdoni007/Take_Easy_Bank/util"
	"github.com/stretchr/testify/require"
	"testing"

)


func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner: util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account,err := testQueris.CreateAccount(context.Background(),arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner,account.Owner)
	require.Equal(t, arg.Balance,account.Balance)
	require.Equal(t, arg.Currency,account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}


