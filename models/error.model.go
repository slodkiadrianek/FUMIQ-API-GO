package models

type ErrorResponse struct {
	Code        int    `json:"code"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

func (e ErrorResponse) Error() string {
	return e.Description
}
func NewError(code int, category, description string) *ErrorResponse {
	return &ErrorResponse{
		Code:        code,
		Category:    category,
		Description: description,
	}
}
