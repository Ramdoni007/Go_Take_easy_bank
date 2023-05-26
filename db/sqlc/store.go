package db

import (
	"context"
	"database/sql"
	"fmt"
)

//Store Provides all functions to execute db queris and transaction

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}

}

//execTx execute a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v , rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

//
type TranferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TranfeTxResult struct {
	//Problem and solve ulang video sebelumnya buat test entry and tranfers :)
}

//TranferTx performs a money tranfers from one account to the other.
// It creates a tranfers record, add accounts , and update account balance within single database transaction

func (store *Store) TranferTx(ctx context.Context, arg TranferTxParams) (TranferTxResult error)

}
