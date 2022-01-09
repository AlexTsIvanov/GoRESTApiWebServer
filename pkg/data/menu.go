package data

import (
	"encoding/json"
	"io"

	"github.com/AlexTsIvanov/OrderService/pkg/data/database"
)

func GetMenu(w io.Writer) error {
	db := database.Connector()

	var menu []database.MenuItem
	db.Find(&menu)
	e := json.NewEncoder(w)
	return e.Encode(menu)
}

func GetMenuItem(id int, w io.Writer) error {
	db := database.Connector()

	var menuItem database.MenuItem
	db.First(&menuItem, id)
	e := json.NewEncoder(w)
	return e.Encode(menuItem)
}

func PostMenu(m *database.MenuItem) {
	db := database.Connector()
	db.Create(&m)
}

func UpdateMenu(id int, m *database.MenuItem) error {
	db := database.Connector()
	err := db.Model(&database.MenuItem{}).Where("ID = ?", id).Updates(m).Error
	return err
}

func DeleteMenu(id int) error {
	db := database.Connector()
	err := db.Delete(&database.MenuItem{}, id).Error
	return err
}

func GetMenuTypes(w io.Writer) error {
	db := database.Connector()

	var types []database.TypeItem
	db.Find(&types)
	e := json.NewEncoder(w)
	return e.Encode(types)
}
