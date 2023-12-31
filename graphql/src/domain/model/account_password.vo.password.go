package model

import (
	"goal-minder/domain/vo"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

// RawLoginPassword 生ログインパスワード
type RawLoginPassword string

// validate 文字列が正しいかどうかを確認する
func (p RawLoginPassword) validate() error {
	return validation.Validate(string(p),
		validation.Required.Error("パスワードは必須です"),
		validation.Match(regexp.MustCompile(`^[!-~]{8,20}$`)).Error("パスワードは半角英数字記号8文字以上20文字以下で入力してください"),
	)
}

// NewRawLoginPassword 生ログインパスワードを作成する
func NewRawLoginPassword(str string) (*RawLoginPassword, error) {
	p := RawLoginPassword(str)
	if err := p.validate(); err != nil {
		return nil, err
	}

	return &p, nil
}

// LoginPassword ログインパスワード
type LoginPassword vo.HashedString

// NewLoginPassword ログインパスワードを作成する
func NewLoginPassword(rawLoginPassword RawLoginPassword) LoginPassword {
	return LoginPassword(vo.NewHashedString(string(rawLoginPassword)))
}

// ValidString 文字列が正しいかどうかを確認する
func (p LoginPassword) ValidString(rawLoginPassword RawLoginPassword) error {
	return vo.HashedString(p).ValidString(string(rawLoginPassword))
}
