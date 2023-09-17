package usecase

import (
	"context"
	"database/sql"

	"goal-minder/domain/model"
	applicationerror "goal-minder/domain/usecase/application_error"
)

type LoginUsecase struct {
	model.AccountPasswordRepository
}

type LoginInput struct {
	Email    model.AccountEmail
	Password string
}

func (u LoginUsecase) Execute(ctx context.Context, db *sql.DB, input LoginInput) (*model.Account, *applicationerror.IncorrectEmailOrPasswordError, error) {
	var accountPassword *model.AccountPassword
	var incorrectEmailOrPasswordError *applicationerror.IncorrectEmailOrPasswordError

	err := Transaction(ctx, db, func(ctx context.Context, tx model.ContextExecutor) error {
		// メールアドレスからアカウントを取得
		ap, err := u.AccountPasswordRepository.FindByEmail(ctx, tx, input.Email)
		if err != nil {
			return err
		}
		accountPassword = ap

		// ログインできるかチェック
		if err := accountPassword.Login(input.Password); err != nil {
			incorrectEmailOrPasswordError = &applicationerror.IncorrectEmailOrPasswordErrorInstanse
			return incorrectEmailOrPasswordError
		}

		if _, err := u.AccountPasswordRepository.Save(ctx, tx, *accountPassword); err != nil {
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
