package services

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"github.com/Alexander272/quiz/backend/internal/models"
	"github.com/Alexander272/quiz/backend/internal/repository"
	"github.com/google/uuid"
)

type QuestionService struct {
	repo   repository.Question
	answer Answer
	media  Media
}

type QuestionDeps struct {
	Repo   repository.Question
	Answer Answer
	Media  Media
}

func NewQuestionService(deps *QuestionDeps) *QuestionService {
	return &QuestionService{
		repo:   deps.Repo,
		answer: deps.Answer,
		media:  deps.Media,
	}
}

type Question interface {
	Get(context.Context, *models.GetQuestionsDTO) ([]*models.Question, error)
	GetById(context.Context, *models.GetQuestionDTO) (*models.Question, error)
	Create(context.Context, *models.QuestionDTO) (string, error)
	Update(context.Context, *models.QuestionDTO) error
	Delete(context.Context, *models.DeleteQuestionDTO) error
}

func (s *QuestionService) Get(ctx context.Context, req *models.GetQuestionsDTO) ([]*models.Question, error) {
	data, err := s.repo.Get(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get questions. error: %w", err)
	}

	if req.HasShuffle {
		rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })
	}

	if !req.HasAnswers {
		return data, nil
	}

	//TODO нужны ли мне здесь правильные ответы?
	answers, err := s.answer.GetByQuiz(ctx, &models.GetAnswersDTO{QuizID: req.QuizID, HasCorrect: !req.HasShuffle && req.HasAnswers})
	if err != nil {
		return nil, err
	}

	for _, d := range data {
		answer := &models.AnswerList{}
		for _, item := range answers {
			if item.QuestionID == d.ID {
				answer = item
				break
			}
		}

		if req.HasShuffle && d.HasShuffle {
			rand.Shuffle(len(answer.List), func(i, j int) { answer.List[i], answer.List[j] = answer.List[j], answer.List[i] })
		}
		d.Answers = answer.List
	}

	return data, nil
}

func (s *QuestionService) GetById(ctx context.Context, req *models.GetQuestionDTO) (*models.Question, error) {
	data, err := s.repo.GetById(ctx, req)
	if err != nil {
		if errors.Is(err, models.ErrNoRows) {
			return nil, err
		}
		return nil, fmt.Errorf("failed to get question by id. error: %w", err)
	}

	//TODO нужны ли мне здесь правильные ответы и перемешивание ответов?
	answers, err := s.answer.GetByQuestion(ctx, &models.GetAnswersDTO{QuestionID: req.ID})
	if err != nil {
		return nil, err
	}
	data.Answers = answers.List

	return data, nil
}

func (s *QuestionService) Create(ctx context.Context, dto *models.QuestionDTO) (string, error) {
	dto.ID = uuid.NewString()
	if dto.Image != "" {
		dst := strings.Replace(dto.Image, "temp", dto.ID, 1)
		if err := s.media.Move(dto.Image, dst); err != nil {
			return dto.ID, err
		}
		dto.Image = dst
	}

	id, err := s.repo.Create(ctx, dto)
	if err != nil {
		return id, fmt.Errorf("failed to create question. error: %w", err)
	}

	for i := range dto.Answers {
		dto.Answers[i].QuestionID = id
	}
	if err := s.answer.CreateSeveral(ctx, dto.Answers); err != nil {
		return id, err
	}

	return id, nil
}

func (s *QuestionService) Update(ctx context.Context, dto *models.QuestionDTO) error {
	if err := s.repo.Update(ctx, dto); err != nil {
		return fmt.Errorf("failed to update question. error: %w", err)
	}

	if err := s.answer.DeleteByQuestionId(ctx, dto.ID); err != nil {
		return err
	}
	for i := range dto.Answers {
		dto.Answers[i].QuestionID = dto.ID
	}
	if err := s.answer.CreateSeveral(ctx, dto.Answers); err != nil {
		return err
	}

	return nil
}

func (s *QuestionService) Delete(ctx context.Context, dto *models.DeleteQuestionDTO) error {
	if err := s.repo.Delete(ctx, dto); err != nil {
		return fmt.Errorf("failed to delete question. error: %w", err)
	}

	if err := s.media.Delete(fmt.Sprintf("media/%s/%s", dto.QuizID, dto.ID)); err != nil {
		return err
	}

	return nil
}
