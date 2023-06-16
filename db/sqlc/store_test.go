package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

// update worker next time add unit testing this bellow function :)
func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	n := 5
	amount := int64(10)

	//run n concurrent tranfers transaction
	errs := make(chan error)
	results := make(chan TransferTxResult)
	for i := 0; i < n; i++ {

		go func() {
			result, err := store.TranferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})
			errs <- err
			results <- result
		}()
	}

	// check result
	for i := 0; i < n; i++ {

		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		//Check Transfers
		transfers := result.Transfer
		require.NotEmpty(t, transfers)
		require.Equal(t, account1.ID, transfers.FromAccountID)
		require.Equal(t, account2.ID, transfers.ToAccountID)
		require.Equal(t, amount, transfers.Amount)
		require.NotZero(t, transfers.ID)
		require.NotZero(t, transfers.CreatedAt)

		// track Record Data in Database
		_, err = store.GetTransfer(context.Background(), transfers.ID)
		require.NoError(t, err)

		//Check entries
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)

		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		//track Record Data in database
		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		//Check Entries
		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, account2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// Todo Add check acoount Balance Soon

	}

}
