package context

import (
	"context"

	"goal-minder/domain/model"
)

type accountID struct{}

var accountIDKey = accountID{}

// SetAccountID アカウントIDをコンテキストに設定する
func SetAccountID(ctx context.Context, accountID model.AccountID) context.Context {
	return context.WithValue(ctx, accountIDKey, accountID)
}

// GetAccountID アカウントIDをコンテキストから取得する
func GetAccountID(ctx context.Context) *model.AccountID {
	accountID, ok := ctx.Value(accountIDKey).(model.AccountID)
	if !ok {
		return nil
	}

	return &accountID
}
