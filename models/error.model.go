package models

import "fmt"

type ErrorResponse struct {
	Code        int
	Category    string
	Description string
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("%s error (%d): %s", e.Category, e.Code, e.Description)
}

func NewError(code int, category string, description string) ErrorResponse {
	return ErrorResponse{
		Code:        code,
		Category:    category,
		Description: description,
	}
}
