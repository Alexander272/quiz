package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/Alexander272/quiz/backend/internal/models"
	"github.com/Alexander272/quiz/backend/internal/repository"
)

type QuestionService struct {
	repo repository.Question
}

func NewQuestionService(repo repository.Question) *QuestionService {
	return &QuestionService{
		repo: repo,
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
	//TODO надо наверное тут получать ответы и перемешивать вопросы
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
	//TODO надо наверное тут получать ответы
	return data, nil
}

func (s *QuestionService) Create(ctx context.Context, dto *models.QuestionDTO) (string, error) {
	id, err := s.repo.Create(ctx, dto)
	if err != nil {
		return id, fmt.Errorf("failed to create question. error: %w", err)
	}
	return id, nil
}

func (s *QuestionService) Update(ctx context.Context, dto *models.QuestionDTO) error {
	if err := s.repo.Update(ctx, dto); err != nil {
		return fmt.Errorf("failed to update question. error: %w", err)
	}
	return nil
}

func (s *QuestionService) Delete(ctx context.Context, dto *models.DeleteQuestionDTO) error {
	if err := s.repo.Delete(ctx, dto); err != nil {
		return fmt.Errorf("failed to delete question. error: %w", err)
	}
	return nil
}
