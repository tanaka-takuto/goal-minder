package application_error

// UnauthorizedError 認証されていません
type UnauthorizedError struct {
	ApplicationError
}

var UnauthorizedErrorInstanse = UnauthorizedError{
	ApplicationError{
		message: "認証されていません",
	},
}
