package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
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

	apiPort, err := env.MustGet("PORT")
	if err != nil {
		fmt.Println("Error getting API port:", err)
		return
	}

	router := gin.Default()

	router.Use(cors.Default())

	router.POST("/query", func(ctx *gin.Context) {
		gptController.CreateQuery(ctx)
	})

	router.Run(fmt.Sprintf(":%s", apiPort))
}
