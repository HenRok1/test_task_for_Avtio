package repositories

import (
	"github.com/HenRok1/test_task_for_Avito/internal/entity"
	"gorm.io/gorm"
)

type SegmentRepository interface {
	CreateSegment(segment *entity.Segment) (*entity.Segment, error)
}

type DefaultSegmentRepository struct {
	DB *gorm.DB
}

func (r *DefaultSegmentRepository) CreateSegment(segment *entity.Segment) (*entity.Segment, error) {
	if err := r.DB.Create(segment).Error; err != nil {
		return nil, err
	}

	return segment, nil
}
