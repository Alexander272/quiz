package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Alexander272/quiz/backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type QuestionRepo struct {
	db *sqlx.DB
}

func NewQuestionRepo(db *sqlx.DB) *QuestionRepo {
	return &QuestionRepo{
		db: db,
	}
}

type Question interface {
	Get(context.Context, *models.GetQuestionsDTO) ([]*models.Question, error)
	GetById(context.Context, *models.GetQuestionDTO) (*models.Question, error)
	Create(context.Context, *models.QuestionDTO) (string, error)
	Update(context.Context, *models.QuestionDTO) error
	Delete(context.Context, *models.DeleteQuestionDTO) error
}

func (r *QuestionRepo) Get(ctx context.Context, req *models.GetQuestionsDTO) ([]*models.Question, error) {
	query := fmt.Sprintf(`SELECT id, number, quiz_id, text, description, image, has_shuffle, level, points, time
		FROM %s WHERE quiz_id=$1 ORDER BY number`,
		QuestionTable,
	)
	data := []*models.Question{}

	if err := r.db.SelectContext(ctx, &data, query, req.QuizID); err != nil {
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}
	return data, nil
}

func (r *QuestionRepo) GetById(ctx context.Context, req *models.GetQuestionDTO) (*models.Question, error) {
	query := fmt.Sprintf(`SELECT id, number, quiz_id, text, description, image, has_shuffle, level, points, time
		FROM %s WHERE id=$1`,
		QuestionTable,
	)
	data := &models.Question{}

	if err := r.db.GetContext(ctx, data, query, req.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRows
		}
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}
	return data, nil
}

func (r *QuestionRepo) Create(ctx context.Context, dto *models.QuestionDTO) (string, error) {
	query := fmt.Sprintf(`INSERT INTO %s (id, number, quiz_id, text, description, image, has_shuffle, level, points, time) 
		VALUES (:id, :number, :quiz_id, :text, :description, :image, :has_shuffle, :level, :points, :time)`,
		QuestionTable,
	)
	// dto.ID = uuid.New().String()

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return "", fmt.Errorf("failed to execute query. error: %w", err)
	}
	return dto.ID, nil
}

func (r *QuestionRepo) Update(ctx context.Context, dto *models.QuestionDTO) error {
	query := fmt.Sprintf(`UPDATE %s SET number=:number, text=:text, description=:description, image=:image, has_shuffle=:has_shuffle, level=:level, 
		points=:points WHERE id=:id`,
		QuestionTable,
	)

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}

func (r *QuestionRepo) Delete(ctx context.Context, dto *models.DeleteQuestionDTO) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=:id`, QuestionTable)

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}
