package data

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"time"

	"github.com/AlexTsIvanov/OrderService/pkg/data/database"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

func CheckUser(authdetails Authentication) (database.User, error) {
	db := database.Connector()

	var authuser database.User
	db.Where("email = ?", authdetails.Email).First(&authuser)
	if authuser.Email == "" {
		return database.User{}, errors.New("Username or password incorrect")
	}
	return authuser, nil
}

func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(email, role string) (string, error) {
	var mySigningKey = []byte(os.Getenv("SECRET_KEY"))

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func EncodeToken(w io.Writer, authuser database.User, validToken string) error {
	var token Token
	token.Email = authuser.Email
	token.Role = authuser.TypeUser
	token.TokenString = validToken
	e := json.NewEncoder(w)
	return e.Encode(token)
}
