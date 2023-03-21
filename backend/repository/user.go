package repository

import (
	"booking-cinema-backend/entities"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (repo *UserRepo) Create(userInput *entities.User) (*entities.User, error) {
	err := repo.DB.Create(userInput).Error
	if err != nil {
		return nil, err
	}
	return userInput, nil
}

func (repo *UserRepo) Get(username string) (*entities.User, error) {
	return getUser(username, repo)
}

func getUser(username string, repo *UserRepo) (*entities.User, error) {
	user := new(entities.User)
	err := repo.DB.Select("role_name", "full_name", "username", "password", "salt", "id").Where(map[string]interface{}{"username": username}).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepo) List() (*[]entities.User, error) {
	users := make([]entities.User, 10)
	err := repo.DB.Select("role_name", "full_name", "username", "created_at", "updated_at").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (repo *UserRepo) Delete(username string) error {
	user, err := getUser(username, repo)
	if err != nil {
		return err
	}
	err = repo.DB.Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepo) Update(userInput *entities.User) error {
	err := repo.DB.Model(&userInput).Where("username = ?", userInput.Username).Updates(map[string]interface{}{"full_name": userInput.FullName, "role_name": userInput.RoleName}).Error
	if err != nil {
		return err
	}
	return nil
}
