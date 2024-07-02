package question

import (
	"errors"
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
	service services.Question
}

func NewHandler(service services.Question) *Handler {
	return &Handler{
		service: service,
	}
}

func Register(api *gin.RouterGroup, service services.Question, middleware *middleware.Middleware) {
	handler := NewHandler(service)

	questions := api.Group("/questions")
	{
		questions.GET("", handler.get)
		questions.GET("/:id", handler.getById)
		questions.POST("", handler.create)
		questions.PUT("/:id", handler.update)
		questions.DELETE("/:id", handler.delete)
	}
}

func (h *Handler) get(c *gin.Context) {
	quizID := c.Query("quiz")
	if quizID == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "id не задан")
		return
	}
	hasShuffle := c.Query("shuffle")
	hasAnswers := c.Query("answers")

	req := &models.GetQuestionsDTO{
		QuizID:     quizID,
		HasShuffle: hasShuffle != "false",
		HasAnswers: hasAnswers != "true",
	}

	data, err := h.service.Get(c, req)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), req)
		return
	}
	c.JSON(http.StatusOK, response.DataResponse{Data: data, Total: len(data)})
}

func (h *Handler) getById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "id не задан")
		return
	}

	data, err := h.service.GetById(c, &models.GetQuestionDTO{ID: id})
	if err != nil {
		if errors.Is(err, models.ErrNoRows) {
			response.NewErrorResponse(c, http.StatusNotFound, err.Error(), err.Error())
			return
		}
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), id)
		return
	}
	c.JSON(http.StatusOK, response.DataResponse{Data: data})
}

func (h *Handler) create(c *gin.Context) {
	dto := &models.QuestionDTO{}
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
	logger.Info("Добавлен вопрос", logger.StringAttr("text", dto.Text), logger.StringAttr("quiz_id", dto.QuizID))

	c.JSON(http.StatusCreated, response.IdResponse{Id: id, Message: "Вопрос добавлен"})
}

func (h *Handler) update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "Id не задан")
		return
	}

	dto := &models.QuestionDTO{}
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
	logger.Info("Добавлен вопрос", logger.StringAttr("text", dto.Text), logger.StringAttr("quiz_id", dto.QuizID))

	c.JSON(http.StatusOK, response.IdResponse{Message: "Вопрос обновлен"})
}

func (h *Handler) delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "Id не задан")
		return
	}

	if err := h.service.Delete(c, &models.DeleteQuestionDTO{ID: id}); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), id)
		return
	}
	logger.Info("Удален вопрос", logger.StringAttr("id", id))

	c.JSON(http.StatusNoContent, response.IdResponse{})
}
