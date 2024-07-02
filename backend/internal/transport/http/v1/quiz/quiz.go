package quiz

import (
	"errors"
	"net/http"
	"time"

	"github.com/Alexander272/quiz/backend/internal/models"
	"github.com/Alexander272/quiz/backend/internal/models/response"
	"github.com/Alexander272/quiz/backend/internal/services"
	"github.com/Alexander272/quiz/backend/internal/transport/http/middleware"
	"github.com/Alexander272/quiz/backend/pkg/error_bot"
	"github.com/Alexander272/quiz/backend/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service services.Quiz
}

func NewHandler(service services.Quiz) *Handler {
	return &Handler{
		service: service,
	}
}

func Register(api *gin.RouterGroup, service services.Quiz, middleware *middleware.Middleware) {
	handler := NewHandler(service)

	quiz := api.Group("/quiz")
	{
		quiz.GET("", handler.get)
		quiz.GET("/:id", handler.getById)
		quiz.POST("", handler.create)
		quiz.PUT("/:id", handler.update)
		quiz.DELETE("/:id", handler.delete)
	}
}

func (h *Handler) get(c *gin.Context) {
	req := &models.GetQuizzesDTO{
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

func (h *Handler) getById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "id не задан")
		return
	}

	quiz, err := h.service.GetById(c, &models.GetQuizDTO{ID: id})
	if err != nil {
		if errors.Is(err, models.ErrNoRows) {
			response.NewErrorResponse(c, http.StatusNotFound, err.Error(), err.Error())
			return
		}
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), id)
		return
	}
	c.JSON(http.StatusOK, response.DataResponse{Data: quiz})
}

func (h *Handler) create(c *gin.Context) {
	dto := &models.QuizDTO{}
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
	logger.Info("Добавлен тест", logger.StringAttr("title", dto.Title), logger.StringAttr("author_id", dto.AuthorID))

	c.JSON(http.StatusCreated, response.IdResponse{Id: id, Message: "Тест добавлен"})
}

func (h *Handler) update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "Id не задан")
		return
	}

	dto := &models.QuizDTO{}
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
	logger.Info("Обновлен тест", logger.StringAttr("title", dto.Title), logger.StringAttr("author_id", dto.AuthorID))

	c.JSON(http.StatusOK, response.IdResponse{Message: "Тест обновлен"})
}

func (h *Handler) delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "Id не задан")
		return
	}

	if err := h.service.Delete(c, &models.DeleteQuizDTO{ID: id}); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), id)
		return
	}
	logger.Info("Удален тест", logger.StringAttr("id", id))

	c.JSON(http.StatusNoContent, response.IdResponse{})
}
