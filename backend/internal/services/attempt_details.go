package services

import (
	"context"
	"fmt"

	"github.com/Alexander272/quiz/backend/internal/models"
	"github.com/Alexander272/quiz/backend/internal/repository"
)

type AttemptDetailsService struct {
	repo repository.AttemptDetails
}

func NewAttemptDetailsService(repo repository.AttemptDetails) *AttemptDetailsService {
	return &AttemptDetailsService{
		repo: repo,
	}
}

type AttemptDetails interface {
	Get(ctx context.Context, req *models.GetAttemptDetails) (*models.AttemptDetails, error)
	CreateAll(ctx context.Context, dto *models.CreateAttemptDetailsDTO) error
	Create(ctx context.Context, dto *models.AttemptDetailDTO) error
	Update(ctx context.Context, dto *models.AttemptDetailDTO) error
}

func (s *AttemptDetailsService) Get(ctx context.Context, req *models.GetAttemptDetails) (*models.AttemptDetails, error) {
	data, err := s.repo.Get(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get attempt details. error: %w", err)
	}
	return data, nil
}

func (s *AttemptDetailsService) CreateAll(ctx context.Context, dto *models.CreateAttemptDetailsDTO) error {
	if err := s.repo.CreateAll(ctx, dto); err != nil {
		return fmt.Errorf("failed to create all attempt details. error: %w", err)
	}
	return nil
}

func (s *AttemptDetailsService) Create(ctx context.Context, dto *models.AttemptDetailDTO) error {
	if err := s.repo.Create(ctx, dto); err != nil {
		return fmt.Errorf("failed to create attempt detail. error: %w", err)
	}
	return nil
}

func (s *AttemptDetailsService) Update(ctx context.Context, dto *models.AttemptDetailDTO) error {
	if err := s.repo.Update(ctx, dto); err != nil {
		return fmt.Errorf("failed to update attempt detail. error: %w", err)
	}
	return nil
}
