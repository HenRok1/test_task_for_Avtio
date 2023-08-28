package repositories

import (
	"database/sql"

	"github.com/HenRok1/test_task_for_Avito/internal/entity"
)

type userRepository interface {
	AddSementsToUser(userID int, segments []string) error
	RemoveSegmentsFromUser(userID int, segments []string) error
	GetActiveSegments(userID int) ([]entity.Segment, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
