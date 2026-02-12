package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teddy-137/task_manager_api/internal/domain"
)

type TaskHandler struct {
	Service domain.TaskService
}

func NewTaskHandler(r *gin.Engine, s domain.TaskService) {
	handler := &TaskHandler{
		Service: s,
	}

	r.GET("/tasks", handler.GetAllTasks)
	r.POST("/tasks", handler.CreatTask)
}

func (h *TaskHandler) GetAllTasks(ctx *gin.Context) {
	tasks, err := h.Service.GetAllTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) CreatTask(ctx *gin.Context) {
	var task domain.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateTask(&task); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, task)

}
