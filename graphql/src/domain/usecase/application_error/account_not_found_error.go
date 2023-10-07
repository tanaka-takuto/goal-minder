package application_error

// AccountNotFoundError 指定されたアカウントが見つかりませんでした
type AccountNotFoundError struct {
	ApplicationError
}

// NewAccountNotFoundError 指定されたアカウントが見つかりませんでしたを生成する
func NewAccountNotFoundError() AccountNotFoundError {
	return AccountNotFoundError{
		ApplicationError{
			message: "指定されたアカウントが見つかりませんでした",
		},
	}
}
