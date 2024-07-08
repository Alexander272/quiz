package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/Alexander272/quiz/backend/internal/models"
	"github.com/Alexander272/quiz/backend/internal/repository"
	"github.com/google/uuid"
)

type QuizService struct {
	repo  repository.Quiz
	media Media
}

func NewQuizService(repo repository.Quiz, media Media) *QuizService {
	return &QuizService{
		repo:  repo,
		media: media,
	}
}

type Quiz interface {
	Get(context.Context, *models.GetQuizzesDTO) ([]*models.Quiz, error)
	GetByAuthor(context.Context, string) ([]*models.Quiz, error)
	GetById(context.Context, *models.GetQuizDTO) (*models.Quiz, error)
	Create(context.Context, *models.QuizDTO) (string, error)
	Update(context.Context, *models.QuizDTO) error
	Delete(context.Context, *models.DeleteQuizDTO) error
}

func (s *QuizService) Check(ctx context.Context, dto *models.UserQuiz) (*models.UserResult, error) {

	return nil, fmt.Errorf("not implemented")
}

func (s *QuizService) Get(ctx context.Context, req *models.GetQuizzesDTO) ([]*models.Quiz, error) {
	data, err := s.repo.Get(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get quizzes. error: %w", err)
	}
	return data, nil
}

func (s *QuizService) GetByAuthor(ctx context.Context, authorId string) ([]*models.Quiz, error) {
	data, err := s.repo.GetByAuthor(ctx, authorId)
	if err != nil {
		return nil, fmt.Errorf("failed to get quizzes by author. error: %w", err)
	}
	return data, nil
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
	dto.ID = uuid.NewString()
	if dto.Image.Filename != "" {
		dto.ImageLink = fmt.Sprintf("media/%s/main/%s", dto.ID, dto.Image.Filename)
		if err := s.media.SaveFile(dto.Image, dto.ImageLink); err != nil {
			return "", err
		}
		dto.Image = nil
	}

	id, err := s.repo.Create(ctx, dto)
	if err != nil {
		return id, fmt.Errorf("failed to create quiz. error: %w", err)
	}
	return id, nil
}

func (s *QuizService) Update(ctx context.Context, dto *models.QuizDTO) error {
	if dto.ImageLink == "" {
		if err := s.media.Delete(fmt.Sprintf("media/%s/main", dto.ID)); err != nil {
			return err
		}
	}

	if dto.Image.Filename != "" {
		dto.ImageLink = fmt.Sprintf("media/%s/main/%s", dto.ID, dto.Image.Filename)
		if err := s.media.SaveFile(dto.Image, dto.ImageLink); err != nil {
			return err
		}
		dto.Image = nil
	}

	if err := s.repo.Update(ctx, dto); err != nil {
		return fmt.Errorf("failed to update quiz. error: %w", err)
	}
	return nil
}

func (s *QuizService) Delete(ctx context.Context, dto *models.DeleteQuizDTO) error {
	if err := s.media.Delete(fmt.Sprintf("media/%s", dto.ID)); err != nil {
		return err
	}

	if err := s.repo.Delete(ctx, dto); err != nil {
		return fmt.Errorf("failed to delete quiz. error: %w", err)
	}
	return nil
}
