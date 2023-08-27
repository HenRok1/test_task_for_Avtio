package repositories

import (
	"database/sql"

	"github.com/HenRok1/test_task_for_Avito/internal/entity"
)

type SegmentRepository interface {
	CreateSegment(segment *entity.Segment) (*entity.Segment, error)
}

type DefaultSegmentRepository struct {
	DB *sql.DB
}

func (r *DefaultSegmentRepository) CreateSegment(segment *entity.Segment) (*entity.Segment, error) {
	query := "INSERT INTO segments (Name) VALUES ($1) RETURNING id"
	err := r.DB.QueryRow(query, segment.Name).Scan(&segment.ID)
	if err != nil {
		return nil, err
	}

	return segment, nil
}
