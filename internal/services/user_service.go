package services

import (
	"github.com/HenRok1/test_task_for_Avito/internal/entity"
	"github.com/HenRok1/test_task_for_Avito/internal/repositories"
)

// type userService interface {
// 	AddUserSegments(userID int, segment_name []string) error
// 	GetActiveUserSegments(userID int) ([]entity.Segment, error)
// }

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) AddUserSegments(userID int, segment_name []string) error {
	// for _, name := range removeSegments {
	// 	err := s.userRepo.RemoveSegmentsFromUser(userID, []string{name})
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	name := entity.UserSegment{
		UserID:      userID,
		SegmentName: segment_name,
	}

	return s.userRepo.AddSementsToUser(userID, name)
}

func (s *UserService) RemoveUserSegments(userID int, removeSegments []string) error {
	return s.userRepo.RemoveSegmentsFromUser(userID, removeSegments)
}
