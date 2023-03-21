package roleservice

import (
	"booking-cinema-backend/entities"

	"gorm.io/gorm"
)

type RoleReader interface {
	Get(roleName string) (r *entities.User, err error)
	List() (role *[]entities.Role,err error)
}

type RoleWriter interface {
	// Create(userInput *entities.User) (*entities.User, error)
}

type RoleRepository interface {
	RoleReader
	RoleWriter
}

type RoleService interface {
	ListRole() (role *[]entities.Role, err error) 
	Validate(roleName string) (err error)
	GetRole(roleName string) (*entities.Role, error)
	// CreateUser(userInput *User) (*entities.User, error) 
	// VerifyUser(username string, userInput User) (bool, error)
	// UpdateUser(e *entity.User) error
	// DeleteUser(id entity.ID) error
}

type Role struct {
	gorm.Model
	RoleDetailID uint
	RoleName string `gorm:"unique;not null;index:rolename_index"`
}