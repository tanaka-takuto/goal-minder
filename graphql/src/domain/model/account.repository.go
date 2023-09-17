package model

import (
	"context"
)

type AccountRepository interface {
	// Create アカウントを作成する
	Create(ctx context.Context, con ContextExecutor, account Account) (*Account, error)

	// ExistsAccountByEmail メールアドレスが存在するかどうかを確認する
	ExistsAccountByEmail(ctx context.Context, con ContextExecutor, email AccountEmail) (bool, error)

	// FindByAccountID アカウントIDからアカウントを取得する
	FindByAccountID(ctx context.Context, con ContextExecutor, accountID AccountID) (*Account, error)
}
