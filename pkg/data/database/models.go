package database

import (
	"github.com/jinzhu/gorm"
)

type MenuItem struct {
	gorm.Model
	Title       string `json:"title"`
	TypeItemID  uint
	Description string  `json:"desc"`
	Price       float32 `json:"price" gt:"0"`
	ActiveItem  bool    `json:"activeItem"`
	OrderItems  []OrderItem
}

type TypeItem struct {
	ID       uint   `gorm:"size:255"`
	Name     string `json:"name"`
	MenuItem []MenuItem
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `gorm:"unique"`
	Address  string `gorm:"default: "`
	Number   int    `gorm:"default: "`
	TypeUser string `gorm:"default:customer"`
	Orders   []Order
}

type Status struct {
	ID     uint
	Status string `json:"status"`
	Orders []Order
}

type OrderItem struct {
	gorm.Model
	MenuItemID uint `json:"menu_item_id"`
	Quantity   uint `json:"quantity"`
	OrderID    uint `gorm:"default: "`
}

type Order struct {
	gorm.Model
	OrderItem []OrderItem
	UserID    uint
	StatusID  uint
}
