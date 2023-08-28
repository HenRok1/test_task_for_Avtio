package services

import (
	"errors"

	"github.com/HenRok1/test_task_for_Avito/internal/entity"
	"github.com/HenRok1/test_task_for_Avito/internal/repositories"
)

// type segmentService interface {
// 	CreateSegment(name string) error
// 	DeleteSegment(name string) error
// }

type SegmentService struct {
	segmentRepo *repositories.SegmentRepository
}

func NewSegmentService(segmentRepo *repositories.SegmentRepository) *SegmentService {
	return &SegmentService{
		segmentRepo: segmentRepo,
	}
}

func (s *SegmentService) CreateSegment(name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}

	segment := entity.Segment{
		Name: name,
	}
	return s.segmentRepo.CreateSegment(segment)
}

func (s *SegmentService) DeleteSegment(name string) error {
	return s.segmentRepo.DeleteSegment(name)
}
