package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	graphql1 "goal-minder/adapter/graphql"
	graphql_model "goal-minder/adapter/graphql/model"
	"goal-minder/adapter/graphql/vo"
	"goal-minder/cmd/di"
	modelContext "goal-minder/domain/context"
	"goal-minder/domain/model"
	"goal-minder/domain/usecase"
	"goal-minder/infra/db"
)

// Account is the resolver for the account field.
func (r *meResolver) Account(ctx context.Context, obj *graphql_model.Me) (*graphql_model.Account, error) {
	var accountID model.AccountID
	if aID := modelContext.GetAccountID(ctx); aID != nil {
		accountID = *aID
	} else {
		// ログイン時のみコンテキストにアカウントIDが設定されていないためオブジェクトから取得する
		aIDByObj, err := vo.AccountID.Decode(obj.Account.ID)
		if err != nil {
			return nil, err
		}
		accountID = aIDByObj
	}

	account, err := di.AccountByIDUsecase().Execute(ctx, db.Con, usecase.AccountByIDInput{AccountID: accountID})
	if err != nil {
		return nil, err
	}
	return &graphql_model.Account{
		ID:    vo.AccountID.New(account.ID),
		Name:  string(account.Name),
		Email: string(account.Email),
	}, nil
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*graphql_model.Me, error) {
	return &graphql_model.Me{}, nil
}

// Me returns graphql1.MeResolver implementation.
func (r *Resolver) Me() graphql1.MeResolver { return &meResolver{r} }

// Query returns graphql1.QueryResolver implementation.
func (r *Resolver) Query() graphql1.QueryResolver { return &queryResolver{r} }

type meResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
