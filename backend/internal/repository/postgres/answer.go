package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/Alexander272/quiz/backend/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AnswerRepo struct {
	db *sqlx.DB
}

func NewAnswerRepo(db *sqlx.DB) *AnswerRepo {
	return &AnswerRepo{
		db: db,
	}
}

type Answer interface {
	GetByQuiz(context.Context, *models.GetAnswersDTO) ([]*models.AnswerList, error)
	GetByQuestion(context.Context, *models.GetAnswersDTO) (*models.AnswerList, error)
	Create(context.Context, *models.AnswerDTO) (string, error)
	CreateSeveral(context.Context, []*models.AnswerDTO) error
	Update(context.Context, *models.AnswerDTO) error
	Delete(context.Context, *models.DeleteAnswerDTO) error
}

func (r *AnswerRepo) GetByQuiz(ctx context.Context, req *models.GetAnswersDTO) ([]*models.AnswerList, error) {
	query := fmt.Sprintf(`SELECT a.id, question_id, a.number, a.text, a.image, is_correct FROM %s AS a
		INNER JOIN %s AS q ON q.id=a.question_id WHERE quiz_id=$1 ORDER BY question_id, a.number`,
		AnswerTable, QuestionTable,
	)
	data := []*models.Answer{}

	if err := r.db.SelectContext(ctx, &data, query, req.QuizID); err != nil {
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}

	list := []*models.AnswerList{}
	for i, d := range data {
		if !req.HasCorrect {
			d.IsCorrect = false
		}

		if i == 0 || list[len(list)-1].QuestionID != d.QuestionID {
			list = append(list, &models.AnswerList{
				QuestionID: d.QuestionID,
				List:       []*models.Answer{d},
			})
		} else {
			list[len(list)-1].List = append(list[len(list)-1].List, d)
		}
	}

	return nil, fmt.Errorf("not implemented")
}

func (r *AnswerRepo) GetByQuestion(ctx context.Context, req *models.GetAnswersDTO) (*models.AnswerList, error) {
	query := fmt.Sprintf(`SELECT id, question_id, number, text, image, is_correct FROM %s WHERE question_id=$1 ORDER BY number`, AnswerTable)
	data := []*models.Answer{}

	if err := r.db.SelectContext(ctx, &data, query, req.QuestionID); err != nil {
		return nil, fmt.Errorf("failed to execute query. error: %w", err)
	}

	if len(data) == 0 {
		return nil, nil
	}

	answer := &models.AnswerList{QuestionID: data[0].QuestionID}
	for _, d := range data {
		if !req.HasCorrect {
			d.IsCorrect = false
		}
		answer.List = append(answer.List, d)
	}
	return answer, nil
}

func (r *AnswerRepo) Create(ctx context.Context, dto *models.AnswerDTO) (string, error) {
	query := fmt.Sprintf(`INSERT INTO %s (id, question_id, number, text, image, is_correct) VALUES (:id, :question_id, :number, :text,
		:image, :is_correct)`,
		AnswerTable,
	)
	dto.ID = uuid.New().String()

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return "", fmt.Errorf("failed to execute query. error: %w", err)
	}
	return dto.ID, nil
}

func (r *AnswerRepo) CreateSeveral(ctx context.Context, dto []*models.AnswerDTO) error {
	values := []string{}
	args := []interface{}{}
	c := 6
	for i, v := range dto {
		id := uuid.New()
		args = append(args, id, v.QuestionID, v.Number, v.Text, v.Image, v.IsCorrect)
		values = append(values, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d)", c*i+1, c*i+2, c*i+3, c*i+4, c*i+5, c*i+6))
	}
	query := fmt.Sprintf(`INSERT INTO %s (id, question_id, number, text, image, is_correct) VALUES %s`, AnswerTable, strings.Join(values, ","))

	if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}

func (r *AnswerRepo) Update(ctx context.Context, dto *models.AnswerDTO) error {
	query := fmt.Sprintf(`UPDATE %s SET number=:number, text=:text, image=:image, is_correct=:is_correct WHERE id=:id`, AnswerTable)

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}

// TODO возможно стоит сделать UpdateSeveral

func (r *AnswerRepo) Delete(ctx context.Context, dto *models.DeleteAnswerDTO) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=:id`, AnswerTable)

	if _, err := r.db.NamedExecContext(ctx, query, dto); err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	return nil
}
