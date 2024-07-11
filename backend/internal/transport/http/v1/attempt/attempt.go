package attempt

import (
	"net/http"

	"github.com/Alexander272/quiz/backend/internal/models"
	"github.com/Alexander272/quiz/backend/internal/models/response"
	"github.com/Alexander272/quiz/backend/internal/services"
	"github.com/Alexander272/quiz/backend/internal/transport/http/middleware"
	"github.com/Alexander272/quiz/backend/pkg/error_bot"
	"github.com/Alexander272/quiz/backend/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service services.Attempt
}

func NewHandler(service services.Attempt) *Handler {
	return &Handler{
		service: service,
	}
}

func Register(api *gin.RouterGroup, service services.Attempt, middleware *middleware.Middleware) {
	handler := NewHandler(service)

	attempts := api.Group("/attempts", middleware.VerifyToken)
	{
		attempts.GET("", handler.get)
		attempts.GET("/:id", handler.getByID)
		attempts.POST("", handler.create)
		attempts.PUT("/:id", handler.update)
		attempts.DELETE("/:id", handler.delete)
	}
}

func (h *Handler) get(c *gin.Context) {
	userID := c.Query("user")
	scheduleID := c.Query("schedule")
	if userID == "" && scheduleID == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "id не задан")
		return
	}

	req := &models.GetAttempt{
		UserID:     userID,
		ScheduleID: scheduleID,
	}
	data, err := h.service.Get(c, req)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), req)
		return
	}
	c.JSON(http.StatusOK, response.DataResponse{Data: data, Total: len(data)})
}

func (h *Handler) getByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "id не задан")
		return
	}

	req := &models.GetAttemptByID{ID: id}
	data, err := h.service.GetByID(c, req)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), req)
		return
	}
	c.JSON(http.StatusOK, response.DataResponse{Data: data})
}

func (h *Handler) create(c *gin.Context) {
	dto := &models.AttemptDTO{}
	if err := c.BindJSON(dto); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "Отправлены некорректные данные")
		return
	}
	// dto.StartTime = time.Now().Unix()

	id, err := h.service.Create(c, dto)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), dto)
		return
	}
	logger.Info("Добавлена попытка", logger.StringAttr("user_id", dto.UserID), logger.StringAttr("schedule_id", dto.ScheduleID))

	c.JSON(http.StatusCreated, response.IdResponse{Id: id, Message: "Добавлена попытка"})
}

func (h *Handler) update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "id не задан")
		return
	}

	dto := &models.AttemptDTO{}
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
	logger.Info("Обновлена попытка", logger.StringAttr("user_id", dto.UserID), logger.StringAttr("schedule_id", dto.ScheduleID))

	c.JSON(http.StatusOK, response.IdResponse{Id: id, Message: "Обновлена попытка"})
}

func (h *Handler) delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "id не задан")
		return
	}

	dto := &models.DeleteAttemptDTO{ID: id}
	if err := h.service.Delete(c, dto); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), dto)
		return
	}
	logger.Info("Удалена попытка", logger.StringAttr("id", dto.ID))

	c.JSON(http.StatusNoContent, response.IdResponse{})
}
