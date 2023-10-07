package model

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

// GoalName 目標名
type GoalName string

const (
	maxLengthGoalName = 100 // 目標名の最大長
)

func (n GoalName) validate() error {
	return validation.Validate(string(n),
		validation.Required.Error("目標名は必須です"),
		validation.Length(0, maxLengthGoalName).Error(fmt.Sprintf("目標名は%d文字以下です", maxLengthGoalName)),
	)
}

// NewGoalName 目標名を生成する
func NewGoalName(name string) (*GoalName, error) {
	n := GoalName(name)
	if err := n.validate(); err != nil {
		return nil, err
	}
	return &n, nil
}
