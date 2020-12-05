package repository

import (
	"github.com/jinzhu/gorm"
	"go-micro-shopping/product/model"
)

type Repository interface {
	Find(uint) (*model.Product, error)
	Create(*model.Product) error
	Update(*model.Product, uint32) (*model.Product, error)
	FindByField(string, string, string) (*model.Product, error)
}
type Product struct {
	Repo *gorm.DB
}

func (p *Product) Find(id uint32) (*model.Product, error) {
	product := &model.Product{}
	product.ID = uint(id)
	if err := p.Repo.First(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
func (p *Product) Create(product *model.Product) error {
	if err := p.Repo.Create(product).Error; err != nil {
		return err
	}
	return nil
}
func (p *Product) Update(product *model.Product, id uint32) (*model.Product, error) {
	if err := p.Repo.Model(product).Update(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
func (p *Product) FindByField(key string, value string, fields string) (*model.Product, error) {
	if len(fields) == 0 {
		fields = "*"
	}
	product := &model.Product{}
	if err := p.Repo.Select(fields).Where(key+" = ?", value).First(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
