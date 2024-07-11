package services

import (
	"context"
	"fmt"

	"github.com/Alexander272/quiz/backend/internal/models"
	"github.com/Alexander272/quiz/backend/internal/repository"
)

type AttemptService struct {
	repo repository.Attempt
}

func NewAttemptService(repo repository.Attempt) *AttemptService {
	return &AttemptService{
		repo: repo,
	}
}

type Attempt interface {
	Get(context.Context, *models.GetAttempt) ([]*models.Attempt, error)
	GetByID(context.Context, *models.GetAttemptByID) (*models.Attempt, error)
	Create(context.Context, *models.AttemptDTO) (string, error)
	Update(context.Context, *models.AttemptDTO) error
	Delete(context.Context, *models.DeleteAttemptDTO) error
}

func (s *AttemptService) Get(ctx context.Context, req *models.GetAttempt) ([]*models.Attempt, error) {
	data, err := s.repo.Get(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get attempt. error: %w", err)
	}
	return data, nil
}

func (s *AttemptService) GetByID(ctx context.Context, req *models.GetAttemptByID) (*models.Attempt, error) {
	data, err := s.repo.GetByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get attempt by id. error: %w", err)
	}
	return data, nil
}

func (s *AttemptService) Create(ctx context.Context, dto *models.AttemptDTO) (string, error) {
	id, err := s.repo.Create(ctx, dto)
	if err != nil {
		return id, fmt.Errorf("failed to create attempt. error: %w", err)
	}
	return id, nil
}

func (s *AttemptService) Update(ctx context.Context, dto *models.AttemptDTO) error {
	if err := s.repo.Update(ctx, dto); err != nil {
		return fmt.Errorf("failed to update attempt. error: %w", err)
	}
	return nil
}

func (s *AttemptService) Delete(ctx context.Context, dto *models.DeleteAttemptDTO) error {
	if err := s.repo.Delete(ctx, dto); err != nil {
		return fmt.Errorf("failed to delete attempt. error: %w", err)
	}
	return nil
}
