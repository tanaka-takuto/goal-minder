package repository

import (
	"context"

	"goal-minder/adapter/sqlboiler/models"
	"goal-minder/domain/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type accountRepository struct{}

func NewAccountRepository() model.AccountRepository {
	return &accountRepository{}
}

// Create implements model.AccountRepository.
func (*accountRepository) Create(ctx context.Context, con model.ContextExecutor, account model.Account) (*model.Account, error) {
	var dAccount models.Account
	dAccount.Name = string(account.Name)
	dAccount.Email = string(account.Email)

	err := dAccount.Insert(ctx, con, boil.Infer())
	if err != nil {
		return nil, err
	}

	// TODO: mapper作りたい
	return &model.Account{
		ID:    model.AccountID(dAccount.ID),
		Name:  model.AccountName(dAccount.Name),
		Email: model.AccountEmail(dAccount.Email),
	}, nil
}

// ExistsAccountByEmail implements model.AccountRepository.
func (*accountRepository) ExistsAccountByEmail(ctx context.Context, con model.ContextExecutor, email model.AccountEmail) (bool, error) {
	return models.Accounts(
		models.AccountWhere.Email.EQ(string(email)),
	).Exists(ctx, con)
}

// FindByAccountID implements model.AccountRepository.
func (*accountRepository) FindByAccountID(ctx context.Context, con model.ContextExecutor, accountID model.AccountID) (*model.Account, error) {
	dAccount, err := models.Accounts(
		models.AccountWhere.ID.EQ(int64(accountID)),
	).One(ctx, con)

	if err != nil {
		return nil, err
	}

	// TODO: mapper作りたい
	return &model.Account{
		ID:    model.AccountID(dAccount.ID),
		Name:  model.AccountName(dAccount.Name),
		Email: model.AccountEmail(dAccount.Email),
	}, nil
}
