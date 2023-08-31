package service

import (
	"dynamic-segmentation-service/pkg/model"
	"dynamic-segmentation-service/pkg/repository"
)

type UserSegmentsService struct {
	repo           repository.UserSegments
	userService    User
	segmentService Segment
}

func NewUserSegmentsService(repo repository.UserSegments, userService User, segmentService Segment) *UserSegmentsService {
	return &UserSegmentsService{repo: repo, userService: userService, segmentService: segmentService}
}

func (s *UserSegmentsService) GetUserSegments(userId int) (model.User, error) {
	user := model.User{}
	user.Id = userId
	return s.repo.GetUserSegments(user)
}

func (s *UserSegmentsService) UpdateUserSegments(userSegments model.UserSegments) (model.User, error) {
	addSlugs := userSegments.AddSlugs
	removeSlugs := userSegments.RemoveSlugs
	userId := userSegments.UserId

	user, err := s.userService.GetUser(userId)
	if err != nil {
		return model.User{}, err
	}

	addSegments, err := s.segmentService.GetSegmentsByTitles(addSlugs)
	if err != nil {
		return model.User{}, err
	}

	segments := append(user.Segments, addSegments...)

	user.Segments = filterSegments(segments, removeSlugs)

	return s.repo.UpdateUserSegments(user)
}

// Метод для удаления сегментов из списка removeSlugs
func filterSegments(savedSegments []model.Segment, removeSlugs []string) []model.Segment {
	filteredSegments := make([]model.Segment, 0)

	slugsToRemoveSet := make(map[string]bool)
	for _, slug := range removeSlugs {
		slugsToRemoveSet[slug] = true
	}

	for _, segment := range savedSegments {
		if !slugsToRemoveSet[segment.Slug] {
			filteredSegments = append(filteredSegments, segment)
		}
	}

	return filteredSegments
}
