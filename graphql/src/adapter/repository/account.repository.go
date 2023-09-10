package repository

import (
	"context"

	"goal-minder/adapter/sqlboiler/models"
	"goal-minder/domain/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type AccountRepository struct {
	db boil.ContextExecutor
}

func NewAccountRepository(db boil.ContextExecutor) *AccountRepository {
	return &AccountRepository{db: db}
}

// Create アカウントを作成する
func (r *AccountRepository) Create(ctx context.Context, account model.Account) (*model.Account, error) {
	var dAccount models.Account
	dAccount.Name = string(account.Name)
	dAccount.Email = string(account.Email)

	err := dAccount.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &model.Account{
		ID:    model.AccountID(dAccount.ID),
		Name:  model.AccountName(dAccount.Name),
		Email: model.AccountEmail(dAccount.Email),
	}, nil
}

// ExistsAccountByEmail メールアドレスが存在するかどうかを確認する
func (r *AccountRepository) ExistsAccountByEmail(ctx context.Context, email model.AccountEmail) (bool, error) {
	return models.Accounts(
		models.AccountWhere.Email.EQ(string(email)),
	).Exists(ctx, r.db)
}
