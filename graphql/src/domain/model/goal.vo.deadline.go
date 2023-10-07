package model

import (
	"time"
)

// GoalDeadline 目標期限
type GoalDeadline time.Time

// NewGoalDeadline 目標期限を生成
func NewGoalDeadline(deadline time.Time) GoalDeadline {
	return GoalDeadline(deadline)
}
