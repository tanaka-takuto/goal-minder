package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.37

import (
	"context"

	graphql1 "github.com/tanaka-takuto/goal-minder/adapter/graphql"
	graphql_model "github.com/tanaka-takuto/goal-minder/adapter/graphql/model"
	"github.com/tanaka-takuto/goal-minder/domain/model"
	"github.com/tanaka-takuto/goal-minder/domain/usecase"
	"github.com/tanaka-takuto/goal-minder/infra/db"
)

// CreateAccount is the resolver for the createAccount field.
func (r *mutationResolver) CreateAccount(ctx context.Context, input *graphql_model.CreateAccountInput) (graphql_model.CreateAccountPayload, error) {
	uInput := usecase.CreateAccountInput{
		Name:     model.AccountName(input.Name),
		Email:    model.AccountEmail(input.Email),
		Password: model.NewLoginPassword(input.Password),
	}

	account, emailAlreadyExistsErr, err := usecase.CreateAccount(ctx, db.Con, usecase.CreateAccountInput(uInput))
	if err != nil {
		return nil, err
	}
	if emailAlreadyExistsErr != nil {
		return graphql_model.EmailAlreadyExistsError{Message: emailAlreadyExistsErr.Error()}, nil
	}

	return &graphql_model.Account{
		ID:    string(account.ID),
		Name:  string(account.Name),
		Email: string(account.Email),
	}, nil
}

// Mutation returns graphql1.MutationResolver implementation.
func (r *Resolver) Mutation() graphql1.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
