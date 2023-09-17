package vo

import "goal-minder/domain/model"

type AccountID GlobalUniqueID

const accountIDKKey = "account"

// NewAccountID アカウントIDを作成する
func NewAccountID(accountID model.AccountID) AccountID {
	return AccountID(newGlobalUniqueID(accountIDKKey, int(accountID)))
}

// Decode アカウントIDを取得する
func (a AccountID) Decode() (model.AccountID, error) {
	decodedID, err := GlobalUniqueID(a).decodeByKey(accountIDKKey)
	if err != nil {
		return 0, err
	}

	return model.AccountID(*decodedID), nil
}

// String 文字列に変換する
func (a AccountID) String() string {
	return string(a)
}
