package data

import (
	"encoding/json"
	"fmt"
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

func GetOrderItemsByCustomerID(id uint, w io.Writer) error {
	db := database.Connector()
	var order database.Order
	db.Where("user_id = ? AND status_id != ?", id, "3").Find(&order)
	var orderItems []database.OrderItem
	db.Where("order_id = ?", order.ID).Find(&orderItems)
	e := json.NewEncoder(w)
	return e.Encode(orderItems)
}

func UpdateOrderItemIDByCustomerID(id uint, itemId uint, orderItem *database.OrderItem) error {
	db := database.Connector()
	err := db.Model(&database.OrderItem{}).Where("ID = ? AND order_id != ?", itemId, id).Updates(orderItem).Error
	return err
}

func PostOrders(m *database.Order) {
	db := database.Connector()
	db.Create(&m)
}

func PostOrderItems(id uint, m *database.OrderItem) {
	db := database.Connector()
	var order database.Order
	db.Where("user_id = ? AND status_id = ?", id, 1).First(&order)
	fmt.Println()
	m.OrderID = order.ID
	db.Create(&m)
}
