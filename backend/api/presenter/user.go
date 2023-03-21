package presenter

type User struct {
	FullName             string    `json:"full_name"`
	Username             string    `json:"username" validate:"required,min=8,max=16"`
	Password             string    `json:"password" validate:"required"`
	PasswordConfirmation string    `json:"password_confirmation"`
	RoleName                 string    `json:"role_name"`
}
