package model

import "fmt"

// AccountID アカウントID
type AccountID int

func (id AccountID) String() string {
	return fmt.Sprintf("%d", id)
}

// AccountName 名前
type AccountName string

// AccountEmail メールアドレス
type AccountEmail string
