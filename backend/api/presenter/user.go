package presenter

type User struct {
	FullName             string    `json:"fullName"`
	Username             string    `json:"username" validate:"required,min=8,max=16"`
	Password             string    `json:"password" validate:"required"`
	PasswordConfirmation string    `json:"passwordConfirmation"`
	RoleName                 string    `json:"roleName"`
}
