package services

import (
	"context"
	"fmt"
	"time"

	"github.com/Alexander272/quiz/backend/internal/models"
	"github.com/Alexander272/quiz/backend/internal/repository"
	"github.com/Alexander272/quiz/backend/pkg/auth"
)

type AttemptService struct {
	repo     repository.Attempt
	details  AttemptDetails
	keycloak *auth.KeycloakClient
}

type AttemptDeps struct {
	Repo     repository.Attempt
	Details  AttemptDetails
	Keycloak *auth.KeycloakClient
}

func NewAttemptService(deps *AttemptDeps) *AttemptService {
	return &AttemptService{
		repo:     deps.Repo,
		details:  deps.Details,
		keycloak: deps.Keycloak,
	}
}

type Attempt interface {
	SaveDetails(context.Context, []*models.AttemptDetailDTO) error
	Finish(context.Context, *models.FinishAttempt) (*models.Attempt, error)
	Get(context.Context, *models.GetAttempt) ([]*models.Attempt, error)
	GetByQuiz(context.Context, *models.GetAttemptByQuiz) ([]*models.Attempt, error)
	GetByID(context.Context, *models.GetAttemptByID) (*models.Attempt, error)
	Create(context.Context, *models.AttemptDTO) (string, error)
	Update(context.Context, *models.AttemptDTO) error
	Delete(context.Context, *models.DeleteAttemptDTO) error
}

func (s *AttemptService) SaveDetails(ctx context.Context, dto []*models.AttemptDetailDTO) error {
	updated := []*models.AttemptDetailDTO{}
	new := []*models.AttemptDetailDTO{}

	//TODO если я не передам id с клиента он создаст новую запись, что и логично, но это проблема
	for _, d := range dto {
		if d.ID == "" {
			new = append(new, d)
		} else {
			updated = append(updated, d)
		}
	}

	if err := s.details.UpdateSeveral(ctx, updated); err != nil {
		return err
	}
	if err := s.details.CreateSeveral(ctx, new); err != nil {
		return err
	}
	return nil
}

func (s *AttemptService) Finish(ctx context.Context, dto *models.FinishAttempt) (*models.Attempt, error) {
	// примерная работа функции
	// GET attempt detail
	// GET questions with right answers (IF attempt details without correct answers)
	// Compare user answers with quiz answers
	// calculate correct answers and points
	// UPDATE attempt
	// RETURN attempt

	// if len(dto.Questions)>0 {
	// 	s.details.CreateAll(ctx, &models.CreateAttemptDetailsDTO{
	// 		AttemptID: dto.ID,
	// 		Questions: dto.Questions,
	// 	})

	// }

	data, err := s.details.Get(ctx, &models.GetAttemptDetails{AttemptID: dto.ID, ShowAnswers: true})
	if err != nil {
		return nil, err
	}

	correct := 0
	points := 0
	totalPoints := 0

	for _, q := range data.Questions {
		if q.IsCorrect {
			correct++
			points += q.Points
		}
		totalPoints += q.Points
	}
	dto.EndTime = time.Now().Unix()

	update := &models.AttemptDTO{ID: dto.ID, EndTime: dto.EndTime, Correct: correct, Points: points, TotalPoints: totalPoints}
	if err := s.Update(ctx, update); err != nil {
		return nil, err
	}

	result := &models.Attempt{
		ID:          dto.ID,
		StartTime:   dto.StartTime,
		EndTime:     dto.EndTime,
		Total:       dto.Total,
		Points:      points,
		Correct:     correct,
		TotalPoints: totalPoints,
	}
	return result, nil
}

func (s *AttemptService) Get(ctx context.Context, req *models.GetAttempt) ([]*models.Attempt, error) {
	data, err := s.repo.Get(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get attempt. error: %w", err)
	}
	return data, nil
}

func (s *AttemptService) GetByQuiz(ctx context.Context, req *models.GetAttemptByQuiz) ([]*models.Attempt, error) {
	data, err := s.repo.GetByQuiz(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get attempt by quiz. error: %w", err)
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
	user, err := s.keycloak.Client.GetUserInfo(ctx, dto.Token, s.keycloak.Realm)
	if err != nil {
		return "", fmt.Errorf("failed to get user info. error: %w", err)
	}

	dto.StartTime = time.Now().Unix()
	if user.Name != nil {
		dto.Username = *user.Name
	}

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
