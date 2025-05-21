package schemas

import (
	"strings"

	z "github.com/Oudwins/zog"
)

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password" `
}
type RegisterUser struct {
	Email           string `json:"email" `
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName" `
	Password        string `json:"password" `
	ConfirmPassword string `json:"confirmPassword" `
}

var RegisterSchema = z.Struct(z.Schema{
	"firstName": z.String().Required(),
	"lastName":  z.String().Required(),
	"email": z.String().Required().Email().Transform(func(val *string, ctx z.Ctx) error {
		*val = strings.ToLower(*val)
		*val = strings.TrimSpace(*val)
		return nil
	}),
	"password":        z.String().Required().Min(8).Max(32).ContainsSpecial().ContainsUpper().ContainsDigit(),
	"confirmPassword": z.String().Required().Min(8).Max(32).ContainsSpecial().ContainsUpper().ContainsDigit(),
})

var LoginSchema = z.Struct(z.Schema{
	"emial": z.String().Required().Email().Transform(func(val *string, ctx z.Ctx) error {
		*val = strings.ToLower(*val)
		*val = strings.TrimSpace(*val)
		return nil
	}),
	"password": z.String().Required().Min(8).Max(32).ContainsSpecial().ContainsUpper().ContainsDigit(),
})
