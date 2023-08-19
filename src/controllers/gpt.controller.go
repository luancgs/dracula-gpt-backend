package controllers

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luancgs/dracula-gpt-backend/src/entities"
	"github.com/luancgs/dracula-gpt-backend/src/services"
)

type GptController interface {
	CreateQuery(ctx *gin.Context)
}

type gptController struct {
	service services.GptService
}

func NewGpt(service services.GptService) GptController {
	return &gptController{
		service: service,
	}
}

func (c *gptController) CreateQuery(ctx *gin.Context) {
	requestBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error reading request body",
		})
		return
	}

	if len(requestBody) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Request body is empty",
		})
		return
	}

	response, err := c.service.CreateQuery(entities.GptQuery{Prompt: string(requestBody)})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating query",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": response,
	})
}
