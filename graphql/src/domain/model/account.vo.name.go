package model

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

// AccountName 名前
type AccountName string

const (
	minAccountNameLength = 10
	maxAccountNameLength = 100
)

func (n AccountName) Validate() error {
	return validation.Validate(string(n),
		validation.Required.Error("名前は必須です"),
		validation.Length(minAccountNameLength, maxAccountNameLength).Error(fmt.Sprintf("名前は%d文字以上%d文字以下です", minAccountNameLength, maxAccountNameLength)),
	)
}
