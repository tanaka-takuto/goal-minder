package usecase

import (
	"context"
	"database/sql"

	"goal-minder/domain/model"
	"goal-minder/domain/usecase/application_error"
	"goal-minder/sdk"
)

type LoginUsecase struct {
	model.AccountPasswordRepository
}

type LoginInput struct {
	Email    model.AccountEmail
	Password model.RawLoginPassword
}

func NewLoginInput(email string, password string) (*LoginInput, *application_error.ValidationError) {
	validationErrorBuilder := application_error.NewValidationErrorBuilder()

	e, err := model.NewAccountEmail(email)
	if err != nil {
		validationErrorBuilder.AddError("email", err.Error())
	}

	p, err := model.NewRawLoginPassword(password)
	if err != nil {
		validationErrorBuilder.AddError("password", err.Error())
	}

	if validationError := validationErrorBuilder.Build(); validationError != nil {
		return nil, validationError
	}

	return &LoginInput{
		Email:    *e,
		Password: *p,
	}, nil
}

func (u LoginUsecase) Execute(ctx context.Context, db *sql.DB, input LoginInput) (*model.AccountID, *application_error.IncorrectEmailOrPasswordError, error) {
	var accountID *model.AccountID
	var incorrectEmailOrPasswordError *application_error.IncorrectEmailOrPasswordError

	err := Transaction(ctx, db, func(ctx context.Context, tx model.ContextExecutor) error {
		// メールアドレスからアカウントを取得
		ap, err := u.AccountPasswordRepository.FindByEmail(ctx, tx, input.Email)
		if err != nil {
			return err
		}

		// ログインできるかチェック
		if err := ap.Login(input.Password); err != nil {
			incorrectEmailOrPasswordError = sdk.Ptr(application_error.NewIncorrectEmailOrPasswordError())
			return nil
		}

		if _, err := u.AccountPasswordRepository.Save(ctx, tx, *ap); err != nil {
			return err
		}
		accountID = &ap.AccountID

		return nil
	})
	if err != nil {
		return nil, nil, err
	}
	if incorrectEmailOrPasswordError != nil {
		return nil, incorrectEmailOrPasswordError, nil
	}

	return accountID, nil, nil
}
