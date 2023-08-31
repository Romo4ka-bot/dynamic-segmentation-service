package service

import (
	"dynamic-segmentation-service/pkg/model"
	"dynamic-segmentation-service/pkg/repository"
)

type User interface {
	CreateUser(user model.User) (model.User, error)
	GetUser(id int) (model.User, error)
}

type Segment interface {
	CreateSegment(segment model.Segment) (model.Segment, error)
	DeleteSegment(id int) error
	GetSegmentsByTitles(titles []string) ([]model.Segment, error)
}

type UserSegments interface {
	UpdateUserSegments(userSegments model.UserSegments) (model.User, error)
	GetUserSegments(userId int) (model.User, error)
}

type Service struct {
	User
	Segment
	UserSegments
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:         NewUserService(repo),
		Segment:      NewSegmentService(repo),
		UserSegments: NewUserSegmentsService(repo, NewUserService(repo), NewSegmentService(repo)),
	}
}
