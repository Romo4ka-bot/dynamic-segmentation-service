package service

import (
	"dynamic-segmentation-service/pkg/model"
	"dynamic-segmentation-service/pkg/repository"
)

type SegmentService struct {
	repo repository.Segment
}

func NewSegmentService(repo repository.Segment) *SegmentService {
	return &SegmentService{repo: repo}
}

func (s *SegmentService) CreateSegment(segment model.Segment) (model.Segment, error) {
	return s.repo.CreateSegment(segment)
}

func (s *SegmentService) DeleteSegment(id int) error {
	segment, err := s.repo.GetSegment(id)
	if err != nil {
		return err
	}
	return s.repo.DeleteSegment(segment)
}

func (s *SegmentService) GetSegmentsByTitles(titles []string) ([]model.Segment, error) {
	return s.repo.GetSegmentsByTitles(titles)
}
