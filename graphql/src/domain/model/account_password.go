package model

import (
	"time"

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

// AccountPassword アカウントパスワード
type AccountPassword struct {
	AccountID  AccountID     // アカウントID
	Password   LoginPassword // ログインパスワード
	SetAt      time.Time     // 設定日時
	LoggedInAt time.Time     // 設定日時
}

// NewAccountPassword アカウントログインを作成する
func NewAccountPassword(accountID AccountID, password LoginPassword) AccountPassword {
	now := time.Now()
	return AccountPassword{
		AccountID:  accountID,
		Password:   password,
		SetAt:      now,
		LoggedInAt: now,
	}
}

// Login ログインする
func (a *AccountPassword) Login(plainStr string) error {
	if err := a.Password.ValidString(plainStr); err != nil {
		return err
	}

	a.LoggedInAt = time.Now()

	return nil
}
