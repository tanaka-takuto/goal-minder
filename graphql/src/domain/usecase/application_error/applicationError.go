package application_error

// ApplicationError アプリケーションエラー
type ApplicationError struct {
	// message エラーメッセージ
	message string
}

// Error エラーメッセージ
func (e *ApplicationError) Error() string {
	return e.message
}
