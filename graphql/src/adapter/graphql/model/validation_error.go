package graphql_model

import "encoding/json"

// NewValidationError バリデーションエラーを作成する
func NewValidationError(err error) ValidationError {
	ve := ValidationError{Message: "バリデーションエラー"}

	validationErrStr, _ := json.Marshal(err)
	var validationDetailMap map[string]string
	json.Unmarshal(validationErrStr, &validationDetailMap)
	for k, v := range validationDetailMap {
		ve.Details = append(ve.Details, &ValidationErrorDetail{
			Field:   k,
			Message: v,
		})
	}

	return ve
}
