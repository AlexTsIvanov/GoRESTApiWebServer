package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func (u *User) IsAuthorized(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] == nil {
			_ = json.NewEncoder(w).Encode("No Token Found")
			return
		}

		var mySigningKey = []byte(os.Getenv("SECRET_KEY"))

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing")
			}
			return mySigningKey, nil
		})

		if err != nil {
			_ = json.NewEncoder(w).Encode("Your Token has been expired")
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "admin" {

				r.Header.Set("Role", "admin")
				handler.ServeHTTP(w, r)
				return

			} else if claims["role"] == "customer" {
				r.Header.Set("Role", "customer")
				r.Header.Set("Userid", fmt.Sprint(claims["id"]))
				handler.ServeHTTP(w, r)
				return
			} else if claims["role"] == "kitchen" {

				r.Header.Set("Role", "kitchen")
				handler.ServeHTTP(w, r)
				return
			}
		}
		_ = json.NewEncoder(w).Encode("Not Authorized")
	})
}

func CheckPermissions(rw http.ResponseWriter, r *http.Request, roles []string) bool {
	for _, role := range roles {
		if role == r.Header["Role"][0] {
			return true
		}
	}
	http.Error(rw, "Unauthorized to view", http.StatusBadRequest)
	return false
}
