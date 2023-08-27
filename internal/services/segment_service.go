package services

import (
	"github.com/HenRok1/test_task_for_Avito/internal/entity"
	"github.com/HenRok1/test_task_for_Avito/internal/repositories"
)

type SegmentService interface {
	CreateSegment(name string) (*entity.Segment, error)
}

type DefaultSegmentService struct {
	SegmentRepository repositories.SegmentRepository
}

func (s *DefaultSegmentService) CreateSegment(name string) (*entity.Segment, error) {
	segment := &entity.Segment{
		Name: name,
	}

	return s.SegmentRepository.CreateSegment(segment)
}
