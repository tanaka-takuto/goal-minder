package model

import (
	"time"

	"github.com/tanaka-takuto/goal-minder/domain/vo"
)

// LoginPassword ログインパスワード
type LoginPassword vo.HashedString

// NewLoginPassword ログインパスワードを作成する
func NewLoginPassword(plainStr string) LoginPassword {
	return LoginPassword(vo.NewHashedString(plainStr))
}

// AccountPassword アカウントパスワード
type AccountPassword struct {
	AccountID AccountID     // アカウントID
	Password  LoginPassword // ログインパスワード
	SetAt     time.Time     // 設定日時
}

// NewAccountPassword アカウントログインを作成する
func NewAccountPassword(accountID AccountID, password LoginPassword) AccountPassword {
	now := time.Now()
	return AccountPassword{
		AccountID: accountID,
		Password:  password,
		SetAt:     now,
	}
}
