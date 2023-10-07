package usecase

import (
	"context"
	"database/sql"

	"goal-minder/domain/model"
	applicationerror "goal-minder/domain/usecase/application_error"
	"goal-minder/sdk"
)

type LoginUsecase struct {
	model.AccountPasswordRepository
	model.AccountRepository
}

type LoginInput struct {
	Email    model.AccountEmail
	Password model.RawLoginPassword
}

func NewLoginInput(email string, password string) (*LoginInput, *applicationerror.ValidationError) {
	validationError := applicationerror.NewValidationError()

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

	return &LoginInput{
		Email:    *e,
		Password: *p,
	}, nil
}

func (u LoginUsecase) Execute(ctx context.Context, db *sql.DB, input LoginInput) (*model.Account, *applicationerror.IncorrectEmailOrPasswordError, error) {
	var account *model.Account
	var incorrectEmailOrPasswordError *applicationerror.IncorrectEmailOrPasswordError

	err := Transaction(ctx, db, func(ctx context.Context, tx model.ContextExecutor) error {
		// メールアドレスからアカウントを取得
		ap, err := u.AccountPasswordRepository.FindByEmail(ctx, tx, input.Email)
		if err != nil {
			return err
		}

		// ログインできるかチェック
		if err := ap.Login(input.Password); err != nil {
			incorrectEmailOrPasswordError = sdk.Ptr(applicationerror.NewIncorrectEmailOrPasswordError())
			return nil
		}

		if _, err := u.AccountPasswordRepository.Save(ctx, tx, *ap); err != nil {
			return err
		}

		a, err := u.AccountRepository.FindByAccountID(ctx, tx, ap.AccountID)
		if err != nil {
			return err
		}
		account = a

		return nil
	})
	if err != nil {
		return nil, nil, err
	}
	if incorrectEmailOrPasswordError != nil {
		return nil, incorrectEmailOrPasswordError, nil
	}

	return account, nil, nil
}
