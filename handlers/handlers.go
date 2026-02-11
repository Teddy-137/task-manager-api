package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/teddy-137/task_manager_api/models"
	"net/http"
)

func Start() {
	router := gin.Default()

	router.GET("/tasks", tasksHandler)
	router.POST("/tasks", tasksHandler)

	router.GET("/tasks/:id", taskHandler)
	router.PUT("/tasks/:id", taskHandler)
	router.DELETE("/tasks/:id", taskHandler)

	router.Run()
}

func tasksHandler(ctx *gin.Context) {
	var tasks []models.Task
	switch ctx.Request.Method {
	case http.MethodGet:
		if err := db.Find(&tasks).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		ctx.JSON(http.StatusOK, tasks)

	case http.MethodPost:
		var input models.Task

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if err := db.Create(&input).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		ctx.JSON(http.StatusCreated, input)
	}
}

func taskHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var task models.Task
	if err := db.First(&task, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}
	switch ctx.Request.Method {
	case http.MethodGet:
		ctx.JSON(http.StatusOK, task)
	case http.MethodPut:
		var input models.Task
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if input.Title != "" {
			task.Title = input.Title
		}
		if input.Description != "" {
			task.Description = input.Description
		}
		if input.Status != "" {
			task.Status = input.Status
		}

		if err := db.Save(&task).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		ctx.JSON(http.StatusOK, task)
	case http.MethodDelete:
		if err := db.Delete(&task).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return

		}
	}
}
