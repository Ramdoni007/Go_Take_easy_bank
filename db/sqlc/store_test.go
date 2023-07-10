package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

// update worker next time add unit testing this bellow function :)
func TestTransferTx(t *testing.T) {
	existed := make(map[int]bool)
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	fmt.Println("SEBELUM SALDO DI UPDATED :", account1.Balance, account2.Balance)

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

		// check Accounts
		// cek dimana saldo akun keluar
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, account1.ID, fromAccount.ID)

		// cek dimana saldo itu masuk
		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, account2.ID, toAccount.ID)

		// cek perbedaan anatara akun yang melakukan transaksi
		fmt.Println("SALDO SETIAP TRANSAKSI:", fromAccount.Balance, toAccount.Balance)
		diff1 := account1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - account2.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0) // jadi jumlah transaksid ke-1 , ke-2 , ke -3 dst harus sesuai kelipatanya

		// ini buat variable k untuk meng analogikan saldo nya
		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true

	}

	// Check Final update Balance saldo kita. dengan cara check database
	UpdatedAccount1, err := testQueris.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	UpdatedAccount2, err := testQueris.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println("SETELAH UPDATE SALDO:", UpdatedAccount1.Balance, UpdatedAccount2.Balance)
	require.Equal(t, account1.Balance-int64(n)*amount, UpdatedAccount1.Balance)
	require.Equal(t, account2.Balance+int64(n)*amount, UpdatedAccount2.Balance)

}

//func TestTransferTxDeadlock(t *testing.T) {
//
//	store := NewStore(testDB)
//
//	account1 := createRandomAccount(t)
//	account2 := createRandomAccount(t)
//	fmt.Println("SEBELUM SALDO DI UPDATED :", account1.Balance, account2.Balance)
//
//	n := 10
//	amount := int64(10)
//
//	//run n concurrent tranfers transaction
//	errs := make(chan error)
//	for i := 0; i < n; i++ {
//		fromAccountID := account1.ID
//		toAccountID := account2.ID
//
//		if i%2 == 1 {
//			fromAccountID = account2.ID
//			toAccountID = account1.ID
//		}
//
//		go func() {
//			_, err := store.TranferTx(context.Background(), TransferTxParams{
//				FromAccountID: fromAccountID,
//				ToAccountID:   toAccountID,
//				Amount:        amount,
//			})
//			errs <- err
//
//		}()
//	}
//
//	// check result
//	for i := 0; i < n; i++ {
//		err := <-errs
//		require.NoError(t, err)
//	}
//
//	// Check Final update Balance saldo kita. dengan cara check database
//	UpdatedAccount1, err := testQueris.GetAccount(context.Background(), account1.ID)
//	require.NoError(t, err)
//	UpdatedAccount2, err := testQueris.GetAccount(context.Background(), account2.ID)
//	require.NoError(t, err)
//
//	fmt.Println("SETELAH UPDATE SALDO:", UpdatedAccount1.Balance, UpdatedAccount2.Balance)
//	require.Equal(t, account1.Balance, UpdatedAccount1.Balance)
//	require.Equal(t, account2.Balance, UpdatedAccount2.Balance)
//
//}

func TestTransferTxDeadlock(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	store := NewStore(testDB)
	fmt.Println(">> before:", account1.Balance, account2.Balance)

	n := 10
	amount := int64(10)
	errs := make(chan error)

	for i := 0; i < n; i++ {
		fromAccountID := account1.ID
		toAccountID := account2.ID

		if i%2 == 1 {
			fromAccountID = account2.ID
			toAccountID = account1.ID
		}

		go func() {
			_, err := store.TranferTx(context.Background(), TransferTxParams{
				FromAccountID: fromAccountID,
				ToAccountID:   toAccountID,
				Amount:        amount,
			})

			errs <- err
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}

	// check the final updated balance
	updatedAccount1, err := testQueris.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := testQueris.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)
	require.Equal(t, account1.Balance, updatedAccount1.Balance)
	require.Equal(t, account2.Balance, updatedAccount2.Balance)

}
