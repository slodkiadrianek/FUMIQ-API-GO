package schemas

type LoginUser struct {
	Email string `json:"email" binding:"required,email"`
	PasswordBody
}
type RegisterUser struct {
	Email           string `json:"email" opts:"required,email"`
	FirstName       string `json:"firstName" opts:"required"`
	LastName        string `json:"lastName" opts:"required"`
	Password        string `json:"password" opts:"required,min=8,max=32"`
	ConfirmPassword string `json:"confirmPassword" opts:"required,min=8,max=32,confirm=Password"`
}

type PasswordBody struct {
	Password string `json:"password" binding:"required,min=8,max=32"`
}
type UserParam struct {
	UserId string `json:"userId" binding:"required"`
}

type ChangePassword struct {
	PasswordBody
	NewPassword     string `json:"newPassword" binding:"required,min=8,max=32"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,min=8,max=32"`
}

type UpdateUser struct {
	Email     string `json:"email" binding:"required,email"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
}

type Token struct {
	Token string `json:"token" binding:"required"`
}
