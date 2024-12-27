package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Alexander272/quiz/backend/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AttemptRepo struct {
	db *sqlx.DB
}

func NewAttemptRepo(db *sqlx.DB) *AttemptRepo {
	return &AttemptRepo{
		db: db,
	}
}

type Attempt interface {
	Get(context.Context, *models.GetAttempt) ([]*models.Attempt, error)
	GetByQuiz(context.Context, *models.GetAttemptByQuiz) ([]*models.Attempt, error)
	GetByID(context.Context, *models.GetAttemptByID) (*models.Attempt, error)
	Create(context.Context, *models.AttemptDTO) (string, error)
	Update(context.Context, *models.AttemptDTO) error
	Delete(context.Context, *models.DeleteAttemptDTO) error
}

func (r *AttemptRepo) Get(ctx context.Context, req *models.GetAttempt) ([]*models.Attempt, error) {
	parts := []string{}
	params := []interface{}{}
	i := 1
	if req.ScheduleID != "" {
		parts = append(parts, fmt.Sprintf("schedule_id=$%d", i))
		params = append(params, req.ScheduleID)
	}
	if req.UserID != "" {
		parts = append(parts, fmt.Sprintf("user_id=$%d", i))
		params = append(params, req.UserID)
	}
	condition := strings.Join(parts, " AND ")

	query := fmt.Sprintf(`SELECT id, schedule_id, user_id, username, start_time, end_time, correct, total, points, total_points FROM %s
		WHERE %s`,
		AttemptTable, condition,
	)
	data := []*models.Attempt{}

	if err := r.db.SelectContext(ctx, &data, query, params...); err != nil {
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}
	return data, nil
}

func (r *AttemptRepo) GetByQuiz(ctx context.Context, req *models.GetAttemptByQuiz) ([]*models.Attempt, error) {
	query := fmt.Sprintf(`SELECT a.id, schedule_id, user_id, username, a.start_time, a.end_time, correct, total, points, total_points 
		FROM %s AS a INNER JOIN %s AS s ON s.id=a.schedule_id WHERE user_id=$1 AND quiz_id=$2 AND s.start_time<=$3 AND s.end_time>=$3
		ORDER BY a.end_time`,
		AttemptTable, ScheduleTable,
	)
	data := []*models.Attempt{}

	if err := r.db.SelectContext(ctx, &data, query, req.UserID, req.QuizID, req.Time); err != nil {
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}
	return data, nil
}

func (r *AttemptRepo) GetByID(ctx context.Context, req *models.GetAttemptByID) (*models.Attempt, error) {
	query := fmt.Sprintf(`SELECT id, schedule_id, user_id, username, start_time, end_time, correct, total, points, total_points FROM %s
		WHERE id=$1`,
		AttemptTable,
	)
	data := &models.Attempt{}

	if err := r.db.GetContext(ctx, data, query, req.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRows
		}
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}
	return data, nil
}

func (r *AttemptRepo) Create(ctx context.Context, dto *models.AttemptDTO) (string, error) {
	query := fmt.Sprintf(`INSERT INTO %s (id, schedule_id, user_id, username, start_time, total, total_points) 
		VALUES (:id, :schedule_id, :user_id, :username, :start_time, :total, :total_points)`,
		AttemptTable,
	)
	dto.ID = uuid.NewString()

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return "", fmt.Errorf("failed to execute query. error: %w", err)
	}
	return dto.ID, nil
}

func (r *AttemptRepo) Update(ctx context.Context, dto *models.AttemptDTO) error {
	query := fmt.Sprintf(`UPDATE %s SET end_time=:end_time, correct=:correct, points=:points, total_points=:total_points WHERE id=:id`,
		AttemptTable,
	)

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}

func (r *AttemptRepo) Delete(ctx context.Context, dto *models.DeleteAttemptDTO) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=:id`, AttemptTable)

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}
