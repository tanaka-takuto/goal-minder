package application_error

// UnauthorizedError 認証されていません
type UnauthorizedError struct {
	ApplicationError
}

// NewUnauthorizedError 認証されていませんを生成する
func NewUnauthorizedError() UnauthorizedError {
	return UnauthorizedError{
		ApplicationError{
			message: "認証されていません",
		},
	}
}
