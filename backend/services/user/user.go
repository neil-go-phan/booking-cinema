package userservice

import (
	"booking-cinema-backend/entities"
	"encoding/base64"
	"errors"
)

var SALT_SIZE uint8 = 8 // 8 byte

type UserService struct {
	repo UserRepository
}


func NewUserService(r UserRepository) *UserService {
	return &UserService{
		repo: r,
	}
}

func (s *UserService) GetUser(username string) (*entities.User, error) {
	return s.repo.Get(username)
}

func (s *UserService) ListUsers() (*[]entities.User, error) {
	return s.repo.List()
}

func (s *UserService) DeleteUser(username string) (error){
	return s.repo.Delete(username)
}

func (s *UserService) CreateUser(userInput *User) (*entities.User, error) {
	if userInput.Password != userInput.PasswordConfirmation {
		return nil, errors.New("password confirm is not match")
	}
	salt, err := generateRandomSalt(SALT_SIZE)
	if err != nil {
		return nil, errors.New("error when generate salt")
	}
// hash password with salt and argon2
	hashedPassword,err := hashPassword(userInput.Password, salt)
	if err != nil {
		return nil, err
	}
	userInput.Password = hashedPassword
	userInput.Salt = base64.RawStdEncoding.EncodeToString(salt)
	entitiesUser := NewEntitiesUser(userInput)
	return s.repo.Create(entitiesUser)
}

func (s *UserService) VerifyUser(username string, userInput User) (bool, error) {
	if username != userInput.Username {
		return false, errors.New("username is incorrect")
	}
	userFromDB, err := s.GetUser(username)
	if err != nil {
		return false, err
	}
	if userFromDB.Username == "" {
		return false, errors.New("username is incorrect")
	}

	return verifyPassword(userInput.Password, userFromDB.Password)
}

// func (s *userService)UpdateUser(userInput *User) (error) {
// 	// user := NewEntitiesUser(userInput)
// 	// err := s.roleService.Validate(user.Role)
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	// return s.repo.Update(user)
// }
