package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AlexTsIvanov/OrderService/pkg/data"
	"github.com/AlexTsIvanov/OrderService/pkg/data/database"
)

type User struct {
	l *log.Logger
}

func NewUser(l *log.Logger) *User {
	return &User{l}
}

func (u *User) SignUp(rw http.ResponseWriter, r *http.Request) {

	var user database.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(rw, "Unable to read body", http.StatusBadRequest)
		return
	}

	err = data.EmailCheck(user)
	if err != nil {
		http.Error(rw, "Email already exists", http.StatusBadRequest)
		return
	}

	user.Password, err = data.GeneratehashPassword(user.Password)
	if err != nil {
		http.Error(rw, "Error hashing password", http.StatusInternalServerError)
		return
	}
	//insert user details in database
	data.CreateUser(&user)
}

func (u *User) SignIn(rw http.ResponseWriter, r *http.Request) {

	var authdetails data.Authentication
	err := json.NewDecoder(r.Body).Decode(&authdetails)
	if err != nil {
		http.Error(rw, "Unable to read body", http.StatusBadRequest)
		return
	}

	authuser, err := data.CheckUser(authdetails)
	if err != nil {
		http.Error(rw, "Username or password incorrect", http.StatusBadRequest)
		return
	}

	check := data.CheckPasswordHash(authuser.Password, authdetails.Password)
	if !check {
		http.Error(rw, "Username or password incorrect", http.StatusBadRequest)
		return
	}

	validToken, err := data.GenerateJWT(authuser.Email, authuser.TypeUser)
	if err != nil {
		http.Error(rw, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	err = data.EncodeToken(rw, authuser, validToken)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}
