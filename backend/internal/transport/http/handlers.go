package http

import (
	"net/http"

	"github.com/Alexander272/quiz/backend/internal/config"
	"github.com/Alexander272/quiz/backend/internal/services"
	"github.com/Alexander272/quiz/backend/internal/transport/http/middleware"
	httpV1 "github.com/Alexander272/quiz/backend/internal/transport/http/v1"
	"github.com/Alexander272/quiz/backend/pkg/auth"
	"github.com/Alexander272/quiz/backend/pkg/limiter"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	keycloak *auth.KeycloakClient
	services *services.Services
}

func NewHandler(services *services.Services, keycloak *auth.KeycloakClient) *Handler {
	return &Handler{
		services: services,
		keycloak: keycloak,
	}
}

func (h *Handler) Init(conf *config.Config) *gin.Engine {
	router := gin.Default()

	router.Use(
		limiter.Limit(conf.Limiter.RPS, conf.Limiter.Burst, conf.Limiter.TTL),
	)

	// Init router
	router.GET("/api/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router, conf)

	return router
}

func (h *Handler) initAPI(router *gin.Engine, conf *config.Config) {
	api := router.Group("/api")
	middleware := middleware.NewMiddleware(h.services, conf.Auth, h.keycloak)
	httpV1.Register(api, &httpV1.Deps{Services: h.services, Conf: conf, Middleware: middleware})
}
