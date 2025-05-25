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

type ResetPassword struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

var ResetPasswordSchema = z.Struct(z.Schema{
	"password":        z.String().Required(),
	"confirmPassword": z.String().Required(),
})

func (r *ResetPassword) Validate() (z.ZogIssueMap, error) {
	if r.ConfirmPassword != r.Password {
		err := errors.New("password and confirm password must match")
		return nil, err
	}
	errMap := ChangePasswordSchema.Validate(r)
	if errMap != nil {
		return errMap, nil
	}
	return nil, nil
}

type ChangePassword struct {
	ConfirmPassword string `json:"confirmPassword"`
	NewPassword     string `json:"newPassword"`
	OldPassword     string `json:"oldPassword"`
}

var ChangePasswordSchema = z.Struct(z.Schema{
	"oldPassword":     z.String().Required(),
	"newPassword":     z.String().Required(),
	"confirmPassword": z.String().Required(),
})

func (c *ChangePassword) Validate() (z.ZogIssueMap, error) {
	if c.ConfirmPassword != c.NewPassword {
		err := errors.New("new password and confirm password must match")
		return nil, err
	}
	errMap := ChangePasswordSchema.Validate(c)
	if errMap != nil {
		return errMap, nil
	}
	return nil, nil
}

type DeleteUser struct {
	Password string `json:"password"`
}

func (r *RegisterUser) Validate() (z.ZogIssueMap, error) {
	if r.Password != r.ConfirmPassword {
		err := errors.New("Password and ConfirmPassword must match")
		return nil, err
	}
	errMap := RegisterSchema.Validate(r)
	if errMap != nil {
		return errMap, nil
	}
	return nil, nil
}

func (l *LoginUser) Validate() (z.ZogIssueMap, error) {
	errMap := LoginSchema.Validate(l)
	if errMap != nil {
		return errMap, nil
	}
	return nil, nil
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
