package application_error

// EmailAlreadyExistsError 指定されたメールアドレスは既に登録されています
type EmailAlreadyExistsError struct {
	ApplicationError
}

// NewEmailAlreadyExistsError 指定されたメールアドレスは既に登録されていますを生成する
func NewEmailAlreadyExistsError() EmailAlreadyExistsError {
	return EmailAlreadyExistsError{
		ApplicationError{
			message: "指定されたメールアドレスは既に登録されています",
		},
	}
}
