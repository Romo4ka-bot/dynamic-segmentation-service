package repository

import (
	"dynamic-segmentation-service/pkg/model"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type SegmentRepository struct {
	db *gorm.DB
}

func NewSegmentRepository(db *gorm.DB) *SegmentRepository {
	return &SegmentRepository{db: db}
}

func (r *SegmentRepository) CreateSegment(segment model.Segment) (model.Segment, error) {
	if err := r.db.Save(&segment).Error; err != nil {
		return model.Segment{}, err
	}
	return segment, nil
}

func (r *SegmentRepository) DeleteSegment(segment model.Segment) error {
	err := r.db.Delete(&segment).Error
	return err
}

func (r *SegmentRepository) GetSegment(id int) (model.Segment, error) {
	segment := model.Segment{}
	err := r.db.First(&segment, id).Error
	return segment, err
}

func (r *SegmentRepository) GetSegmentsByTitles(titles []string) ([]model.Segment, error) {
	var segments []model.Segment
	err := r.db.Model(segments).Where("slug = ANY(?)", pq.Array(titles)).Find(&segments).Error
	return segments, err
}
