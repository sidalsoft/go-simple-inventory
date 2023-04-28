package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"go-simple-inventory/database"
	"go-simple-inventory/models"
)

func GetAllItems() []models.Item {
	var items []models.Item = []models.Item{}

	database.DB.Order("created_at desc").Find(&items)

	return items
}

func GetItemByID(id string) (models.Item, error) {
	var item models.Item

	result := database.DB.First(&item, "id = ?", id)

	if result.RowsAffected == 0 {
		return models.Item{}, errors.New("item not found")
	}

	return item, nil
}

func CreateItem(itemRequest models.ItemRequest) models.Item {
	var newItem models.Item = models.Item{
		ID:        uuid.New().String(),
		Name:      itemRequest.Name,
		Price:     itemRequest.Price,
		Quantity:  itemRequest.Quantity,
		CreatedAt: time.Now(),
	}

	database.DB.Create(&newItem)

	return newItem
}

func UpdateItem(itemRequest models.ItemRequest, id string) (models.Item, error) {
	item, err := GetItemByID(id)

	if err != nil {
		return models.Item{}, err
	}

	item.Name = itemRequest.Name
	item.Price = itemRequest.Price
	item.Quantity = itemRequest.Quantity
	item.UpdatedAt = time.Now()

	database.DB.Save(&item)

	return item, nil
}

func DeleteItem(id string) bool {
	item, err := GetItemByID(id)

	if err != nil {
		return false
	}

	database.DB.Delete(&item)

	return true
}
