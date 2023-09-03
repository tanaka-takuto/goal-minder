package repository

import (
	"context"
	"database/sql"
)

// Transaction トランザクションを実行する
func Transaction(ctx context.Context, db *sql.DB, txExec func(context.Context, *sql.Tx) error) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			panic(err)
		}
	}()

	err = txExec(ctx, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
