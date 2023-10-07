package application_error

// IncorrectEmailOrPasswordError メールアドレスまたはパスワードが間違っています
type IncorrectEmailOrPasswordError struct {
	ApplicationError
}

// NewIncorrectEmailOrPasswordError メールアドレスまたはパスワードが間違っていますを生成する
func NewIncorrectEmailOrPasswordError() IncorrectEmailOrPasswordError {
	return IncorrectEmailOrPasswordError{
		ApplicationError{
			message: "メールアドレスまたはパスワードが間違っています",
		},
	}
}
