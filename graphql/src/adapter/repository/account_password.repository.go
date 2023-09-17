package repository

import (
	"context"
	"fmt"

	"goal-minder/adapter/sqlboiler/models"
	"goal-minder/domain/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type accountPasswordRepository struct{}

func NewAccountPasswordRepository() model.AccountPasswordRepository {
	return &accountPasswordRepository{}
}

// FindByEmail implements model.AccountPasswordRepository.
func (r *accountPasswordRepository) FindByEmail(ctx context.Context, con model.ContextExecutor, email model.AccountEmail) (*model.AccountPassword, error) {
	dAccountPassword, err := models.AccountPasswords(
		qm.InnerJoin(fmt.Sprintf("%v a ON a.%v = %v.%v", models.TableNames.Account, models.AccountColumns.ID, models.TableNames.AccountPassword, models.AccountPasswordColumns.AccountID)),
		qm.Where(fmt.Sprintf("a.%v = ?", models.AccountColumns.Email), email),
	).One(ctx, con)
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

// Save implements model.AccountPasswordRepository.
func (r *accountPasswordRepository) Save(ctx context.Context, con model.ContextExecutor, accountPassword model.AccountPassword) (*model.AccountPassword, error) {
	var dAccountPassword models.AccountPassword
	dAccountPassword.AccountID = int64(accountPassword.AccountID)
	dAccountPassword.HashedPassword = string(accountPassword.Password)
	dAccountPassword.SetAt = accountPassword.SetAt
	dAccountPassword.LoggedInAt = accountPassword.LoggedInAt

	err := dAccountPassword.Upsert(ctx, con, boil.Infer(), boil.Infer())
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
