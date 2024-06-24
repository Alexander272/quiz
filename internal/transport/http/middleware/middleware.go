package middleware

import (
	"github.com/Alexander272/quiz/backend/internal/config"
	"github.com/Alexander272/quiz/backend/internal/services"
	"github.com/Alexander272/quiz/backend/pkg/auth"
)

type Middleware struct {
	keycloak *auth.KeycloakClient
	services *services.Services
	auth     config.AuthConfig
}

func NewMiddleware(services *services.Services, auth config.AuthConfig, keycloak *auth.KeycloakClient) *Middleware {
	return &Middleware{
		keycloak: keycloak,
		services: services,
		auth:     auth,
	}
}
