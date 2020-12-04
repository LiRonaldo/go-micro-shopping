package repository

import (
	"github.com/jinzhu/gorm"
	"go-micro-shopping/user/model"
)

type Repository interface {
	Find(id uint32) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) (*model.User, error)
	FindByField(key string, value string, field string) (*model.User, error)
}
type User struct {
	Dao *gorm.DB
}

func (repo *User) Find(id uint32) (*model.User, error) {
	user := &model.User{}
	user.ID = uint(id)
	if err := repo.Dao.First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func (repo *User) Create(user *model.User) error {
	if err := repo.Dao.Create(user).Error; err != nil {
		return err
	}
	return nil
}
func (repo *User) Update(user *model.User) (*model.User, error) {
	//updates 方法参数是实体，不是指针。所以要加&，
	if err := repo.Dao.Model(user).Updates(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func (repo *User) FindByField(key string, value string, fields string) (*model.User, error) {
	if len(fields) == 0 {
		fields = "*"
	}
	user := &model.User{}
	if err := repo.Dao.Select(fields).Where(key+" = ?", value).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
