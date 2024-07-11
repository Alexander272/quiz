package schedule

import (
	"net/http"
	"time"

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
	service services.Schedule
}

func NewHandler(service services.Schedule) *Handler {
	return &Handler{
		service: service,
	}
}

func Register(api *gin.RouterGroup, service services.Schedule, middleware *middleware.Middleware) {
	handler := NewHandler(service)

	schedule := api.Group("schedule")
	{
		schedule.GET("", handler.get)
		schedule.GET("/:quiz", handler.getByQuiz)
		schedule.POST("", handler.create)
		schedule.PUT("/:id", handler.update)
		schedule.DELETE("/:id", handler.delete)
	}
}

func (h *Handler) get(c *gin.Context) {
	req := &models.GetSchedule{
		Time: time.Now().Unix(),
	}

	data, err := h.service.Get(c, req)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), req)
		return
	}
	c.JSON(http.StatusOK, response.DataResponse{Data: data, Total: len(data)})
}

func (h *Handler) getByQuiz(c *gin.Context) {
	quiz := c.Param("quiz")
	if quiz == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "id не задан")
		return
	}

	req := &models.GetScheduleByQuiz{QuizID: quiz}
	data, err := h.service.GetByQuiz(c, req)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), req)
		return
	}
	c.JSON(http.StatusOK, response.DataResponse{Data: data, Total: len(data)})
}

func (h *Handler) create(c *gin.Context) {
	dto := &models.ScheduleDTO{}
	if err := c.BindJSON(dto); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "Отправлены некорректные данные")
		return
	}

	id, err := h.service.Create(c, dto)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), dto)
		return
	}
	logger.Info("Расписание добавлено",
		logger.StringAttr("quiz_id", dto.QuizID),
		logger.StringAttr("start_time", time.Unix(dto.StartTime, 0).Format(constants.DateFormat)),
		logger.StringAttr("end_time", time.Unix(dto.EndTime, 0).Format(constants.DateFormat)),
	)

	c.JSON(http.StatusCreated, response.IdResponse{Id: id, Message: "Расписание добавлено"})
}

func (h *Handler) update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "id не задан")
		return
	}

	dto := &models.ScheduleDTO{}
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
	logger.Info("Расписание обновлено",
		logger.StringAttr("quiz_id", dto.QuizID),
		logger.StringAttr("start_time", time.Unix(dto.StartTime, 0).Format(constants.DateFormat)),
		logger.StringAttr("end_time", time.Unix(dto.EndTime, 0).Format(constants.DateFormat)),
	)

	c.JSON(http.StatusOK, response.IdResponse{Id: id, Message: "Расписание обновлено"})
}

func (h *Handler) delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "id не задан")
		return
	}

	dto := &models.DeleteScheduleDTO{ID: id}
	if err := h.service.Delete(c, dto); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), dto)
		return
	}
	logger.Info("Расписание удалено", logger.StringAttr("id", id))

	c.JSON(http.StatusNoContent, response.IdResponse{})
}
