package data

import (
	"encoding/json"
	"io"

	"github.com/AlexTsIvanov/OrderService/pkg/data/database"
)

func GetOrders(w io.Writer) error {
	db := database.Connector()

	var orders []database.Order
	db.Find(&orders)
	e := json.NewEncoder(w)
	return e.Encode(orders)
}

func GetOrderByID(id int, w io.Writer) error {
	db := database.Connector()

	var order database.Order
	db.First(&order, id)
	e := json.NewEncoder(w)
	return e.Encode(order)
}

func GetOrderItems(w io.Writer) error {
	db := database.Connector()

	var orderItems []database.OrderItem
	db.Find(&orderItems)
	e := json.NewEncoder(w)
	return e.Encode(orderItems)
}

func GetOrderItemsByOrderID(id int, w io.Writer) error {
	db := database.Connector()

	var orderItems []database.OrderItem
	db.Where("order_id = ?", id).Find(&orderItems)
	e := json.NewEncoder(w)
	return e.Encode(orderItems)
}

func PostOrders(m *database.Order) {
	db := database.Connector()
	db.Create(&m)
}

func PostOrderItems(m *database.OrderItem) {
	db := database.Connector()
	db.Create(&m)
}
