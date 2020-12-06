package repository

import (
	"github.com/jinzhu/gorm"
	"go-micro-shopping/order/model"
)

type Repository interface {
	Create(*model.Order) error
	Find(orderId string) (*model.Order, error)
	Update(*model.Order, int64) (model.Order, error)
}

type Order struct {
	Repo *gorm.DB
}

func (repo *Order) Create(order *model.Order) error {
	if err := repo.Repo.Create(order).Error; err != nil {
		return err
	}
	return nil
}
func (repo *Order) Find(orderId string) (*model.Order, error) {
	order := &model.Order{}
	if err := repo.Repo.Where("order_id=?", orderId).Find(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}
func (repo *Order) Update(order *model.Order, id uint64) (*model.Order, error) {
	if err := repo.Repo.Model(order).Updates(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}
