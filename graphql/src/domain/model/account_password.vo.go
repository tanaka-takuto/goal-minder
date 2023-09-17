package model

import (
	"goal-minder/domain/vo"
)

// LoginPassword ログインパスワード
type LoginPassword vo.HashedString

// NewLoginPassword ログインパスワードを作成する
func NewLoginPassword(plainStr string) LoginPassword {
	return LoginPassword(vo.NewHashedString(plainStr))
}

// ValidString 文字列が正しいかどうかを確認する
func (p LoginPassword) ValidString(plainStr string) error {
	return vo.HashedString(p).ValidString(plainStr)
}
