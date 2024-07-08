package media

import (
	"fmt"
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
	service services.Media
}

func NewHandler(service services.Media) *Handler {
	return &Handler{
		service: service,
	}
}

func Register(api *gin.RouterGroup, service services.Media, middleware *middleware.Middleware) {
	handler := NewHandler(service)

	media := api.Group("media")
	{
		media.POST("", handler.create)
		media.DELETE("", handler.delete)
	}
}

func (h *Handler) create(c *gin.Context) {
	dto := &models.MediaDTO{}
	if err := c.ShouldBind(&dto); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "Отправлены некорректные данные")
		return
	}

	dst := fmt.Sprintf("media/%s", dto.Path)
	if err := h.service.SaveFile(dto.Image, dst); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), dto)
		return
	}
	logger.Info("Добавлено изображение", logger.StringAttr("path", dst), logger.StringAttr("filename", dto.Image.Filename))

	c.JSON(http.StatusCreated, response.DataResponse{Data: dst})
}

func (h *Handler) delete(c *gin.Context) {
	path := c.Query("path")
	if path == "" {
		c.JSON(http.StatusNoContent, response.IdResponse{})
		return
	}

	if err := h.service.Delete(path); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка: "+err.Error())
		error_bot.Send(c, err.Error(), path)
		return
	}
	logger.Info("Удалено изображение", logger.StringAttr("path", path))

	c.JSON(http.StatusNoContent, response.IdResponse{})
}
