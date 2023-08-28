package services

import (
	"github.com/HenRok1/test_task_for_Avito/internal/entity"
	"github.com/HenRok1/test_task_for_Avito/internal/repositories"
)

type userService interface {
	AddUserSegments(userID int, addSegments, removeSegments []string) error
	GetActiveUserSegments(userID int) ([]entity.Segment, error)
}

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) AddUserSegments(userID int, addSegments, removeSegments []string) error {
	// for _, name := range removeSegments {
	// 	err := s.userRepo.RemoveSegmentsFromUser(userID, []string{name})
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	return s.userRepo.AddSementsToUser(userID, addSegments)
}
