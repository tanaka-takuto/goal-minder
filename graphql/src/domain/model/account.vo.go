package model

// AccountID アカウントID
type AccountID int

func (id AccountID) String() string {
	return string(id)
}

// AccountName 名前
type AccountName string

// AccountEmail メールアドレス
type AccountEmail string
