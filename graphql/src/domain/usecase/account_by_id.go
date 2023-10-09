package usecase

import (
	"context"
	"database/sql"

	"goal-minder/domain/model"
)

type AccountByIDUsecase struct {
	model.AccountRepository
}

type AccountByIDInput struct {
	AccountID model.AccountID
}

func (u AccountByIDUsecase) Execute(ctx context.Context, db *sql.DB, input AccountByIDInput) (*model.Account, error) {
	account, err := u.AccountRepository.FindByAccountID(ctx, db, input.AccountID)
	if err != nil {
		return nil, err
	}

	return account, nil
}
