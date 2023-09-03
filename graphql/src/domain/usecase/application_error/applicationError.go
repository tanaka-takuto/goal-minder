package application_error

// ApplicationError アプリケーションエラー
type ApplicationError struct {
	error

	// message エラーメッセージ
	message string
}

// Error エラーメッセージ
func (e *ApplicationError) Error() string {
	return e.message
}
