package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/AlexTsIvanov/OrderService/pkg/data"
	"github.com/AlexTsIvanov/OrderService/pkg/data/database"
	"github.com/gorilla/mux"
)

type Menu struct {
	l *log.Logger
}

func NewMenu(l *log.Logger) *Menu {
	return &Menu{l: l}
}

// GetMenu returns the menu list from the database
func (m *Menu) GetMenu(rw http.ResponseWriter, r *http.Request) {

	//fetch menu from database
	err := data.GetMenu(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

func (m *Menu) GetMenuItem(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}
	//fetch menu from database
	err = data.GetMenuItem(id, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}

}

func (m *Menu) PostMenu(rw http.ResponseWriter, r *http.Request) {
	validUsers := []string{"admin"}

	if CheckPermissions(rw, r, validUsers) {
		prod := &database.MenuItem{}

		err := database.FromJSON(r.Body, prod)
		if err != nil {
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		}

		data.PostMenu(prod)
	}
}

func (m *Menu) UpdateMenu(rw http.ResponseWriter, r *http.Request) {
	validUsers := []string{"admin"}

	if CheckPermissions(rw, r, validUsers) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(rw, "Unable to convert id", http.StatusBadRequest)
			return
		}
		prod := &database.MenuItem{}
		err = database.FromJSON(r.Body, prod)
		if err != nil {
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		}

		err = data.UpdateMenu(id, prod)
		if err != nil {
			http.Error(rw, "MenuItem not found", http.StatusInternalServerError)
			return
		}
	}
}

func (m *Menu) DeleteMenu(rw http.ResponseWriter, r *http.Request) {
	validUsers := []string{"admin"}

	if CheckPermissions(rw, r, validUsers) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(rw, "Unable to convert id", http.StatusBadRequest)
			return
		}

		err = data.DeleteMenu(id)
		if err != nil {
			http.Error(rw, "Error deleting entry", http.StatusInternalServerError)
			return
		}
	}
}

func (m *Menu) GetMenuTypes(rw http.ResponseWriter, r *http.Request) {

	err := data.GetMenuTypes(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}
