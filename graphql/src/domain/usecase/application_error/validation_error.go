package application_error

// validationErrorDetails バリデーションエラー詳細
type validationErrorDetails map[string][]string

// ValidationError バリデーションエラー
type ValidationError struct {
	ApplicationError
	Details validationErrorDetails // バリデーションエラー詳細
}

// validationErrorBuilder バリデーションエラービルダー
type validationErrorBuilder struct {
	details validationErrorDetails
}

// NewValidationErrorBuilder バリデーションエラービルダーを作成する
func NewValidationErrorBuilder() validationErrorBuilder {
	return validationErrorBuilder{
		details: make(map[string][]string),
	}
}

// AddError バリデーションエラー詳細を追加する
func (v *validationErrorBuilder) AddError(field string, message string) {
	v.details[field] = append(v.details[field], message)
}

// hasError バリデーションエラーがあるかどうかを返す
func (v *validationErrorBuilder) hasError() bool {
	for _, messages := range v.details {
		if len(messages) > 0 {
			return true
		}
	}

	return false
}

// Build バリデーションエラーを作成する
func (v *validationErrorBuilder) Build() *ValidationError {
	if !v.hasError() {
		return nil
	}

	return &ValidationError{
		ApplicationError: ApplicationError{
			message: "バリデーションエラー",
		},
		Details: v.details,
	}
}
