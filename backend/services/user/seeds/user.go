package userseeds

import (
	userservice "booking-cinema-backend/services/user"
	"fmt"
	"log"
	"math/rand"
)

func CreateAdminUser(userService *userservice.UserService) error {
	user := &userservice.User{
		FullName:             "Super admin",
		Username:             "superadmin",
		Password:             PASSWORD_DEFAULT,
		PasswordConfirmation: PASSWORD_DEFAULT,
		RoleName:             "super admin",
	}
	_, err := userService.CreateUser(user)
	if err != nil {
		return err
	}
	log.Println("add a super admin user successfull")
	return nil
}
func CreateUser(userService *userservice.UserService) error {
	randomUser := generateRandomUserService()
	_, err := userService.CreateUser(randomUser)
	if err != nil {
		return err
	}
	log.Println("add a new seed user successfull")
	return nil
}

var PASSWORD_DEFAULT = "259ad599004d7dca7fa7bb5f95bba5698a8948b5ec93c3d121c5d7645567bc1edf5f06fb9ea939d01cd55be29196f4e6f3e3cc3aa04e5ae1709daf59bcbc4b06" //'goldenowl2023' hash by SHA512

func generateRandomUserService() *userservice.User {
	user := &userservice.User{
		FullName:             randomName(),
		Username:             randomUsername(),
		Password:             PASSWORD_DEFAULT,
		PasswordConfirmation: PASSWORD_DEFAULT,
		RoleName:             "user",
	}
	return user
}

var lastNames []string = []string{"Nguyen", "Phan", "Le", "Tran", "Pham", "Huynh", "Hoang", "Vo", "Truong", "Bui", "Ly", "Do"}

var middleNames []string = []string{"Anh", "Bach", "Duc", "Cao", "Hai", "Cuong", "Gia", "Trung", "Xuan", "Ngoc", "Tuan", "Tien"}

var firstNames []string = []string{"Quyen", "Phuc", "Hung", "Dung", "Quynh", "Vu", "Anh", "Truc", "Trinh", "Bao", "Toan", "Duy", "Thang", "Viet"}

func randomPickStr(arr []string) string {
	return arr[rand.Intn(len(arr))]
}

func randomName() string {
	return randomPickStr(lastNames) + " " + randomPickStr(middleNames) + " " + randomPickStr(firstNames)
}

var usernames []string = []string{"quyen", "phuc", "hung", "dung", "quynh", "vu", "anh", "truc", "trinh", "bao", "toan", "duy", "thang", "viet"}

func randomInt() int {
	min := 10000000
	max := 99999999
	return rand.Intn(max-min+1) + min
}

func randomUsername() string {
	return fmt.Sprintf("%s%v", randomPickStr(usernames), randomInt())
}
