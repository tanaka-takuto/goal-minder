package usecase

import (
	"context"
	"database/sql"
	"time"

	"goal-minder/domain/model"
	"goal-minder/domain/usecase/application_error"
	"goal-minder/sdk"
)

type SetGoalUsecase struct {
	model.AccountRepository
	model.GoalRepository
}

type SetGoalInput struct {
	AccountID model.AccountID
	Name      model.GoalName
	Detail    model.GoalDetail
	Deadline  *model.GoalDeadline
	Scale     *model.GoalScale
}

func NewSetGoalInput(accountID model.AccountID, name string, detail string, deadline *time.Time, scale *int) (*SetGoalInput, *application_error.ValidationError) {
	validationErrorBuilder := application_error.NewValidationErrorBuilder()

	n, err := model.NewGoalName(name)
	if err != nil {
		validationErrorBuilder.AddError("name", err.Error())
	}

	det, err := model.NewGoalDetail(detail)
	if err != nil {
		validationErrorBuilder.AddError("detail", err.Error())
	}

	var ded *model.GoalDeadline
	if deadline != nil {
		ded = sdk.Ptr(model.NewGoalDeadline(*deadline))
	}

	s, err := model.NewGoalScale(*scale)
	if err != nil {
		validationErrorBuilder.AddError("scale", err.Error())
	}

	if validationError := validationErrorBuilder.Build(); validationError != nil {
		return nil, validationError
	}

	return &SetGoalInput{
		AccountID: accountID,
		Name:      *n,
		Detail:    *det,
		Deadline:  ded,
		Scale:     s,
	}, nil
}

func (u SetGoalUsecase) Execute(ctx context.Context, db *sql.DB, input SetGoalInput) (*model.Goal, *application_error.AccountNotFoundError, error) {
	var goal *model.Goal
	var accountNotFoundError *application_error.AccountNotFoundError
	err := Transaction(ctx, db, func(ctx context.Context, tx model.ContextExecutor) error {

		// アカウントの存在チェック
		a, err := u.AccountRepository.FindByAccountID(ctx, tx, input.AccountID)
		if err != nil {
			return err
		}
		if a == nil {
			accountNotFoundError = sdk.Ptr(application_error.NewAccountNotFoundError())
			return nil
		}

		// 目標を作成
		newGoal := model.NewGoal(input.AccountID, input.Name, input.Detail, input.Deadline, input.Scale)
		g, err := u.GoalRepository.Create(ctx, tx, newGoal)
		if err != nil {
			return err
		}
		goal = g

		return nil
	})
	if err != nil {
		return nil, nil, err
	} else if accountNotFoundError != nil {
		return nil, accountNotFoundError, nil
	}

	return goal, nil, nil
}
