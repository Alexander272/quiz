package postgres

import (
	"context"
	"fmt"

	"github.com/Alexander272/quiz/backend/internal/models"
	"github.com/Alexander272/quiz/backend/internal/repository/postgres/pq_models"
	"github.com/Alexander272/quiz/backend/internal/utils"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type AttemptDetailsRepo struct {
	db *sqlx.DB
}

func NewAttemptDetailsRepo(db *sqlx.DB) *AttemptDetailsRepo {
	return &AttemptDetailsRepo{
		db: db,
	}
}

type AttemptDetails interface {
	Get(ctx context.Context, req *models.GetAttemptDetails) (*models.AttemptDetails, error)
	CreateAll(ctx context.Context, dto *models.CreateAttemptDetailsDTO) error
	Create(ctx context.Context, dto *models.AttemptDetailDTO) error
	Update(ctx context.Context, dto *models.AttemptDetailDTO) error
}

func (r *AttemptDetailsRepo) Get(ctx context.Context, req *models.GetAttemptDetails) (*models.AttemptDetails, error) {
	/*
		SELECT id, attempt_id, a.question_id, d.question_id, answers, a.correct
		FROM public.attempt_details AS d
		LEFT JOIN LATERAL (SELECT question_id, array_agg(id) AS correct FROM answer WHERE is_correct=true
			GROUP BY question_id) AS a ON a.question_id=d.question_id
		WHERE attempt_id=''
	*/
	query := fmt.Sprintf(`SELECT d.id, attempt_id, d.question_id, q.points, answers, a.correct FROM %s AS d
		LEFT JOIN LATERAL (SELECT question_id, array_agg(id) AS correct FROM %s WHERE is_correct=true
			GROUP BY question_id) AS a ON a.question_id=d.question_id
		LEFT JOIN LATERAL (SELECT id, points FROM %s) AS q ON a.question_id=q.id
		WHERE attempt_id=$1`,
		AttemptDetailsTable, AnswerTable, QuestionTable,
	)

	tmp := []*pq_models.AttemptDetails{}
	if err := r.db.SelectContext(ctx, &tmp, query, req.AttemptID); err != nil {
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}

	data := &models.AttemptDetails{AttemptID: req.AttemptID}
	for _, d := range tmp {
		isCorrect := utils.AreSlicesEqual(d.Answers, d.Correct)
		correct := []string{}
		if req.ShowAnswers {
			correct = d.Correct
		}

		question := &models.AttemptQuestion{
			ID:        d.QuestionID,
			Answers:   d.Answers,
			Points:    d.Points,
			IsCorrect: req.ShowAnswers && isCorrect,
			Correct:   correct,
		}

		data.Questions = append(data.Questions, question)
	}
	return data, nil
}

func (r *AttemptDetailsRepo) CreateAll(ctx context.Context, dto *models.CreateAttemptDetailsDTO) error {
	query := fmt.Sprintf(`INSERT INTO %s (id, attempt_id, question_id, answers) 
		VALUES (:id, :attempt_id, :question_id, :answers)`,
		AttemptDetailsTable,
	)

	tmp := []*pq_models.AttemptDetailsDTO{}
	for _, q := range dto.Questions {
		tmp = append(tmp, &pq_models.AttemptDetailsDTO{
			ID:         uuid.NewString(),
			AttemptID:  dto.AttemptID,
			QuestionID: q.ID,
			Answers:    pq.StringArray(q.Answers),
		})
	}

	if _, err := r.db.NamedExecContext(ctx, query, tmp); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}

func (r *AttemptDetailsRepo) Create(ctx context.Context, dto *models.AttemptDetailDTO) error {
	query := fmt.Sprintf(`INSERT INTO %s (id, attempt_id, question_id, answers) 
		VALUES (:id, :attempt_id, :question_id, :answers)`,
		AttemptDetailsTable,
	)
	dto.ID = uuid.NewString()

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}

func (r *AttemptDetailsRepo) Update(ctx context.Context, dto *models.AttemptDetailDTO) error {
	query := fmt.Sprintf(`UPDATE %s SET answers=:answers WHERE id=:id`, AttemptDetailsTable)

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}
