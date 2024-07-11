package postgres

import (
	"context"
	"fmt"

	"github.com/Alexander272/quiz/backend/internal/models"
	"github.com/Alexander272/quiz/backend/internal/repository/postgres/pq_models"
	"github.com/jmoiron/sqlx"
)

type ResultRepo struct {
	db *sqlx.DB
}

func NewResultRepo(db *sqlx.DB) *ResultRepo {
	return &ResultRepo{
		db: db,
	}
}

type Result interface {
	Get(context.Context, *models.GetResults) ([]*models.Result, error)
}

func (r *ResultRepo) Get(ctx context.Context, req *models.GetResults) ([]*models.Result, error) {
	//TODO надо это все очень внимательно протестировать
	query := fmt.Sprintf(`SELECT a.question_id, array_agg(COALESCE(r.answer_id::text,'')) AS answer,
		array_agg(COALESCE(CASE WHEN is_correct THEN a.id::text END,'')) AS correct,
		bool_and(CASE WHEN a.id=r.answer_id THEN is_correct END) AS is_correct
		FROM %s AS a LEFT JOIN %s AS r ON a.id=r.answer_id WHERE quiz_id=$1 OR quiz_id IS NULL
		GROUP BY a.question_id
		HAVING bool_and(CASE WHEN a.id=r.answer_id THEN is_correct END) IS NOT NULL`,
		AnswerTable, ResultTable,
	)
	data := []*pq_models.Result{}

	if err := r.db.SelectContext(ctx, &data, query, req.QuizID); err != nil {
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}

	results := []*models.Result{}
	for _, d := range data {
		// if d.IsCorrect == nil {
		// 	continue
		// }

		answer := []string{}
		for _, a := range d.UserAnswer {
			if a != "" {
				answer = append(answer, a)
			}
		}
		correct := []string{}
		for _, a := range d.CorrectAnswer {
			if a != "" {
				correct = append(correct, a)
			}
		}

		results = append(results, &models.Result{
			QuestionID:    d.QuestionID,
			IsCorrect:     d.IsCorrect,
			UserAnswer:    answer,
			CorrectAnswer: correct,
		})
	}

	return results, nil
}

func (r *ResultRepo) Create(ctx context.Context) error {
	return fmt.Errorf("not implemented")
}

func (r *ResultRepo) Update(ctx context.Context) error {
	return fmt.Errorf("not implemented")
}
