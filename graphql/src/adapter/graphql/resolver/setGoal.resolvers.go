package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	graphql_model "goal-minder/adapter/graphql/model"
	"goal-minder/adapter/graphql/scalar"
	"goal-minder/adapter/graphql/vo"
	"goal-minder/cmd/di"
	modelContext "goal-minder/domain/context"
	"goal-minder/domain/usecase"
	"goal-minder/domain/usecase/application_error"
	"goal-minder/infra/db"
	"goal-minder/sdk"
)

// SetGoal is the resolver for the setGoal field.
func (r *mutationResolver) SetGoal(ctx context.Context, input *graphql_model.SetGoalInput) (graphql_model.SetGoaltPayload, error) {
	accountID := modelContext.GetAccountID(ctx)
	if accountID == nil {
		return graphql_model.AccountNotFoundError{
			Message: sdk.Ptr(application_error.NewAccountNotFoundError().ApplicationError).Error(),
		}, nil
	}

	uInput, vErr := usecase.NewSetGoalInput(*accountID, input.Name, input.Detail, input.Deadline.ToTime(), input.Scale)
	if vErr != nil {
		return graphql_model.NewValidationError(*vErr), nil
	}

	goal, accountNotFoundError, err := di.SetGoalUsecase().Execute(ctx, db.Con, *uInput)
	if err != nil {
		return nil, err
	} else if accountNotFoundError != nil {
		return graphql_model.AccountNotFoundError{
			Message: accountNotFoundError.Error(),
		}, nil
	}

	// TODO: マッパーがほしい
	var deadline *scalar.Date
	if goal.Deadline != nil {
		deadline = sdk.Ptr(scalar.Date(*goal.Deadline))
	}

	var scale *int
	if goal.Scale != nil {
		scale = sdk.Ptr(int(*goal.Scale))
	}

	return &graphql_model.Goal{
		ID:       vo.GoalID.New(goal.ID),
		Name:     string(goal.Name),
		Detail:   string(goal.Detail),
		Deadline: deadline,
		Scale:    scale,
	}, nil
}
