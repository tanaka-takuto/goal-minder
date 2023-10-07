package graphql_model

import (
	"goal-minder/domain/usecase/application_error"
)

// NewValidationError バリデーションエラーを作成する
func NewValidationError(err application_error.ValidationError) ValidationError {
	ve := ValidationError{Message: err.Error()}

	for k, v := range err.Details {
		ve.Details = append(ve.Details, &ValidationErrorDetail{
			Field:   k,
			Message: v,
		})
	}

	return ve
}
