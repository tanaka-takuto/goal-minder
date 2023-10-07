package model

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

// AccountName 名前
type AccountName string

const (
	minLengthAccountName = 10
	maxLengthAccountName = 100
)

// validate 名前のバリデーション
func (n AccountName) validate() error {
	return validation.Validate(string(n),
		validation.Required.Error("名前は必須です"),
		validation.Length(minLengthAccountName, maxLengthAccountName).Error(fmt.Sprintf("名前は%d文字以上%d文字以下です", minLengthAccountName, maxLengthAccountName)),
	)
}

// NewAccountName 名前を生成
func NewAccountName(name string) (*AccountName, error) {
	n := AccountName(name)
	if err := n.validate(); err != nil {
		return nil, err
	}
	return &n, nil
}
