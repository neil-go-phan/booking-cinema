package userservice

import (
	"booking-cinema-backend/entities"
	"booking-cinema-backend/services/role"

	"gorm.io/gorm"
)

type UserReader interface {
	Get(username string) (*entities.User, error)
	List() (*[]entities.User, error)
}

type UserWriter interface {
	Create(userInput *entities.User) (*entities.User, error)
	Delete(username string) error
	// Update(userInput *entities.User) (error)
}

type UserRepository interface {
	UserReader
	UserWriter
}

type UserService interface {
	GetUser(username string) (*entities.User, error)
	ListUsers() (user *[]entities.User, err error)
	CreateUser(userInput *User) (*entities.User, error)
	VerifyUser(username string, userInput User) (bool, error)
	DeleteUser(username string) error
	// UpdateUser(userInput *User) (error)
}

type User struct {
	gorm.Model
	FullName             string
	Username             string `gorm:"unique;not null;index:username_index"`
	Password             string
	PasswordConfirmation string `gorm:"-"`
	Salt                 string
	RoleName             string 
	Role                 roleservice.Role `gorm:"foreignKey:RoleName"`
}
