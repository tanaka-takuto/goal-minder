package model

import (
	"context"
)

type GoalRepository interface {
	// Create 目標を作成する
	Create(ctx context.Context, con ContextExecutor, goal Goal) (*Goal, error)
}
