package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teddy-137/task_manager_api/models"
)

func Start() {

	router := gin.Default()

	router.GET("/tasks", tasksHandler)
	router.POST("/tasks", tasksHandler)
	router.GET("/tasks/:id", taskHundler)
	router.PUT("/tasks/:id", taskHundler)
	router.DELETE("/tasks/:id", taskHundler)

	router.Run()
}

func tasksHandler(ctx *gin.Context) {
	switch ctx.Request.Method {
	case http.MethodGet:
		var tasks []models.Task
		db.Find(&tasks)
		ctx.JSON(http.StatusOK, tasks)
	case http.MethodPost:
		var input models.Task
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
			return
		}

		if err := db.Create(&input).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, input)
	}
}

func taskHundler(ctx *gin.Context) {
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

		db.Save(&task)

		ctx.JSON(http.StatusOK, task)
	case http.MethodDelete:
		db.Delete(&task)
		ctx.JSON(http.StatusNoContent, gin.H{"message": "task deleted."})
	default:
		ctx.JSON(http.StatusMethodNotAllowed, nil)
	}

}
