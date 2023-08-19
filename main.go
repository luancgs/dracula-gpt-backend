package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gofor-little/env"
	"github.com/luancgs/dracula-gpt-backend/src/controllers"
	"github.com/luancgs/dracula-gpt-backend/src/services"
)

var (
	gptService    services.GptService       = services.NewGpt()
	gptController controllers.GptController = controllers.NewGpt(gptService)
)

func main() {
	if err := env.Load("./.env"); err != nil {
		panic(err)
	}

	r := gin.Default()

	r.POST("/query", func(ctx *gin.Context) {
		gptController.CreateQuery(ctx)
	})

	r.Run(":8080")
}
