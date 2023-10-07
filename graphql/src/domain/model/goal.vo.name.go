package model

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

// GoalName 目標名
type GoalName string

const (
	minLengthGoalName = 10
	maxLengthGoalName = 100
)

func (n GoalName) Validate() error {
	return validation.Validate(string(n),
		validation.Required.Error("目標名は必須です"),
		validation.Length(minLengthGoalName, maxLengthGoalName).Error(fmt.Sprintf("目標名は%d文字以上%d文字以下です", minLengthGoalName, maxLengthGoalName)),
	)
}
