package repository

import (
	"dynamic-segmentation-service/pkg/model"
	"github.com/jinzhu/gorm"
)

type User interface {
	CreateUser(user model.User) (model.User, error)
	GetUser(id int) (model.User, error)
}

type Segment interface {
	CreateSegment(segment model.Segment) (model.Segment, error)
	DeleteSegment(segment model.Segment) error
	GetSegment(id int) (model.Segment, error)
	GetSegmentsByTitles(titles []string) ([]model.Segment, error)
}

type UserSegments interface {
	UpdateUserSegments(userSegments model.User) (model.User, error)
	GetUserSegments(user model.User) (model.User, error)
}

type Repository struct {
	User
	Segment
	UserSegments
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:         NewUserRepository(db),
		Segment:      NewSegmentRepository(db),
		UserSegments: NewUserSegmentsRepository(db),
	}
}
