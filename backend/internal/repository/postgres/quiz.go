package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Alexander272/quiz/backend/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type QuizRepo struct {
	db *sqlx.DB
}

func NewQuizRepo(db *sqlx.DB) *QuizRepo {
	return &QuizRepo{
		db: db,
	}
}

type Quiz interface {
	GetById(context.Context, *models.GetQuizDTO) (*models.Quiz, error)
	Create(context.Context, *models.QuizDTO) (string, error)
	Update(context.Context, *models.QuizDTO) error
	Delete(context.Context, *models.DeleteQuizDTO) error
}

func (r *QuizRepo) GetById(ctx context.Context, req *models.GetQuizDTO) (*models.Quiz, error) {
	query := fmt.Sprintf(`SELECT id, title, description, image, is_drawing, number_of_attempts, category_id, start_time, end_time, has_shuffle, has_skippable,
		show_list, show_answers, show_results, time, author_id FROM %s WHERE id=$1`,
		QuizTable,
	)

	quiz := &models.Quiz{}
	if err := r.db.GetContext(ctx, &quiz, query, req.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRows
		}
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}
	return quiz, nil
}

func (r *QuizRepo) Create(ctx context.Context, dto *models.QuizDTO) (string, error) {
	query := fmt.Sprintf(`INSERT INTO %s (id, title, description, image, is_drawing, number_of_attempts, category_id, start_time, end_time, has_shuffle, 
		has_skippable, show_list, show_answers, show_results, time, author_id) VALUES (:id, :title, :description, :image, :is_drawing, :number_of_attempts,
		:category_id, :start_time, :end_time, :has_shuffle, :has_skippable, :show_list, :show_answers, :show_results, :time, :author_id)`,
		QuizTable,
	)
	dto.ID = uuid.New().String()

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return "", fmt.Errorf("failed to execute query. error: %w", err)
	}
	return dto.ID, nil
}

func (r *QuizRepo) Update(ctx context.Context, dto *models.QuizDTO) error {
	query := fmt.Sprintf(`UPDATE %s SET title=:title, description=:description, image=:image, is_drawing=:is_drawing, number_of_attempts=:number_of_attempts,
		category_id=:category_id, start_time=:start_time, end_time=:end_time, has_shuffle=:has_shuffle, has_skippable=:has_skippable, show_list=:show_list,
		show_answers=:show_answers, show_results=:show_results, time=:time, author_id=:author_id WHERE id=:id`,
		QuizTable,
	)

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}

func (r *QuizRepo) Delete(ctx context.Context, dto *models.DeleteQuizDTO) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=:id`, QuizTable)

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}
