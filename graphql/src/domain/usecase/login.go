package usecase

import (
	"context"
	"database/sql"

	"github.com/tanaka-takuto/goal-minder/adapter/repository"
	"github.com/tanaka-takuto/goal-minder/domain/model"
	applicationerror "github.com/tanaka-takuto/goal-minder/domain/usecase/application_error"
)

type LoginInput struct {
	Email    model.AccountEmail
	Password string
}

func Login(ctx context.Context, db *sql.DB, input LoginInput) (*model.Account, *applicationerror.IncorrectEmailOrPasswordError, error) {
	var accountPassword *model.AccountPassword
	var incorrectEmailOrPasswordError *applicationerror.IncorrectEmailOrPasswordError

	err := repository.Transaction(ctx, db, func(ctx context.Context, tx *sql.Tx) error {
		// メールアドレスからアカウントを取得
		accountPasswordRepo := repository.NewAccountPasswordRepository(tx)
		ap, err := accountPasswordRepo.FindByEmail(ctx, input.Email)
		if err != nil {
			return err
		}
		accountPassword = ap

		// ログインできるかチェック
		if err := accountPassword.Login(input.Password); err != nil {
			incorrectEmailOrPasswordError = &applicationerror.IncorrectEmailOrPasswordErrorInstanse
			return incorrectEmailOrPasswordError
		}

		if _, err := accountPasswordRepo.Save(ctx, *accountPassword); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, nil, err
	}
	if incorrectEmailOrPasswordError != nil {
		return nil, incorrectEmailOrPasswordError, nil
	}

	return &model.Account{ID: accountPassword.AccountID}, nil, nil
}
