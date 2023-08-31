package repository

import (
	"dynamic-segmentation-service/pkg/model"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user model.User) (model.User, error) {
	if err := r.db.Save(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UserRepository) GetUser(id int) (model.User, error) {
	user := model.User{}
	if err := r.db.First(&user, id).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}
