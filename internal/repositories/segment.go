package repositories

import (
	"database/sql"

	"github.com/HenRok1/test_task_for_Avito/internal/entity"
)

type segmentRepository interface {
	CreateSegment(segment *entity.Segment) error
	DeleteSegment(name string) error
}

type SegmentRepository struct {
	DB *sql.DB
}

func NewSegmentRepository(db *sql.DB) *SegmentRepository {
	return &SegmentRepository{
		DB: db,
	}
}

func (repo *SegmentRepository) CreateSegment(segment entity.Segment) error {
	query := "INSERT INTO segments (name) VALUES ($1)"
	_, err := repo.DB.Exec(query, segment.Name)
	if err != nil {
		return err
	}
	return nil
}

func (repo *SegmentRepository) DeleteSegment(name string) error {
	query := "DELETE FROM segments WHERE name = $1"
	_, err := repo.DB.Exec(query, name)
	if err != nil {
		return err
	}
	return nil
}

//TODO: сделать проверку для удаления, если Сегмента нет в БД, то не делать удаление
