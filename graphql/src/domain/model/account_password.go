package model

import (
	"time"
)

// AccountPassword アカウントパスワード
type AccountPassword struct {
	AccountID  AccountID     // アカウントID
	Password   LoginPassword // ログインパスワード
	SetAt      time.Time     // 設定日時
	LoggedInAt time.Time     // ログイン日時
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
func (a *AccountPassword) Login(rawLoginPassword RawLoginPassword) error {
	if err := a.Password.ValidString(rawLoginPassword); err != nil {
		return err
	}

	a.LoggedInAt = time.Now()

	return nil
}
