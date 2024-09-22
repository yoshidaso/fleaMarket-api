package repository

import "gin-freemarket/models"

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
}
type ItemRepository struct {
	items []models.Item
}

func NewItemMemoryRpository(items []models.Item) IItemRepository {
	return &ItemRepository{items: items}
}

func (r *ItemRepository) FindAll() (*[]models.Item, error) {
	return &r.items, nil
}
