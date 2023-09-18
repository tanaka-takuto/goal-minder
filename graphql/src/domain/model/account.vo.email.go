package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// AccountEmail メールアドレス
type AccountEmail string

func (e AccountEmail) Validate() error {
	return validation.Validate(string(e),
		validation.Required.Error("メールアドレスは必須です"),
		is.Email.Error("メールアドレスの形式が正しくありません"),
	)
}
