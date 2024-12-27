package details

import (
	"net/http"

	"github.com/Alexander272/quiz/backend/internal/constants"
	"github.com/Alexander272/quiz/backend/internal/models"
	"github.com/Alexander272/quiz/backend/internal/models/response"
	"github.com/Alexander272/quiz/backend/internal/services"
	"github.com/Alexander272/quiz/backend/internal/transport/http/middleware"
	"github.com/Alexander272/quiz/backend/pkg/error_bot"
	"github.com/Alexander272/quiz/backend/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service services.AttemptDetails
}

func NewHandler(service services.AttemptDetails) *Handler {
	return &Handler{
		service: service,
	}
}

func Register(api *gin.RouterGroup, service services.AttemptDetails, middleware *middleware.Middleware) {
	handler := NewHandler(service)

	api.GET("/:id/details", handler.get)
	details := api.Group("/details")
	{
		details.POST("", handler.create)
		details.PUT("/:id", handler.update)
	}
}

func (h *Handler) get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "id не задан")
		return
	}
	req := &models.GetAttemptDetails{AttemptID: id}

	show := c.Query("show")
	if show == "true" {
		req.ShowAnswers = true
	}

	data, err := h.service.Get(c, req)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), req)
		return
	}

	c.JSON(http.StatusOK, response.DataResponse{Data: data})
}

func (h *Handler) create(c *gin.Context) {
	dto := &models.AttemptDetailDTO{}
	if err := c.BindJSON(dto); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "Отправлены некорректные данные")
		return
	}

	if err := h.service.Create(c, dto); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), dto)
		return
	}

	u, exists := c.Get(constants.CtxUser)
	if !exists {
		response.NewErrorResponse(c, http.StatusUnauthorized, "empty user", "сессия не найдена")
		return
	}
	user := u.(models.User)
	logger.Info("Сохранен ответ", logger.StringAttr("user_id", user.ID), logger.AnyAttr("dto", dto))

	c.JSON(http.StatusNoContent, response.IdResponse{})
}

func (h *Handler) update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "id не задан")
		return
	}

	dto := &models.AttemptDetailDTO{}
	if err := c.BindJSON(dto); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "Отправлены некорректные данные")
		return
	}
	dto.ID = id

	if err := h.service.Update(c, dto); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), dto)
		return
	}

	u, exists := c.Get(constants.CtxUser)
	if !exists {
		response.NewErrorResponse(c, http.StatusUnauthorized, "empty user", "сессия не найдена")
		return
	}
	user := u.(models.User)
	logger.Info("Обновлен ответ", logger.StringAttr("user_id", user.ID), logger.AnyAttr("dto", dto))

	c.JSON(http.StatusNoContent, response.IdResponse{})
}
