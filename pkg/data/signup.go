package data

import (
	"errors"

	"github.com/AlexTsIvanov/OrderService/pkg/data/database"
	"golang.org/x/crypto/bcrypt"
)

func EmailCheck(user database.User) error {
	db := database.Connector()

	var dbuser database.User
	db.Where("email = ?", user.Email).First(&dbuser)

	//checks if email is already register or not
	if dbuser.Email != "" {
		return errors.New("Email already in use")
	}
	return nil
}

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CreateUser(user *database.User) {
	db := database.Connector()
	db.Create(&user)
}
