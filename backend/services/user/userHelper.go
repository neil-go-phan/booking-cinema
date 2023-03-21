package userservice

import (
	"booking-cinema-backend/entities"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// Establish the parameters to use for Argon2.
var ARGON2_MEMORY uint32 = 64 * 1024
var ARGON2_TIME uint32 = 3
var ARGON2_THREADS uint8 = 2
var ARGON2_KEYLENGTH uint32 = 32

func NewEntitiesUser(userInput *User) *entities.User {
	user := &entities.User{
		FullName:  userInput.FullName,
		Username:  userInput.Username,
		Password:  userInput.Password,
		Salt:      userInput.Salt,
		RoleName:  userInput.RoleName,
		CreatedAt: userInput.CreatedAt,
		UpdatedAt: userInput.UpdatedAt,
	}
	return user
}

func NewEntitiesRole(roleName string) *entities.Role {
	role := &entities.Role{
		RoleName: roleName,
	}
	return role
}

func generateRandomSalt(saltSize uint8) ([]byte, error) {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])

	if err != nil {
		return nil, err
	}

	return salt, nil
}

func hashPassword(password string, salt []byte) (string, error) {
	hash := argon2.IDKey([]byte(password), salt, ARGON2_TIME, ARGON2_MEMORY, ARGON2_THREADS, ARGON2_KEYLENGTH)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)
	// standard encoded hash representation.
	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, ARGON2_MEMORY, ARGON2_TIME, ARGON2_THREADS, b64Salt, b64Hash)

	return encodedHash, nil
}

func verifyPassword(plain, hash string) (bool, error) {
	hashParts := strings.Split(hash, "$")
	var argon2Const struct {
		memory  uint32
		time    uint32
		threads uint8
	}
	_, err := fmt.Sscanf(hashParts[3], "m=%d,t=%d,p=%d", &argon2Const.memory, &argon2Const.time, &argon2Const.threads)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(hashParts[4])
	if err != nil {
		return false, err
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(hashParts[5])
	if err != nil {
		return false, err
	}

	hashToCompare := argon2.IDKey([]byte(plain), salt, argon2Const.time, argon2Const.memory, argon2Const.threads, uint32(len(decodedHash)))

	return subtle.ConstantTimeCompare(decodedHash, hashToCompare) == 1, nil
}
