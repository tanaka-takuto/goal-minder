package usecase

import (
	"context"
	"database/sql"

	"goal-minder/adapter/repository"
	"goal-minder/domain/model"
	applicationerror "goal-minder/domain/usecase/application_error"
)

type CreateAccountInput struct {
	Name     model.AccountName
	Email    model.AccountEmail
	Password model.LoginPassword
}

func CreateAccount(ctx context.Context, db *sql.DB, input CreateAccountInput) (*model.Account, *applicationerror.EmailAlreadyExistsError, error) {
	var account *model.Account
	var emailAlreadyExistsError *applicationerror.EmailAlreadyExistsError
	err := repository.Transaction(ctx, db, func(ctx context.Context, tx *sql.Tx) error {
		accountRepo := repository.NewAccountRepository(tx)

		// メールアドレスの存在チェック
		if exists, err := accountRepo.ExistsAccountByEmail(ctx, input.Email); err != nil {
			return err
		} else if exists {
			emailAlreadyExistsError = &applicationerror.EmailAlreadyExistsErrorInstanse
			return emailAlreadyExistsError
		}

		// アカウントを作成
		newAccount := model.NewAccount(input.Name, input.Email)
		a, err := accountRepo.Create(ctx, newAccount)
		if err != nil {
			return err
		}
		account = a

		// ログイン情報の保存
		accountPassword := model.NewAccountPassword(account.ID, input.Password)
		if _, err := repository.NewAccountPasswordRepository(tx).Save(ctx, accountPassword); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, nil, err
	} else if emailAlreadyExistsError != nil {
		return nil, emailAlreadyExistsError, nil
	}

	return account, nil, nil
}
