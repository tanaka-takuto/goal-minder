package repository

import (
	"context"

	models "github.com/tanaka-takuto/goal-minder/adapter/sqlboiler"
	"github.com/tanaka-takuto/goal-minder/domain/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type AccountPasswordRepository struct {
	db boil.ContextExecutor
}

func NewAccountPasswordRepository(db boil.ContextExecutor) *AccountPasswordRepository {
	return &AccountPasswordRepository{db: db}
}

// Save 保存する
func (r *AccountPasswordRepository) Save(ctx context.Context, accountPassword model.AccountPassword) (*model.AccountPassword, error) {
	var dAccountPassword models.AccountPassword
	dAccountPassword.AccountID = int64(accountPassword.AccountID)
	dAccountPassword.HashedPassword = string(accountPassword.Password)
	dAccountPassword.SetAt = accountPassword.SetAt

	err := dAccountPassword.Upsert(ctx, r.db, boil.Infer(), boil.Infer())
	if err != nil {
		return nil, err
	}

	return &model.AccountPassword{
		AccountID: model.AccountID(dAccountPassword.AccountID),
		Password:  model.LoginPassword(dAccountPassword.HashedPassword),
	}, nil
}
