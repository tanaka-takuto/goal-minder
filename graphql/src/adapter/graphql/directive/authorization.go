package directive

import (
	"context"
	modelContext "goal-minder/domain/context"
	"goal-minder/domain/usecase/application_error"

	"github.com/99designs/gqlgen/graphql"
)

// AuthorizationDirective 認証認可ディレクティブ
func AuthorizationDirective(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	accountID := modelContext.GetAccountID(ctx)

	if accountID == nil {
		return nil, &application_error.UnauthorizedErrorInstanse
	}

	// TODO: 操作ごとに認可があればここでチェックする

	return next(ctx)
}
