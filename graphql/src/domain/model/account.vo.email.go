package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// AccountEmail メールアドレス
type AccountEmail string

// validate メールアドレスのバリデーション
func (e AccountEmail) validate() error {
	return validation.Validate(string(e),
		validation.Required.Error("メールアドレスは必須です"),
		is.Email.Error("メールアドレスの形式が正しくありません"),
	)
}

// NewAccountEmail メールアドレスを作成する
func NewAccountEmail(str string) (*AccountEmail, error) {
	e := AccountEmail(str)
	if err := e.validate(); err != nil {
		return nil, err
	}

	return &e, nil
}
