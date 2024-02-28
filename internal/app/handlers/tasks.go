package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"tech/internal/entities/service"
)

func CreateTask(c *gin.Context) {
	var (
		request service.CreateTaskRequest
		err     error
	)

	if err = c.ShouldBindJSON(&request); err != nil {
		log.Println("CreateTask handler cannot bind the request:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": "Неверная структура запроса"})
		return
	}

	taskID := service.CreateTask()

	go service.SendRequest(taskID, &request)

	c.JSON(http.StatusOK, gin.H{"id": taskID})
}

func GetTaskByID(c *gin.Context) {
	var (
		taskIDStr = c.Param("taskId")
		taskIDInt int
		err       error
	)

	taskIDInt, err = strconv.Atoi(taskIDStr)
	if err != nil {
		log.Println("GetTaskByID handler cannot convert task ID:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": "Получен неверный ID задачи"})
		return
	}

	task := service.GetTaskByID(uint(taskIDInt))

	if task.ID == 0 {
		c.JSON(http.StatusNotFound, task)
	} else {
		c.JSON(http.StatusOK, task)
	}
}
