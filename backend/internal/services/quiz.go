package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/Alexander272/quiz/backend/internal/models"
	"github.com/Alexander272/quiz/backend/internal/repository"
)

type QuizService struct {
	repo repository.Quiz
}

func NewQuizService(repo repository.Quiz) *QuizService {
	return &QuizService{
		repo: repo,
	}
}

type Quiz interface {
	GetById(context.Context, *models.GetQuizDTO) (*models.Quiz, error)
	Create(context.Context, *models.QuizDTO) (string, error)
	Update(context.Context, *models.QuizDTO) error
	Delete(context.Context, *models.DeleteQuizDTO) error
}

func (s *QuizService) GetById(ctx context.Context, req *models.GetQuizDTO) (*models.Quiz, error) {
	data, err := s.repo.GetById(ctx, req)
	if err != nil {
		if errors.Is(err, models.ErrNoRows) {
			return nil, err
		}
		return nil, fmt.Errorf("failed to get quiz by id. error: %w", err)
	}
	return data, nil
}

func (s *QuizService) Create(ctx context.Context, dto *models.QuizDTO) (string, error) {
	id, err := s.repo.Create(ctx, dto)
	if err != nil {
		return id, fmt.Errorf("failed to create quiz. error: %w", err)
	}
	return id, nil
}

func (s *QuizService) Update(ctx context.Context, dto *models.QuizDTO) error {
	if err := s.repo.Update(ctx, dto); err != nil {
		return fmt.Errorf("failed to update quiz. error: %w", err)
	}
	return nil
}

func (s *QuizService) Delete(ctx context.Context, dto *models.DeleteQuizDTO) error {
	if err := s.repo.Delete(ctx, dto); err != nil {
		return fmt.Errorf("failed to delete quiz. error: %w", err)
	}
	return nil
}
