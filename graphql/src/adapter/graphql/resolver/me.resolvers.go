package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.37

import (
	"context"
	"fmt"

	"github.com/tanaka-takuto/goal-minder/adapter/graphql"
	graphql_model "github.com/tanaka-takuto/goal-minder/adapter/graphql/model"
	modelContext "github.com/tanaka-takuto/goal-minder/domain/context"
)

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*graphql_model.Account, error) {
	accountID := modelContext.GetAccountID(ctx)
	fmt.Println(*accountID)
	return &graphql_model.Account{
		ID:    "1",
		Name:  "tanaka",
		Email: "",
	}, nil
}

// Query returns graphql1.QueryResolver implementation.
func (r *Resolver) Query() graphql.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
