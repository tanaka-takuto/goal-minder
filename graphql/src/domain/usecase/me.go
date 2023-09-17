package usecase

import (
	"context"
	"database/sql"

	"goal-minder/domain/model"
)

type MeUsecase struct {
	model.AccountRepository
}

type MeInput struct {
	AccountID model.AccountID
}

func (u MeUsecase) Execute(ctx context.Context, db *sql.DB, input MeInput) (*model.Account, error) {
	account, err := u.AccountRepository.FindByAccountID(ctx, db, input.AccountID)
	if err != nil {
		return nil, err
	}

	return account, nil
}
