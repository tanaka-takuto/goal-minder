package application_error

// ValidationError バリデーションエラー
type ValidationError struct {
	ApplicationError
	Details map[string][]string // バリデーションエラー詳細
}

// NewValidationError バリデーションエラーを生成する
func NewValidationError() ValidationError {
	return ValidationError{
		ApplicationError: ApplicationError{
			message: "バリデーションエラー",
		},
		Details: map[string][]string{},
	}
}

// Add バリデーションエラー詳細を追加する
func (v *ValidationError) Add(field string, message string) {
	v.Details[field] = append(v.Details[field], message)
}

// HasError バリデーションエラーがあるかどうかを返す
func (v *ValidationError) HasError() bool {
	return len(v.Details) > 0
}
