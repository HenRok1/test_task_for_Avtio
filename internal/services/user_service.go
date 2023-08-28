package services

import (
	"github.com/HenRok1/test_task_for_Avito/internal/repositories"
)

// type userService interface {
// 	AddUserSegment(userID int, addSegment, removeSegments []string) error
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
