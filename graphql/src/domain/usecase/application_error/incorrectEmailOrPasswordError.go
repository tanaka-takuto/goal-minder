package application_error

// IncorrectEmailOrPasswordError メールアドレスまたはパスワードが間違っています
type IncorrectEmailOrPasswordError struct {
	ApplicationError
}

var IncorrectEmailOrPasswordErrorInstanse = IncorrectEmailOrPasswordError{
	ApplicationError{
		message: "メールアドレスまたはパスワードが間違っています",
	},
}
