package handlers

import (
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
	for _, s := range r.Header["Role"] {
		if s == "kitchen" {
			//fetch menu from database
			err := data.GetOrders(rw)
			if err != nil {
				http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
			}
			return
		}
	}
	http.Error(rw, "Unauthorized to view", http.StatusBadRequest)

}

func (m *Order) GetOrderByID(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}
	//fetch menu from database
	err = data.GetOrderByID(id, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}

}

func (m *Order) GetOrderItems(rw http.ResponseWriter, r *http.Request) {

	//fetch menu from database
	err := data.GetOrderItems(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}

}

func (m *Order) GetOrderItemsByOrderID(rw http.ResponseWriter, r *http.Request) {
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

func (p *Order) PostOrders(rw http.ResponseWriter, r *http.Request) {
	prod := &database.Order{}

	err := database.FromJSON(r.Body, prod)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.PostOrders(prod)
}

func (p *Order) PostOrderItems(rw http.ResponseWriter, r *http.Request) {
	prod := &database.OrderItem{}

	err := database.FromJSON(r.Body, prod)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.PostOrderItems(prod)
}
