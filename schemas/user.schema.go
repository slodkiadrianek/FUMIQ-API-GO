package schemas

import (
<<<<<<< HEAD
	"errors"
=======
>>>>>>> e50232b (VALIDATION)
	"strings"

	z "github.com/Oudwins/zog"
)

<<<<<<< HEAD
type UserId struct {
	UserId string `json:"userId"`
}

func (u *UserId) Validate() (z.ZogIssueMap, error) {
	errMap := UserIdSchema.Validate(u)
	if errMap != nil {
		return errMap, nil
	}
	return nil, nil
}

var UserIdSchema = z.Struct(z.Schema{
	"userId": z.String().Required(),
})

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUser struct {
	Email           string `json:"email"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

func (r *RegisterUser) Validate() (z.ZogIssueMap, error) {
	if r.Password != r.ConfirmPassword {
		return nil, errors.New("password and ConfirmPassword must match")
	}
	errMap := RegisterSchema.Validate(r)
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
	"newPassword":     z.String().Required().Min(8).Max(32).ContainsSpecial().ContainsUpper().ContainsDigit(),
	"confirmPassword": z.String().Required().Min(8).Max(32).ContainsSpecial().ContainsUpper().ContainsDigit(),
})

func (c *ChangePassword) Validate() (z.ZogIssueMap, error) {
	if c.ConfirmPassword != c.NewPassword {
		return nil, errors.New("new password and confirm password must match")
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

func (d *DeleteUser) Validate() (z.ZogIssueMap, error) {
	errMap := DeleteUserSchema.Validate(d)
	if errMap != nil {
		return errMap, nil
	}
	return nil, nil
}

type ResetPassword struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

func (r *ResetPassword) Validate() (z.ZogIssueMap, error) {
	if r.Password != r.ConfirmPassword {
		return nil, errors.New("password and confirm password must match")
	}
	errMap := ResetPasswordSchema.Validate(r)
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

type UpdateUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func (u *UpdateUser) Validate() (z.ZogIssueMap, error) {
	errMap := UpdateUserSchema.Validate(u)
	if errMap != nil {
		return errMap, nil
	}
	return nil, nil
}

var UpdateUserSchema = z.Struct(z.Schema{
	"firstName": z.String().Required(),
	"lastName":  z.String().Required(),
	"email": z.String().Required().Email().Transform(func(val *string, ctx z.Ctx) error {
		*val = strings.ToLower(*val)
		*val = strings.TrimSpace(*val)
		return nil
	}),
})

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
	"email": z.String().Required().Email().Transform(func(val *string, ctx z.Ctx) error {
=======
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
>>>>>>> e50232b (VALIDATION)
		*val = strings.ToLower(*val)
		*val = strings.TrimSpace(*val)
		return nil
	}),
	"password": z.String().Required().Min(8).Max(32).ContainsSpecial().ContainsUpper().ContainsDigit(),
})
<<<<<<< HEAD

var DeleteUserSchema = z.Struct(z.Schema{
	"password": z.String().Required().Min(8).Max(32).ContainsSpecial().ContainsUpper().ContainsDigit(),
})

var ResetPasswordSchema = z.Struct(z.Schema{
	"password":        z.String().Required().Min(8).Max(32).ContainsSpecial().ContainsUpper().ContainsDigit(),
	"confirmPassword": z.String().Required().Min(8).Max(32).ContainsSpecial().ContainsUpper().ContainsDigit(),
})
=======
>>>>>>> e50232b (VALIDATION)
