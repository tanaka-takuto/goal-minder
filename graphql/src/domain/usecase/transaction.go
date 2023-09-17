package usecase

import (
	"context"
	"goal-minder/domain/model"
)

// Transaction トランザクションを実行する
func Transaction(ctx context.Context, db model.Beginner, txExec func(context.Context, model.ContextExecutor) error) error {
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
