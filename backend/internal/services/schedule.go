package services

import (
	"context"
	"fmt"

	"github.com/Alexander272/quiz/backend/internal/models"
	"github.com/Alexander272/quiz/backend/internal/repository"
)

type ScheduleService struct {
	repo repository.Schedule
}

func NewScheduleService(repo repository.Schedule) *ScheduleService {
	return &ScheduleService{
		repo: repo,
	}
}

type Schedule interface {
	Get(context.Context, *models.GetSchedule) ([]*models.Schedule, error)
	GetByQuiz(context.Context, *models.GetScheduleByQuiz) ([]*models.Schedule, error)
	Create(context.Context, *models.ScheduleDTO) (string, error)
	Update(context.Context, *models.ScheduleDTO) error
	Delete(context.Context, *models.DeleteScheduleDTO) error
}

func (s *ScheduleService) Get(ctx context.Context, req *models.GetSchedule) ([]*models.Schedule, error) {
	data, err := s.repo.Get(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get schedule. error: %w", err)
	}
	return data, nil
}

func (s *ScheduleService) GetByQuiz(ctx context.Context, req *models.GetScheduleByQuiz) ([]*models.Schedule, error) {
	data, err := s.repo.GetByQuiz(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get schedule by quiz. error: %w", err)
	}
	return data, nil
}

func (s *ScheduleService) Create(ctx context.Context, dto *models.ScheduleDTO) (string, error) {
	id, err := s.repo.Create(ctx, dto)
	if err != nil {
		return id, fmt.Errorf("failed to create schedule. error: %w", err)
	}
	return id, nil
}

func (s *ScheduleService) Update(ctx context.Context, dto *models.ScheduleDTO) error {
	if err := s.repo.Update(ctx, dto); err != nil {
		return fmt.Errorf("failed to update schedule. error: %w", err)
	}
	return nil
}

func (s *ScheduleService) Delete(ctx context.Context, dto *models.DeleteScheduleDTO) error {
	if err := s.repo.Delete(ctx, dto); err != nil {
		return fmt.Errorf("failed to delete schedule. error: %w", err)
	}
	return nil
}
