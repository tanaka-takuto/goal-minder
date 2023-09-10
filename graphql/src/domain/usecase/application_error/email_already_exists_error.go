package application_error

// EmailAlreadyExistsError 指定されたメールアドレスは既に登録されています
type EmailAlreadyExistsError struct {
	ApplicationError
}

var EmailAlreadyExistsErrorInstanse = EmailAlreadyExistsError{
	ApplicationError{
		message: "指定されたメールアドレスは既に登録されています",
	},
}
