package postgres

import (
	"context"
	"fmt"

	"github.com/Alexander272/quiz/backend/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ScheduleRepo struct {
	db *sqlx.DB
}

func NewScheduleRepo(db *sqlx.DB) *ScheduleRepo {
	return &ScheduleRepo{
		db: db,
	}
}

type Schedule interface {
	Get(context.Context, *models.GetSchedule) ([]*models.Schedule, error)
	GetByQuiz(context.Context, *models.GetScheduleByQuiz) ([]*models.Schedule, error)
	Create(context.Context, *models.ScheduleDTO) (string, error)
	Update(context.Context, *models.ScheduleDTO) error
	Delete(context.Context, *models.DeleteScheduleDTO) error
}

func (r *ScheduleRepo) Get(ctx context.Context, req *models.GetSchedule) ([]*models.Schedule, error) {
	query := fmt.Sprintf(`SELECT id, quiz_id, start_time, end_time, number_of_attempts FROM %s
		WHERE start_time<=$1 AND end_time>=$1`,
		ScheduleTable,
	)
	data := []*models.Schedule{}

	if err := r.db.SelectContext(ctx, &data, query, req.Time); err != nil {
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}
	return data, nil
}

func (r *ScheduleRepo) GetByQuiz(ctx context.Context, req *models.GetScheduleByQuiz) ([]*models.Schedule, error) {
	query := fmt.Sprintf(`SELECT id, quiz_id, start_time, end_time, number_of_attempts FROM %s WHERE quiz_id=$1`, ScheduleTable)
	data := []*models.Schedule{}

	if err := r.db.SelectContext(ctx, &data, query, req.QuizID); err != nil {
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}
	return data, nil
}

func (r *ScheduleRepo) Create(ctx context.Context, dto *models.ScheduleDTO) (string, error) {
	query := fmt.Sprintf(`INSERT INTO %s (id, quiz_id, start_time, end_time, number_of_attempts) 
		VALUES (:id, :quiz_id, :start_time, :end_time, :number_of_attempts)`,
		ScheduleTable,
	)
	dto.ID = uuid.New().String()

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return "", fmt.Errorf("failed to execute query. error: %w", err)
	}
	return dto.ID, nil
}

func (r *ScheduleRepo) Update(ctx context.Context, dto *models.ScheduleDTO) error {
	query := fmt.Sprintf(`UPDATE %s SET start_time=:start_time, end_time=:end_time, number_of_attempts WHERE id=:id`, ScheduleTable)

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}

func (r *ScheduleRepo) Delete(ctx context.Context, dto *models.DeleteScheduleDTO) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=:id`, ScheduleTable)

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}
