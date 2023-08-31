package repository

import (
	"dynamic-segmentation-service/pkg/model"
	"github.com/jinzhu/gorm"
)

type UserSegmentsRepository struct {
	db *gorm.DB
}

func NewUserSegmentsRepository(db *gorm.DB) *UserSegmentsRepository {
	return &UserSegmentsRepository{db: db}
}

func (r *UserSegmentsRepository) GetUserSegments(user model.User) (model.User, error) {
	if err := r.db.Model(user).Preload("Segments").Find(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UserSegmentsRepository) UpdateUserSegments(user model.User) (model.User, error) {
	if err := r.db.Model(user).Association("Segments").Replace(user.Segments).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}
