package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"tech/internal/app/handlers"
)

func RunApp() {
	router := gin.Default()
	router.Use(gin.Recovery())

	router.POST("/task", handlers.CreateTask)
	router.GET("/task/:taskId", handlers.GetTaskByID)

	err := router.Run(":8080")
	if err != nil {
		log.Fatalln("Cannot start app:", err.Error())
	}
}
