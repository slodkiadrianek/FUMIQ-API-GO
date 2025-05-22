package schemas

import (
	"errors"
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

func (r *RegisterUser) Validate() (error, z.ZogIssueMap) {
	if r.Password != r.ConfirmPassword {
		err := errors.New("Password and ConfirmPassword must match")
		return err, nil
	}
	errMap := RegisterSchema.Validate(r)
	if errMap != nil {
		return nil, errMap
	}
	return nil, nil
}

func (l *LoginUser) Validate() z.ZogIssueMap {
	errMap := LoginSchema.Validate(l)
	if errMap != nil {
		return errMap
	}
	return nil
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
