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

func (repo *UserRepository) AddSementsToUser(userID int, segments []string) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	for _, name := range segments {
		query := "INSERT INTO user_segments (user_id, segment_name) VALUES ($1, $2)"
		_, err := tx.Exec(query, userID, name)
		if err != nil {
			return err
		}
	}
	return nil
}

func (repo *UserRepository) RemoveSegmentsFromUser(userID int, segments []string) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	for _, name := range segments {
		query := "DELETE FROM user_segments WHERE user_id = $1 AND segment_name = $2"
		_, err := tx.Exec(query, userID, name)
		if err != nil {
			return err
		}
	}
	return nil

}
