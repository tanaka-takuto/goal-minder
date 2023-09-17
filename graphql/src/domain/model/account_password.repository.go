package model

import (
	"context"
)

type AccountPasswordRepository interface {
	// Save アカウントパスワードを保存する
	Save(ctx context.Context, con ContextExecutor, accountPassword AccountPassword) (*AccountPassword, error)

	// FindByEmail メールアドレスからアカウントパスワードを取得する
	FindByEmail(ctx context.Context, con ContextExecutor, email AccountEmail) (*AccountPassword, error)
}
