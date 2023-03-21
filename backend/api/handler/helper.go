package handler

import (
	"booking-cinema-backend/api/presenter"
	userservice "booking-cinema-backend/services/user"
	"crypto/rand"
	"errors"
	"regexp"
	"time"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v5"
)

var validate *validator.Validate
var TOKEN_SERECT_KEY = []byte("GolenOwl2023")

// TODO: create a env.example contain all const
const ACCESS_TOKEN_LIFE = 5 * time.Minute // 5 min
const REFRESH_TOKEN_LIFE = 24 * time.Hour // 1 day
const RANDOM_TOKEN_STRING_SIZE = 8

func newServicesUser(userInput *presenter.User) *userservice.User {
	user := &userservice.User{
		FullName:             userInput.FullName,
		Username:             userInput.Username,
		Password:             userInput.Password,
		PasswordConfirmation: userInput.PasswordConfirmation,
		RoleName:                 "user",
	}
	return user
}

func validateSignUp(user *presenter.User) error {
	err := validateFullname(user)
	if err != nil {
		return err
	}
	err = validateUsernameAndPassword(user)
	if err != nil {
		return err
	}
	return nil
}

func validateFullname(user *presenter.User) error {
	validate = validator.New()
	err := validate.Var(user.FullName, "required,min=8,max=50")
	if err != nil {
		return err
	}
	checkRegexFullName := checkRegexp(user.FullName, "full_name")
	if !checkRegexFullName {
		return errors.New("full name must not contain special character")
	}
	return nil
}

func validateUsernameAndPassword(user *presenter.User) error {
	validate = validator.New()
	match := checkRegexp(user.Password, "username_password")
	if !match {
		return errors.New("password must not contain special character")
	}
	match = checkRegexp(user.Username, "username_password")
	if !match {
		return errors.New("username must not contain special character")
	}
	err := validate.Struct(user)
	if err != nil {
		return err
	}
	return nil
}

func checkRegexp(checkedString string, checkType string) bool {
	switch checkType {
	case "username_password":
		match, _ := regexp.MatchString("^[a-zA-Z0-9_]*$", checkedString)
		return match
	case "full_name":
		match, _ := regexp.MatchString("^[a-zA-Z0-9_ ]*$", checkedString)
		return match
	}
	return false
}

func generateAccessToken(username, role string) (string, error) {
	expirationTime := time.Now().Add(ACCESS_TOKEN_LIFE)

	claims := &presenter.JWTClaim{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(TOKEN_SERECT_KEY)
}

func GenerateRefreshToken(username, role string) (string, error) {
	randomString, err := generateRandomTokenString()
	if err != nil {
		return "", err
	}
	expirationTime := time.Now().Add(REFRESH_TOKEN_LIFE)
	claims := &presenter.JWTClaim{
		Username:     username,
		Role:         role,
		RandomString: randomString,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(TOKEN_SERECT_KEY)
}

func generateRandomTokenString() ([]byte, error) {
	var randomString = make([]byte, RANDOM_TOKEN_STRING_SIZE)

	_, err := rand.Read(randomString[:])

	if err != nil {
		return nil, err
	}

	return randomString, nil
}
