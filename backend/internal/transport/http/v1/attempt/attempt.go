package attempt

import (
	"net/http"
	"strings"
	"time"

	"github.com/Alexander272/quiz/backend/internal/constants"
	"github.com/Alexander272/quiz/backend/internal/models"
	"github.com/Alexander272/quiz/backend/internal/models/response"
	"github.com/Alexander272/quiz/backend/internal/services"
	"github.com/Alexander272/quiz/backend/internal/transport/http/middleware"
	"github.com/Alexander272/quiz/backend/internal/transport/http/v1/attempt/details"
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

func Register(api *gin.RouterGroup, services *services.Services, middleware *middleware.Middleware) {
	handler := NewHandler(services.Attempt)

	attempts := api.Group("/attempts", middleware.VerifyToken)
	{
		attempts.GET("", handler.get)
		attempts.GET("/:id", handler.getByID)
		// attempts.GET("/quiz/:quiz")
		attempts.POST("/save", handler.saveDetails)
		attempts.POST("/finish", handler.finish)
		attempts.POST("", handler.create)
		attempts.PUT("/:id", handler.update)
		attempts.DELETE("/:id", handler.delete)
	}

	details.Register(attempts, services.AttemptDetails, middleware)
}

func (h *Handler) get(c *gin.Context) {
	quiz := c.Query("quiz")
	scheduleID := c.Query("schedule")
	if quiz == "" && scheduleID == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "id не задан")
		return
	}

	if quiz != "" {
		h.getByQuiz(c)
		return
	}

	req := &models.GetAttempt{
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

func (h *Handler) getByQuiz(c *gin.Context) {
	quizID := c.Query("quiz")
	active := c.Query("active")

	u, exists := c.Get(constants.CtxUser)
	if !exists {
		response.NewErrorResponse(c, http.StatusUnauthorized, "empty user", "сессия не найдена")
		return
	}
	user := u.(models.User)

	req := &models.GetAttemptByQuiz{
		QuizID: quizID,
		UserID: user.ID,
		Time:   time.Now().Unix(),
	}
	data, err := h.service.GetByQuiz(c, req)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), req)
		return
	}

	if active == "true" && len(data) > 0 {
		if data[0].EndTime == 0 {
			data = []*models.Attempt{data[0]}
		} else {
			data = []*models.Attempt{}
		}
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

func (h *Handler) saveDetails(c *gin.Context) {
	dto := []*models.AttemptDetailDTO{}
	if err := c.BindJSON(&dto); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "Отправлены некорректные данные")
		return
	}

	if err := h.service.SaveDetails(c, dto); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), dto)
		return
	}
	c.JSON(http.StatusOK, response.IdResponse{})
}

func (h *Handler) finish(c *gin.Context) {
	dto := &models.FinishAttempt{}
	if err := c.BindJSON(&dto); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "Отправлены некорректные данные")
		return
	}

	data, err := h.service.Finish(c, dto)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), dto)
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
	//TODO на клиенте происходит авто сохранения => надо как-то обновлять (и понять что именно надо обновлять) сохраненные данные
	// все же просто, записи в который есть id нужно обновить, остальные нужно создать

	dto.Token = strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", 1)

	u, exists := c.Get(constants.CtxUser)
	if !exists {
		response.NewErrorResponse(c, http.StatusUnauthorized, "empty user", "сессия не найдена")
		return
	}
	user := u.(models.User)
	dto.UserID = user.ID

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

	u, exists := c.Get(constants.CtxUser)
	if !exists {
		response.NewErrorResponse(c, http.StatusUnauthorized, "empty user", "сессия не найдена")
		return
	}
	user := u.(models.User)
	dto.UserID = user.ID

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
