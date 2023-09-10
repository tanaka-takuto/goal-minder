package repository

import (
	"context"
	"fmt"

	"goal-minder/adapter/sqlboiler/models"
	"goal-minder/domain/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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
	dAccountPassword.LoggedInAt = accountPassword.LoggedInAt

	err := dAccountPassword.Upsert(ctx, r.db, boil.Infer(), boil.Infer())
	if err != nil {
		return nil, err
	}

	// TODO: mapper作りたい
	return &model.AccountPassword{
		AccountID:  model.AccountID(dAccountPassword.AccountID),
		Password:   model.LoginPassword(dAccountPassword.HashedPassword),
		SetAt:      dAccountPassword.SetAt,
		LoggedInAt: dAccountPassword.LoggedInAt,
	}, nil
}

// FindByEmail メールアドレスからアカウントパスワードを取得する
func (r *AccountPasswordRepository) FindByEmail(ctx context.Context, email model.AccountEmail) (*model.AccountPassword, error) {
	dAccountPassword, err := models.AccountPasswords(
		qm.InnerJoin(fmt.Sprintf("%v a ON a.%v = %v.%v", models.TableNames.Account, models.AccountColumns.ID, models.TableNames.AccountPassword, models.AccountPasswordColumns.AccountID)),
		qm.Where(fmt.Sprintf("a.%v = ?", models.AccountColumns.Email), email),
	).One(ctx, r.db)
	if err != nil {
		return nil, err
	}

	// TODO: mapper作りたい
	return &model.AccountPassword{
		AccountID:  model.AccountID(dAccountPassword.AccountID),
		Password:   model.LoginPassword(dAccountPassword.HashedPassword),
		SetAt:      dAccountPassword.SetAt,
		LoggedInAt: dAccountPassword.LoggedInAt,
	}, nil
}
