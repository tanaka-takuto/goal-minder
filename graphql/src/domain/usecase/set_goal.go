package usecase

import (
	"context"
	"database/sql"

	"goal-minder/domain/model"
	applicationerror "goal-minder/domain/usecase/application_error"
	"goal-minder/sdk"

	validation "github.com/go-ozzo/ozzo-validation"
)

type SetGoalUsecase struct {
	model.AccountRepository
	model.AccountPasswordRepository
}

type SetGoalInput struct {
	Name     model.AccountName
	Email    model.AccountEmail
	Password model.RawLoginPassword
}

func (i SetGoalInput) Validate() error {
	err := validation.ValidateStruct(&i,
		validation.Field(&i.Name),
		validation.Field(&i.Email),
		validation.Field(&i.Password),
	)
	return err
}

func (u SetGoalUsecase) Execute(ctx context.Context, db *sql.DB, input SetGoalInput) (*model.Account, *applicationerror.EmailAlreadyExistsError, error) {
	var account *model.Account
	var emailAlreadyExistsError *applicationerror.EmailAlreadyExistsError
	err := Transaction(ctx, db, func(ctx context.Context, tx model.ContextExecutor) error {

		// メールアドレスの存在チェック
		if exists, err := u.AccountRepository.ExistsAccountByEmail(ctx, tx, input.Email); err != nil {
			return err
		} else if exists {
			emailAlreadyExistsError = sdk.Ptr(applicationerror.NewEmailAlreadyExistsError())
			return emailAlreadyExistsError
		}

		// アカウントを作成
		newAccount := model.NewAccount(input.Name, input.Email)
		a, err := u.AccountRepository.Create(ctx, tx, newAccount)
		if err != nil {
			return err
		}
		account = a

		// ログイン情報の保存
		accountPassword := model.NewAccountPassword(account.ID, model.NewLoginPassword(input.Password))
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
