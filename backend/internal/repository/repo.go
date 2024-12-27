package repository

import (
	"github.com/Alexander272/quiz/backend/internal/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Role interface {
	postgres.Role
}
type MenuItem interface {
	postgres.MenuItem
}
type Menu interface {
	postgres.Menu
}

type Quiz interface {
	postgres.Quiz
}
type Question interface {
	postgres.Question
}
type Answer interface {
	postgres.Answer
}
type Schedule interface {
	postgres.Schedule
}
type Attempt interface {
	postgres.Attempt
}
type AttemptDetails interface {
	postgres.AttemptDetails
}

type Repository struct {
	Role
	MenuItem
	Menu

	Quiz
	Question
	Answer
	Schedule
	Attempt
	AttemptDetails
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Role:     postgres.NewRoleRepo(db),
		MenuItem: postgres.NewMenuItemRepo(db),
		Menu:     postgres.NewMenuRepo(db),

		Quiz:           postgres.NewQuizRepo(db),
		Question:       postgres.NewQuestionRepo(db),
		Answer:         postgres.NewAnswerRepo(db),
		Schedule:       postgres.NewScheduleRepo(db),
		Attempt:        postgres.NewAttemptRepo(db),
		AttemptDetails: postgres.NewAttemptDetailsRepo(db),
	}
}
