package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teddy-137/task_manager_api/internal/domain"
)

type UserHandler struct {
	Service domain.UserService
}

func NewUserHandler(r *gin.Engine, s domain.UserService) {
	handler := &UserHandler{
		Service: s,
	}

	r.GET("/users", handler.GetAllUsers)
	r.POST("/users", handler.CreateUser)
}

func (h *UserHandler) GetAllUsers(ctx *gin.Context) {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var user domain.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
