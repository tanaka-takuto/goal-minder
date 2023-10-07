package model

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

// GoalDetail 目標詳細
type GoalDetail string

const (
	maxLengthGoalDetail = 1000 // 目標詳細の最大長
)

// validate 目標詳細のバリデーション
func (d GoalDetail) validate() error {
	return validation.Validate(string(d),
		validation.Required.Error("目標詳細は必須です"),
		validation.Length(0, maxLengthGoalDetail).Error(fmt.Sprintf("目標詳細は%d文字以下です", maxLengthGoalDetail)),
	)
}

// NewGoalDetail 目標詳細を生成する
func NewGoalDetail(detail string) (*GoalDetail, error) {
	d := GoalDetail(detail)
	if err := d.validate(); err != nil {
		return nil, err
	}
	return &d, nil
}
