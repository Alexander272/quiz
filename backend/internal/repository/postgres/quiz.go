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
	Get(context.Context, *models.GetQuizzesDTO) ([]*models.Quiz, error)
	GetByAuthor(context.Context, string) ([]*models.Quiz, error)
	GetById(context.Context, *models.GetQuizDTO) (*models.Quiz, error)
	Create(context.Context, *models.QuizDTO) (string, error)
	Update(context.Context, *models.QuizDTO) error
	Delete(context.Context, *models.DeleteQuizDTO) error
}

func (r *QuizRepo) Get(ctx context.Context, req *models.GetQuizzesDTO) ([]*models.Quiz, error) {
	query := fmt.Sprintf(`SELECT q.id, title, description, image, number_of_attempts, category_id, time, author_id, s.id AS schedule_id
		FROM %s AS q INNER JOIN %s AS s ON q.id=s.quiz_id WHERE s.start_time<=$1 AND s.end_time>=$1`,
		QuizTable, ScheduleTable,
	)

	data := []*models.Quiz{}
	if err := r.db.SelectContext(ctx, &data, query, req.Time); err != nil {
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}
	return data, nil
}

func (r *QuizRepo) GetByAuthor(ctx context.Context, authorId string) ([]*models.Quiz, error) {
	query := fmt.Sprintf(`SELECT id, title, description, image, number_of_attempts, category_id, start_time, end_time, time, author_id
		FROM %s WHERE author_id=$1`,
		QuizTable,
	)

	data := []*models.Quiz{}
	if err := r.db.SelectContext(ctx, &data, query, authorId); err != nil {
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}
	return data, nil
}

func (r *QuizRepo) GetById(ctx context.Context, req *models.GetQuizDTO) (*models.Quiz, error) {
	query := fmt.Sprintf(`SELECT q.id, title, description, image, is_drawing, category_id, has_shuffle, has_skippable, show_list, show_answers, 
		show_results, time, author_id,
		COALESCE(s.id::text,'') AS schedule_id, COALESCE(s.start_time,0) AS start_time, 
		COALESCE(s.end_time,0) AS end_time, COALESCE(s.number_of_attempts,0) AS  number_of_attempts
		FROM %s AS q LEFT JOIN %s AS s ON q.id=s.quiz_id
		WHERE q.id=$1`,
		QuizTable, ScheduleTable,
	)

	quiz := &models.Quiz{}
	if err := r.db.GetContext(ctx, quiz, query, req.ID); err != nil {
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
	// dto.ID = uuid.New().String()
	if dto.CategoryID == "" {
		dto.CategoryID = uuid.Nil.String()
	}

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
	if dto.CategoryID == "" {
		dto.CategoryID = uuid.Nil.String()
	}

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
