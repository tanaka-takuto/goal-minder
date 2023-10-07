package repository

import (
	"context"
	"goal-minder/adapter/sqlboiler/models"
	"goal-minder/domain/model"
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type goalRepository struct{}

func NewGoalRepository() model.GoalRepository {
	return &goalRepository{}
}

// Create implements model.GoalRepository.
func (*goalRepository) Create(ctx context.Context, con model.ContextExecutor, goal model.Goal) (*model.Goal, error) {

	dGoal := models.Goal{}
	dGoal.AccountID = int64(goal.AccountID)
	dGoal.Name = string(goal.Name)
	dGoal.Detail = string(goal.Detail)
	if goal.Deadline != nil {
		dGoal.Deadline = null.TimeFrom(time.Time(*goal.Deadline))
	}
	if goal.Scale != nil {
		dGoal.Scale = null.IntFrom(int(*goal.Scale))
	}

	if err := dGoal.Insert(ctx, con, boil.Infer()); err != nil {
		return nil, err
	}

	goal.ID = model.GoalID(dGoal.ID)

	return &goal, nil
}
