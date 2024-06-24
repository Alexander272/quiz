package v1

import (
	"github.com/Alexander272/quiz/backend/internal/config"
	"github.com/Alexander272/quiz/backend/internal/services"
	"github.com/Alexander272/quiz/backend/internal/transport/http/middleware"
	"github.com/Alexander272/quiz/backend/internal/transport/http/v1/auth"
	"github.com/Alexander272/quiz/backend/internal/transport/http/v1/quiz"
	"github.com/gin-gonic/gin"
)

// type Handler struct {
// 	services   *services.Services
// 	conf       *config.Config
// 	middleware *middleware.Middleware
// }

// func NewHandler(deps *Deps) *Handler {
// 	return &Handler{
// 		services:   deps.Services,
// 		conf:       deps.Conf,
// 		middleware: deps.Middleware,
// 	}
// }

// func (h *Handler) Init(group *gin.RouterGroup) {
// 	// v1 := group.Group("/v1")

// 	// auth.Register(v1, auth.Deps{Service: h.services.Session, Auth: h.conf.Auth})
// }

type Deps struct {
	Services   *services.Services
	Conf       *config.Config
	Middleware *middleware.Middleware
}

func Register(api *gin.RouterGroup, deps *Deps) {
	v1 := api.Group("/v1")
	auth.Register(v1, &auth.Deps{Service: deps.Services.Session, Auth: deps.Conf.Auth})
	quiz.Register(v1, deps.Services.Quiz, deps.Middleware)
}
