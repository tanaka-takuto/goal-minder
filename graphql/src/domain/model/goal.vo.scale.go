package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// GoalScale 目標規模
type GoalScale int

// validate 目標詳細のバリデーション
func (s GoalScale) validate() error {
	return validation.Validate(int(s),
		validation.Min(0).Error("目標規模は0以上です"))
}

// NewGoalScale 目標詳細を生成する
func NewGoalScale(scale int) (*GoalScale, error) {
	s := GoalScale(scale)
	if err := s.validate(); err != nil {
		return nil, err
	}
	return &s, nil
}
