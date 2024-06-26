package answer

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
	service services.Answer
}

func NewHandler(service services.Answer) *Handler {
	return &Handler{
		service: service,
	}
}

func Register(api *gin.RouterGroup, service services.Answer, middleware *middleware.Middleware) {
	handler := NewHandler(service)

	answers := api.Group("/answers")
	{
		answers.GET("/quiz/:id", handler.getByQuiz)
		answers.GET("/:questionId", handler.getByQuestion)
		answers.POST("", handler.create)
		answers.POST("/several", handler.createSeveral)
		answers.PUT("/:id", handler.update)
		answers.DELETE("/:id", handler.delete)
	}
}

func (h *Handler) getByQuiz(c *gin.Context) {
	quizId := c.Param("id")
	if quizId == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "id не задан")
		return
	}
	req := &models.GetAnswersDTO{QuizID: quizId, HasCorrect: true}

	data, err := h.service.GetByQuiz(c, req)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), req)
		return
	}
	c.JSON(http.StatusOK, response.DataResponse{Data: data, Total: len(data)})
}

func (h *Handler) getByQuestion(c *gin.Context) {
	questionId := c.Param("questionId")
	if questionId == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "id не задан")
		return
	}
	req := &models.GetAnswersDTO{QuestionID: questionId, HasCorrect: true}

	data, err := h.service.GetByQuestion(c, req)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), req)
		return
	}
	c.JSON(http.StatusOK, response.DataResponse{Data: data.List, Total: len(data.List)})
}

func (h *Handler) create(c *gin.Context) {
	dto := &models.AnswerDTO{}
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
	logger.Info("Добавлен ответ", logger.StringAttr("text", dto.Text), logger.StringAttr("question_id", dto.QuestionID))

	c.JSON(http.StatusCreated, response.IdResponse{Id: id, Message: "Ответ добавлен"})
}

func (h *Handler) createSeveral(c *gin.Context) {
	dto := []*models.AnswerDTO{}
	if err := c.BindJSON(&dto); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "Отправлены некорректные данные")
		return
	}

	if err := h.service.CreateSeveral(c, dto); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), dto)
		return
	}
	logger.Info("Добавлен ответы", logger.StringAttr("question_id", dto[0].QuestionID))

	c.JSON(http.StatusCreated, response.IdResponse{Message: "Ответы добавлены"})
}

func (h *Handler) update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "Id не задан")
		return
	}

	dto := &models.AnswerDTO{}
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
	logger.Info("Добавлен ответ", logger.StringAttr("text", dto.Text), logger.StringAttr("question_id", dto.QuestionID))

	c.JSON(http.StatusOK, response.IdResponse{Message: "Ответ обновлен"})
}

func (h *Handler) delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NewErrorResponse(c, http.StatusBadRequest, "empty param", "Id не задан")
		return
	}

	if err := h.service.Delete(c, &models.DeleteAnswerDTO{ID: id}); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), id)
		return
	}
	logger.Info("Удален ответ", logger.StringAttr("id", id))

	c.JSON(http.StatusNoContent, response.IdResponse{})
}
