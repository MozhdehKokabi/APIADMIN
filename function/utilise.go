package function

import (
	"APIADMIN/auth"
	"APIADMIN/repository"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func GenerateJWT(username string) (string, error) {
	token, err := auth.GenerateJWT(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func CheckPassword(username, password string) bool {
	var storedPassword string
	err := repository.Db.QueryRow("select password from students where username= $1", username).Scan(&storedPassword)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	return err == nil
}
