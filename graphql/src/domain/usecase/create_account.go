package usecase

import (
	"context"
	"database/sql"

	"goal-minder/domain/model"
	applicationerror "goal-minder/domain/usecase/application_error"
	"goal-minder/sdk"
)

type CreateAccountUsecase struct {
	model.AccountRepository
	model.AccountPasswordRepository
}

type createAccountInput struct {
	name     model.AccountName
	email    model.AccountEmail
	password model.RawLoginPassword
}

func NewCreateAccountInput(name string, email string, password string) (*createAccountInput, *applicationerror.ValidationError) {
	validationError := applicationerror.NewValidationError()

	n, err := model.NewAccountName(name)
	if err != nil {
		validationError.Add("name", err.Error())
	}

	e, err := model.NewAccountEmail(email)
	if err != nil {
		validationError.Add("email", err.Error())
	}

	p, err := model.NewRawLoginPassword(password)
	if err != nil {
		validationError.Add("password", err.Error())
	}

	if validationError.HasError() {
		return nil, &validationError
	}

	return &createAccountInput{
		name:     *n,
		email:    *e,
		password: *p,
	}, nil
}

func (u CreateAccountUsecase) Execute(ctx context.Context, db *sql.DB, input createAccountInput) (*model.Account, *applicationerror.EmailAlreadyExistsError, error) {
	var account *model.Account
	var emailAlreadyExistsError *applicationerror.EmailAlreadyExistsError
	err := Transaction(ctx, db, func(ctx context.Context, tx model.ContextExecutor) error {

		// メールアドレスの存在チェック
		if exists, err := u.AccountRepository.ExistsAccountByEmail(ctx, tx, input.email); err != nil {
			return err
		} else if exists {
			emailAlreadyExistsError = sdk.Ptr(applicationerror.NewEmailAlreadyExistsError())
			return nil
		}

		// アカウントを作成
		newAccount := model.NewAccount(input.name, input.email)
		a, err := u.AccountRepository.Create(ctx, tx, newAccount)
		if err != nil {
			return err
		}
		account = a

		// ログイン情報の保存
		accountPassword := model.NewAccountPassword(account.ID, model.NewLoginPassword(input.password))
		if _, err := u.AccountPasswordRepository.Save(ctx, tx, accountPassword); err != nil {
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
