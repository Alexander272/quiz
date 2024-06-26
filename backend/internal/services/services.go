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

	Quiz
	Question
	Answer
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

	answer := NewAnswerService(deps.Repos.Answer)
	question := NewQuestionService(deps.Repos.Question, answer)
	quiz := NewQuizService(deps.Repos.Quiz)

	return &Services{
		MenuItem:   menuItem,
		Menu:       menu,
		Role:       role,
		Session:    session,
		Permission: permission,

		Quiz:     quiz,
		Question: question,
		Answer:   answer,
	}
}
