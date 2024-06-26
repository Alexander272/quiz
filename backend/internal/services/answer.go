package services

import (
	"context"
	"fmt"

	"github.com/Alexander272/quiz/backend/internal/models"
	"github.com/Alexander272/quiz/backend/internal/repository"
)

type AnswerService struct {
	repo repository.Answer
}

func NewAnswerService(repo repository.Answer) *AnswerService {
	return &AnswerService{
		repo: repo,
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

func (s *AnswerService) GetByQuiz(ctx context.Context, req *models.GetAnswersDTO) ([]*models.AnswerList, error) {
	data, err := s.repo.GetByQuiz(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get answers by quiz. error: %w", err)
	}
	return data, nil
}

func (s *AnswerService) GetByQuestion(ctx context.Context, req *models.GetAnswersDTO) (*models.AnswerList, error) {
	data, err := s.repo.GetByQuestion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get answers by question. error: %w", err)
	}
	return data, nil
}

func (s *AnswerService) Create(ctx context.Context, dto *models.AnswerDTO) (string, error) {
	id, err := s.repo.Create(ctx, dto)
	if err != nil {
		return id, fmt.Errorf("failed to create answer. error: %w", err)
	}
	return id, nil
}

func (s *AnswerService) CreateSeveral(ctx context.Context, dto []*models.AnswerDTO) error {
	if err := s.repo.CreateSeveral(ctx, dto); err != nil {
		return fmt.Errorf("failed to create several answers. error: %w", err)
	}
	return nil
}

func (s *AnswerService) Update(ctx context.Context, dto *models.AnswerDTO) error {
	if err := s.repo.Update(ctx, dto); err != nil {
		return fmt.Errorf("failed to update answer. error: %w", err)
	}
	return nil
}

func (s *AnswerService) Delete(ctx context.Context, dto *models.DeleteAnswerDTO) error {
	if err := s.repo.Delete(ctx, dto); err != nil {
		return fmt.Errorf("failed to delete answer. error: %w", err)
	}
	return nil
}
