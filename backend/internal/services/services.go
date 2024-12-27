package services

import (
	"time"

	"github.com/Alexander272/quiz/backend/internal/repository"
	"github.com/Alexander272/quiz/backend/pkg/auth"
)

type Services struct {
	Menu
	MenuItem
	Role
	Session
	Permission

	Media
	Quiz
	Question
	Answer
	Schedule
	Attempt
	AttemptDetails
}

type Deps struct {
	Repos           *repository.Repository
	Keycloak        *auth.KeycloakClient
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
	BotUrl          string
	ChannelId       string
	ErrorBotUrl     string
}

func NewServices(deps Deps) *Services {
	menuItem := NewMenuItemService(deps.Repos.MenuItem)
	menu := NewMenuService(deps.Repos.Menu, menuItem)
	role := NewRoleService(deps.Repos.Role)
	session := NewSessionService(deps.Keycloak, role)
	permission := NewPermissionService("configs/privacy.conf", menu, role)

	media := NewMediaService()
	answer := NewAnswerService(deps.Repos.Answer)
	question := NewQuestionService(&QuestionDeps{Repo: deps.Repos.Question, Answer: answer, Media: media})
	quiz := NewQuizService(deps.Repos.Quiz, media)
	schedule := NewScheduleService(deps.Repos.Schedule)
	attemptDetails := NewAttemptDetailsService(deps.Repos.AttemptDetails)
	attempt := NewAttemptService(&AttemptDeps{
		Repo:     deps.Repos.Attempt,
		Details:  attemptDetails,
		Keycloak: deps.Keycloak,
	})

	return &Services{
		MenuItem:   menuItem,
		Menu:       menu,
		Role:       role,
		Session:    session,
		Permission: permission,

		Media:          media,
		Quiz:           quiz,
		Question:       question,
		Answer:         answer,
		Schedule:       schedule,
		Attempt:        attempt,
		AttemptDetails: attemptDetails,
	}
}
