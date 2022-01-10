package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/AlexTsIvanov/OrderService/pkg/data"
	"github.com/AlexTsIvanov/OrderService/pkg/data/database"
	"github.com/gorilla/mux"
)

type Order struct {
	l *log.Logger
}

func NewOrder(l *log.Logger) *Order {
	return &Order{l}
}

func (m *Order) GetOrders(rw http.ResponseWriter, r *http.Request) {
	validUsers := []string{"kitchen", "admin"}

	if CheckPermissions(rw, r, validUsers) {
		//fetch menu from database
		err := data.GetOrders(rw)
		if err != nil {
			http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
		}
		return
	}

}

func (m *Order) GetOrderByID(rw http.ResponseWriter, r *http.Request) {
	validUsers := []string{"kitchen", "admin"}

	if CheckPermissions(rw, r, validUsers) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(rw, "Unable to convert id", http.StatusBadRequest)
			return
		}

		err = data.GetOrderByID(id, rw)
		if err != nil {
			http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
		}
	}

}

func (m *Order) GetOrderItems(rw http.ResponseWriter, r *http.Request) {
	validUsers := []string{"kitchen", "admin"}

	if CheckPermissions(rw, r, validUsers) {
		err := data.GetOrderItems(rw)
		if err != nil {
			http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
		}
	}
}

func (m *Order) GetOrderItemsByOrderID(rw http.ResponseWriter, r *http.Request) {
	validUsers := []string{"kitchen", "admin"}

	if CheckPermissions(rw, r, validUsers) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(rw, "Unable to convert id", http.StatusBadRequest)
			return
		}
		//fetch menu from database
		err = data.GetOrderItemsByOrderID(id, rw)
		if err != nil {
			http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
		}
	}
}

func (m *Order) GetOrderItemsByCustomerID(rw http.ResponseWriter, r *http.Request) {
	validUsers := []string{"customer", "kitchen", "admin"}

	if CheckPermissions(rw, r, validUsers) {

		id, err := strconv.ParseUint(r.Header["Userid"][0], 10, 32)
		if err != nil {
			http.Error(rw, "Unable to convert id", http.StatusBadRequest)
			return
		}
		//fetch menu from database
		err = data.GetOrderItemsByCustomerID(uint(id), rw)
		if err != nil {
			http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
		}
	}
}

func (m *Order) PutOrderItemIDByCustomerID(rw http.ResponseWriter, r *http.Request) {
	validUsers := []string{"customer", "admin"}

	if CheckPermissions(rw, r, validUsers) {
		vars := mux.Vars(r)
		id, _ := strconv.ParseUint(r.Header["Userid"][0], 10, 32)
		itemId, err := strconv.Atoi(vars["itemId"])
		if err != nil {
			http.Error(rw, "Unable to convert id", http.StatusBadRequest)
			return
		}
		orderItem := &database.OrderItem{}
		err = database.FromJSON(r.Body, orderItem)
		if err != nil {
			http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
		}

		err = data.UpdateOrderItemIDByCustomerID(uint(id), uint(itemId), orderItem)
		if err != nil {
			http.Error(rw, "OrderItem not found", http.StatusInternalServerError)
			return
		}
	}
}

func (p *Order) PostOrders(rw http.ResponseWriter, r *http.Request) {
	validUsers := []string{"customer", "admin"}

	if CheckPermissions(rw, r, validUsers) {
		prod := &database.Order{}
		fmt.Println(r.Header["Userid"])
		id, _ := strconv.ParseUint(r.Header["Userid"][0], 10, 32)
		prod.UserID = uint(id)
		prod.StatusID = 1
		err := database.FromJSON(r.Body, prod)
		if err != nil {
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		}

		data.PostOrders(prod)
	}
}

func (p *Order) PostOrderItems(rw http.ResponseWriter, r *http.Request) {
	validUsers := []string{"customer", "admin"}

	if CheckPermissions(rw, r, validUsers) {
		prod := &database.OrderItem{}

		id, _ := strconv.ParseUint(r.Header["Userid"][0], 10, 32)
		err := database.FromJSON(r.Body, prod)
		if err != nil {
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		}

		data.PostOrderItems(uint(id), prod)
	}
}
