package model

// Account アカウント
type Account struct {
	ID    AccountID    // アカウントID
	Name  AccountName  // 名前
	Email AccountEmail // メールアドレス
}

// NewAccount アカウントを作成する
func NewAccount(name AccountName, email AccountEmail) Account {
	return Account{
		Name:  name,
		Email: email,
	}
}
