package rest

import "strings"

type ErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

type ValidationError struct {
	ErrorResponses []*ErrorResponse
}

func (e ValidationError) Error() string {
	sb := strings.Builder{}
	for _, e := range e.ErrorResponses {
		sb.WriteString(e.FailedField + ": " + e.Tag + ": " + e.Value + "\n")
	}
	return sb.String()
}

func NewValidationError(errors []*ErrorResponse) error {
	return &ValidationError{
		ErrorResponses: errors,
	}
}
