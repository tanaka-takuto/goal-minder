package model

import (
	"fmt"
	"goal-minder/domain/vo"

	validation "github.com/go-ozzo/ozzo-validation"
)

// RawLoginPassword 生ログインパスワード
type RawLoginPassword string

const (
	minRawLoginPasswordLength = 8
	maxRawLoginPasswordLength = 20
)

func (p RawLoginPassword) Validate() error {
	return validation.Validate(string(p),
		validation.Required.Error("パスワードは必須です"),
		validation.Length(8, 20).Error(fmt.Sprintf("パスワードは%d文字以上%d文字以下です", minRawLoginPasswordLength, maxRawLoginPasswordLength)),
	)
}

// LoginPassword ログインパスワード
type LoginPassword vo.HashedString

// NewLoginPassword ログインパスワードを作成する
func NewLoginPassword(rawLoginPassword RawLoginPassword) LoginPassword {
	return LoginPassword(vo.NewHashedString(string(rawLoginPassword)))
}

// ValidString 文字列が正しいかどうかを確認する
func (p LoginPassword) ValidString(plainStr string) error {
	return vo.HashedString(p).ValidString(plainStr)
}
